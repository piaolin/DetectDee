package cmd

import (
	"DetectDee/utils"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"regexp"
	"strings"
	"sync"
	"time"
)

type detectArgsType struct {
	name      []string
	site      []string
	check     bool
	proxy     string
	timeout   int
	isNSFW    bool
	precisely bool
	//unique    bool
}

var (
	detectArgs   detectArgsType
	wg           sync.WaitGroup
	existInfo    = "[+] %-15s %-15s: %s\n"
	nonExistInfo = "[-] %-15s %-15s: non exists\n"
	sleepTable   = make(map[string]int64)
	sleepChannel = make(map[string]chan bool)
)

func init() {
	detectCmd.Flags().StringSliceVarP(&detectArgs.name, "name", "n", []string{}, "name[s]")
	_ = detectCmd.MarkFlagRequired("name")
	detectCmd.Flags().StringSliceVarP(&detectArgs.site, "site", "s", []string{}, "Limit analysis to just the listed sites. Add multiple options to specify more than one site.")
	detectCmd.Flags().BoolVarP(&detectArgs.check, "check", "c", false, "self-check")
	detectCmd.Flags().StringVarP(&detectArgs.proxy, "proxy", "p", "", "Make requests over a proxy. e.g. socks5://127.0.0.1:1080")
	detectCmd.Flags().IntVarP(&detectArgs.timeout, "timeout", "t", 60, "Time (in seconds) to wait for response to requests")
	detectCmd.Flags().BoolVar(&detectArgs.isNSFW, "nsfw", false, "Include checking of NSFW sites from default list.")
	detectCmd.Flags().BoolVar(&detectArgs.precisely, "precisely", false, "Check precisely")
	//detectCmd.Flags().BoolVar(&detectArgs.unique, "unique", false, "Make new requests client for each site")
	rootCmd.AddCommand(detectCmd)
}

var detectCmd = &cobra.Command{
	Use:   "detect",
	Short: "Hunt down social media accounts by username across social networks",
	Long:  ``,
	Run:   detect,
}

func detect(_ *cobra.Command, _ []string) {
	log.Infoln("Detect for", detectArgs.name)
	if Verbose {
		log.Infoln("Debug Mode")
		log.SetLevel(log.DebugLevel)
	}

	log.Debugln(detectArgs)
	siteData, err := ioutil.ReadFile("data.json")
	if err != nil {
		log.Fatalln(err)
	}
	r := gjson.ParseBytes(siteData)

	siteDataMap := r.Map()
	if len(detectArgs.site) != 0 {
		log.Infoln("Detect to ", detectArgs.site)
		siteDataMap = make(map[string]gjson.Result)
		for _, site := range detectArgs.site {
			siteDataMap[site] = r.Get(strings.ToLower(site))
		}
	}

	for site, siteBody := range siteDataMap {
		// set delay for each requests to site
		sleepTable[site] = 0
		if sleep := siteBody.Get("sleep"); sleep.Exists() && sleep.Value() != nil {
			sleepTable[site] = sleep.Int()
			sleepChannel[site] = make(chan bool, 1)
			log.Debugf("Delay %ds before each request to the %s\n", sleep.Int(), site)
		}
	}

	for _, name := range detectArgs.name {
		for site, siteBody := range siteDataMap {
			go detectSite(name, site, siteBody)
			wg.Add(1)
		}
	}
	wg.Wait()
}

func detectSite(name, site string, siteBody gjson.Result) {
	defer wg.Done()

	// flag for precisely mode
	flag := false

	if !siteBody.Get("isNFSW").Bool() && detectArgs.isNSFW {
		return
	}

	// Check name
	if nameCheck := siteBody.Get("nameCheck"); nameCheck.Exists() && len(nameCheck.Str) != 0 {
		match, err := regexp.MatchString(nameCheck.String(), name)
		if err != nil {
			log.Fatalln(err)
		}

		if !match {
			log.Debugf(nonExistInfo, name, site)
			return
		}
	}

	detectReq := siteBody.Get("detect").Array()

	for index, detectData := range detectReq {

		// set header
		header := make(map[string]string)
		if h := detectData.Get("header"); h.Exists() {
			headerTemp, ok := h.Value().(map[string]interface{})
			if !ok {
				log.Fatalln(headerTemp, "is invalid")
			}
			for key, value := range headerTemp {
				strKey := fmt.Sprintf("%v", key)
				strValue := fmt.Sprintf("%v", value)
				header[strKey] = strValue
			}

		}

		// if body is set, set method to post
		var body string
		if d := detectData.Get("body"); d.Exists() {
			body = d.String()
			if strings.Contains(body, "%s") {
				body = fmt.Sprintf(body, name)
			}
		}

		// set url with name
		var url string
		if u := detectData.Get("url"); u.Exists() {
			url = u.String()
			if strings.Contains(url, "%s") {
				url = fmt.Sprintf(url, name)
			}
		} else {
			log.Fatalln("Why no URL???")
		}

		// delay
		if ch, ok := sleepChannel[site]; ok {
			ch <- true
			time.Sleep(time.Duration(sleepTable[site]) * time.Second)
			<-ch
		}

		rep, err := utils.Requests(url, body, detectArgs.proxy, header, detectArgs.timeout)
		if err != nil {
			log.Errorln(err)
			continue
		}

		//log.Infoln(rep.Request.Body)
		//log.Debugln(rep.Status())
		//log.Debugln(rep.Proto())
		//
		//log.Debugln(rep.Request.Header)
		//
		//log.Debugln(rep.String())

		// statusCode, existRegex must both be true
		statusCode := detectData.Get("statusCode")
		existRegex := detectData.Get("existRegex")

		statusCodeCheck := !statusCode.Exists() || strings.Contains(rep.Status(), statusCode.String())
		existRegexCheck := true
		if existRegex.Exists() && len(existRegex.Str) != 0 {
			match, err := regexp.MatchString(existRegex.String(), rep.String())
			if err != nil {
				log.Errorln(err)
			}
			existRegexCheck = match
		}

		// nonExistRegex have a veto power
		nonExistRegex := detectData.Get("nonExistRegex")
		nonExistRegexCheck := false
		if nonExistRegex.Exists() && len(nonExistRegex.Str) != 0 {
			match, err := regexp.MatchString(nonExistRegex.String(), rep.String())
			if err != nil {
				log.Errorln(err)
			}
			nonExistRegexCheck = match
		}
		if statusCodeCheck && existRegexCheck && !nonExistRegexCheck {
			flag = true
		}

		userPage := detectData.Get("userPage").String()
		if strings.Contains(userPage, "%s") {
			userPage = fmt.Sprintf(userPage, name)
		}

		// precisely mode
		if !flag {
			log.Debugf(nonExistInfo, name, site)
			break
		} else if !detectArgs.precisely {
			// flag=true && precisely=false
			log.Infof(existInfo, name, site, userPage)
			break
		} else if index == len(detectReq)-1 {
			// flag=true && precisely=true
			log.Infof(existInfo, name, site, userPage)
		}
	}

}
