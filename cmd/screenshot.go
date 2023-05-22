package cmd

import (
	"DetectDee/utils"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/chromedp/chromedp"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type screenshotArgsType struct {
	proxy   string
	chrome  bool
	thread  int
	timeout int
	path    string
	file    string
	result  string
}

var (
	screenshotArgs          screenshotArgsType
	screenshotWg            sync.WaitGroup
	screenshotTargets       = make(chan []string, 10)
	execPathNotFound        = "[-] Chrome executable file not found, please install Chrome or specify the chrome.exe path with --path"
	createResultFolderError = "[-] failed to create result folder:%v"
)

func init() {
	screenshotCmd.Flags().StringVar(&screenshotArgs.proxy, "proxy", "", "Make requests over a proxy. e.g. socks5://127.0.0.1:1080")
	screenshotCmd.Flags().StringVar(&screenshotArgs.path, "path", "", "Chrome ExecPath")
	screenshotCmd.Flags().StringVarP(&screenshotArgs.result, "dir", "d", "screenshots", "Folder path of the screenshot")
	screenshotCmd.Flags().BoolVar(&screenshotArgs.chrome, "chrome", false, "Show chrome")
	screenshotCmd.Flags().IntVarP(&screenshotArgs.thread, "thread", "t", 3, "Chrome number")
	screenshotCmd.Flags().StringVarP(&screenshotArgs.file, "file", "f", "result.txt", "Url list file")
	screenshotCmd.Flags().IntVar(&screenshotArgs.timeout, "timeout", 60, "Timeout")
	rootCmd.AddCommand(screenshotCmd)
}

var screenshotCmd = &cobra.Command{
	Use:   "screenshot",
	Short: "screenshot result with chrome headless",
	Long:  ``,
	Run:   screenshot,
}

func screenshot(_ *cobra.Command, _ []string) {
	if Verbose {
		log.Infoln("Debug Mode")
		log.SetLevel(log.DebugLevel)
	}

	targetList, err := utils.ParseResult(screenshotArgs.file)
	if err != nil {
		return
	}

	go utils.AddTarget(targetList, screenshotTargets)

	err = CreateDirIfNotExists(screenshotArgs.result)
	if err != nil {
		log.Errorf(createResultFolderError, err)
		return
	}
	log.Infoln("[+] starting screenshot tasks")

	for i := 1; i < screenshotArgs.thread+1; i++ {
		go navigate(i)
		screenshotWg.Add(1)
	}
	screenshotWg.Wait()
}

func navigate(workerNum int) {
	defer screenshotWg.Done()

	// create context
	opts := append(chromedp.DefaultExecAllocatorOptions[:], chromedp.Flag("headless", !screenshotArgs.chrome), chromedp.ProxyServer(screenshotArgs.proxy), chromedp.Flag("mute-audio", true), chromedp.IgnoreCertErrors, chromedp.DisableGPU, chromedp.NoFirstRun, chromedp.ExecPath(screenshotArgs.path), chromedp.WindowSize(1920, 1080), chromedp.NoDefaultBrowserCheck, chromedp.NoSandbox)

	if screenshotArgs.proxy != "" {
		opts = append(opts, chromedp.Flag("proxy-server", screenshotArgs.proxy))
	}

	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)

	//ctx, cancel = context.WithTimeout(ctx, time.Duration(screenshotArgs.timeout)*time.Second)

	ctx, cancel = chromedp.NewContext(ctx)

	// Check & Make Browser not close
	err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate("https://github.com/piaolin"),
	})
	if err != nil {
		if strings.Contains(err.Error(), "executable file not found") {
			log.Fatalln(execPathNotFound)
		} else {
			log.Fatalln(err)
		}
	}

	// Run
	for len(screenshotTargets) > 0 {
		target := <-screenshotTargets
		name := target[0]
		site := target[1]
		urlStr := target[2]

		log.Debugf("[Worker%2d] starting screenshot task %s,remaining tasks:%d\n", workerNum, urlStr, len(screenshotTargets))
		var buf []byte
		//ctx, cancel = chromedp.NewContext(
		//	ctx,
		//)
		if err := chromedp.Run(ctx, fullScreenshot(urlStr, 100, &buf)); err != nil {
			log.Errorf("[-] Failed to take (URL:%s) screenshot: %v", urlStr, err)
			continue
		}
		pngFilePath := fmt.Sprintf("./%s/%s-%s.png", screenshotArgs.result, name, site)
		if err := ioutil.WriteFile(pngFilePath, buf, 0644); err != nil {
			log.Errorf("[-] Failed to write file %v", err)
			continue
		}

		//if err := page.Close().Do(cdp.WithExecutor(ctx, chromedp.FromContext(ctx).Target)); err != nil {
		//	log.Errorln(err)
		//}

		log.Infof("[+] screenshot success %s", urlStr)
		log.Debugf("[Worker%2d] finished screenshot task %s,remaining tasks:%d\n", workerNum, urlStr, len(screenshotTargets))
	}

	defer cancel()
}

func fullScreenshot(urlstr string, quality int, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),             //chromedp.WaitVisible("style"),
		chromedp.Sleep(1 * time.Second),       //chromedp.OuterHTML(`document.querySelector("body")`, &htmlContent, chromedp.ByJSPath),
		chromedp.FullScreenshot(res, quality), //chromedp.ActionFunc(func(ctx context.Context) error {
		//	_, _, _, _, _, contentSize, err := page.GetLayoutMetrics().Do(ctx)
		//	if err != nil {
		//		return err
		//	}
		//
		//	width, height := int64(math.Ceil(contentSize.Width)), int64(math.Ceil(contentSize.Height))
		//
		//	err = emulation.SetDeviceMetricsOverride(width, height, 1, false).
		//		WithScreenOrientation(&emulation.ScreenOrientation{
		//			Type:  emulation.OrientationTypePortraitPrimary,
		//			Angle: 0,
		//		}).
		//		Do(ctx)
		//	if err != nil {
		//		return err
		//	}
		//
		//	*res, err = page.CaptureScreenshot().
		//		WithQuality(quality).
		//		WithClip(&page.Viewport{
		//			X:      contentSize.X,
		//			Y:      contentSize.Y,
		//			Width:  contentSize.Width,
		//			Height: contentSize.Height,
		//			Scale:  1,
		//		}).Do(ctx)
		//	if err != nil {
		//		return err
		//	}
		//	return nil
		//}),
	}
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
