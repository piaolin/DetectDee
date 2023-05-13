package cmd

import (
	"bufio"
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type mirrorArgsType struct {
	proxy   string
	chrome  bool
	thread  int
	timeout int
	path    string
	file    string
	result  string
}

var (
	mirrorArgs              mirrorArgsType
	mirrorWg                sync.WaitGroup
	targets                 = make(chan string, 1000)
	execPathNotFound        = "[-] Chrome executable file not found, please install Chrome or specify the chrome.exe path with --path"
	urlFileReadError        = "[-] failed to read targets file:%v"
	createResultFolderError = "[-] failed to create result folder:%v"
)

func init() {
	mirrorCmd.Flags().StringVar(&mirrorArgs.proxy, "proxy", "", "Make requests over a proxy. e.g. socks5://127.0.0.1:1080")
	mirrorCmd.Flags().StringVar(&mirrorArgs.path, "path", "", "Chrome ExecPath")
	mirrorCmd.Flags().StringVarP(&mirrorArgs.result, "result", "r", "screenshots", "Folder path of the screenshot")
	mirrorCmd.Flags().BoolVar(&mirrorArgs.chrome, "chrome", false, "Show chrome")
	mirrorCmd.Flags().IntVarP(&mirrorArgs.thread, "thread", "t", 3, "Chrome number")
	mirrorCmd.Flags().StringVarP(&mirrorArgs.file, "file", "f", "result.txt", "Url list file")
	mirrorCmd.Flags().IntVar(&mirrorArgs.timeout, "timeout", 60, "Timeout")
	rootCmd.AddCommand(mirrorCmd)
}

var mirrorCmd = &cobra.Command{
	Use:   "mirror",
	Short: "mirror data.json using github",
	Long:  ``,
	Run:   mirror,
}

func mirror(_ *cobra.Command, _ []string) {
	if Verbose {
		log.Infoln("Debug Mode")
		log.SetLevel(log.DebugLevel)
	}
	file, err := os.Open(mirrorArgs.file)
	if err != nil {
		log.Errorf(urlFileReadError, err)
		return
	}
	defer file.Close()
	err = CreateDirIfNotExists(mirrorArgs.result)
	if err != nil {
		log.Errorf(createResultFolderError, err)
		return
	}
	log.Infoln("[+] starting screenshot tasks")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "" {
			targets <- scanner.Text()
		}
	}
	for i := 1; i < mirrorArgs.thread+1; i++ {
		go navigate(i)
		mirrorWg.Add(1)
	}
	mirrorWg.Wait()
}

func navigate(workerNum int) {
	defer mirrorWg.Done()

	// create context
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", !mirrorArgs.chrome),
		chromedp.ProxyServer(mirrorArgs.proxy),
		chromedp.Flag("mute-audio", true),
		chromedp.IgnoreCertErrors,
		chromedp.DisableGPU,
		chromedp.NoFirstRun,
		chromedp.ExecPath(mirrorArgs.path),
		chromedp.WindowSize(1920, 1080),
		chromedp.NoDefaultBrowserCheck,
		chromedp.NoSandbox,
	)

	if mirrorArgs.proxy != "" {
		opts = append(opts, chromedp.Flag("proxy-server", mirrorArgs.proxy))
	}

	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)

	ctx, cancel = context.WithTimeout(ctx, time.Duration(mirrorArgs.timeout)*time.Second)

	ctx, cancel = chromedp.NewContext(ctx)

	// Check
	err := chromedp.Run(
		ctx,
		chromedp.Tasks{
			chromedp.Navigate("about:blank"),
		},
	)
	if err != nil {
		if strings.Contains(err.Error(), "executable file not found") {
			log.Fatalln(execPathNotFound)
		} else {
			log.Fatalln(err)
		}
	}

	// Run
	for len(targets) > 0 {
		url := <-targets
		log.Debugf("[Worker%2d] starting screenshot task %s,remaining tasks:%d\n", workerNum, url, len(targets))
		var buf []byte
		if err := chromedp.Run(ctx, fullScreenshot(url, 100, &buf)); err != nil {
			log.Errorf("[-] Failed to take (URL:%s) screenshot: %v", url, err)
			continue
		}
		pngFilePath := fmt.Sprintf("./%s/%s.png", mirrorArgs.result, getDomain(url))
		if err := ioutil.WriteFile(pngFilePath, buf, 0644); err != nil {
			log.Errorf("[-] Failed to write file %v", err)
			continue
		}
		log.Infof("[+] screenshot success %s", url)
		log.Debugf("[Worker%2d] finished screenshot task %s,remaining tasks:%d\n", workerNum, url, len(targets))

	}

	defer cancel()
}

func fullScreenshot(urlstr string, quality int64, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		//chromedp.WaitVisible("style"),
		chromedp.Sleep(5 * time.Second),
		//chromedp.OuterHTML(`document.querySelector("body")`, &htmlContent, chromedp.ByJSPath),
		chromedp.ActionFunc(func(ctx context.Context) error {
			_, _, _, _, _, contentSize, err := page.GetLayoutMetrics().Do(ctx)
			if err != nil {
				return err
			}

			width, height := int64(math.Ceil(contentSize.Width)), int64(math.Ceil(contentSize.Height))

			err = emulation.SetDeviceMetricsOverride(width, height, 1, false).
				WithScreenOrientation(&emulation.ScreenOrientation{
					Type:  emulation.OrientationTypePortraitPrimary,
					Angle: 0,
				}).
				Do(ctx)
			if err != nil {
				return err
			}

			*res, err = page.CaptureScreenshot().
				WithQuality(quality).
				WithClip(&page.Viewport{
					X:      contentSize.X,
					Y:      contentSize.Y,
					Width:  contentSize.Width,
					Height: contentSize.Height,
					Scale:  1,
				}).Do(ctx)
			if err != nil {
				return err
			}
			return nil
		}),
	}
}

func getDomain(urlstr string) string {
	u, err := url.Parse(urlstr)
	if err != nil {
		fmt.Println(err)
		return "getdomain_error"
	}
	return u.Hostname()
}

func CreateDirIfNotExists(dirName string) error {
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		err := os.MkdirAll(dirName, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}
