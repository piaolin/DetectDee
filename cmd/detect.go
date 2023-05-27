package cmd

import (
	"DetectDee/utils"
	"fmt"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"regexp"
	"sort"
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
	retry     int
	file      string
	google    bool
	email     []string
	phone     []string
	output    string
	//unique    bool
	token string
}

var (
	detectArgs          detectArgsType
	wg                  sync.WaitGroup
	nonSiteData         = "[-] There is no site data of %s\n"
	existInfo           = "[+] %-15s %-15s: %s\n"
	existOutputInfo     = "%s, %s, %s\n"
	nonExistInfo        = "[-] %-15s %-15s: non exists\n"
	reqErrorInfo        = "[!] %-15s %-15s: %s requests error, retry %d/%d\n"
	sleepMap            = make(map[string]int64)
	sleepChannel        = make(map[string]chan bool)
	searchInfo          = "[+] %-15s %-15s: Please check in %s\n"
	disableSiteInfo     = "[!] %-15s %-15s: data of %s is temporarily unavailable\n"
	detectCompletedInfo = "[+] Detect completed, save to %s\n"
	nsfwInfo            = "[!] %-15s %-15s: %s is nsfw\n"
	writeContent        = make(chan string)
	writeDone           = make(chan bool)
	whoisMap            sync.Map
	detectResultMap     sync.Map
	chatgptUserLabel    = "[+] %s ChatgptUserLabel: %s\n"
)

func init() {
	detectCmd.Flags().StringSliceVarP(&detectArgs.name, "name", "n", []string{}, "name[s], e.g. piaolin,poq79,SomeOneYouLike")
	detectCmd.Flags().StringSliceVarP(&detectArgs.email, "email", "e", []string{}, "email[s], e.g. mail@gmail.com,45715485@qq.com")
	detectCmd.Flags().StringSliceVarP(&detectArgs.phone, "phone", "p", []string{}, "phone[s], e.g. 15725753684,13575558962")
	//_ = detectCmd.MarkFlagRequired("name")
	detectCmd.Flags().StringSliceVarP(&detectArgs.site, "site", "s", []string{}, "Limit analysis to just the listed sites. Add multiple options to specify more than one site.")
	detectCmd.Flags().BoolVarP(&detectArgs.check, "check", "c", false, "self-check")
	detectCmd.Flags().StringVar(&detectArgs.proxy, "proxy", "", "Make requests over a proxy. e.g. socks5://127.0.0.1:1080")
	detectCmd.Flags().IntVarP(&detectArgs.timeout, "timeout", "t", 10, "Time (in seconds) to wait for response to requests")
	detectCmd.Flags().BoolVar(&detectArgs.isNSFW, "nsfw", false, "Include checking of NSFW sites from default list.")
	detectCmd.Flags().IntVarP(&detectArgs.retry, "retry", "r", 3, "Retry times after request failed")
	detectCmd.Flags().BoolVar(&detectArgs.precisely, "precisely", false, "Check precisely")
	detectCmd.Flags().StringVarP(&detectArgs.file, "file", "f", "data.json", "Site data file")
	detectCmd.Flags().BoolVarP(&detectArgs.google, "google", "g", false, "Show google search result")
	detectCmd.Flags().StringVarP(&detectArgs.output, "output", "o", "result.txt", "Result file")
	detectCmd.Flags().StringVar(&detectArgs.token, "token", "", "chatgpt api token")

	//detectCmd.Flags().BoolVar(&detectArgs.unique, "unique", false, "Make new requests client for each site")
	rootCmd.AddCommand(detectCmd)
}

var detectCmd = &cobra.Command{
	Use:   "detect",
	Short: "Hunt down social media accounts by username, email or phone across social networks",
	Long:  ``,
	Run:   detect,
}

func detect(_ *cobra.Command, _ []string) {
	if len(detectArgs.name) != 0 {
		log.Infoln("Detect for", detectArgs.name)
	}
	if len(detectArgs.email) != 0 {
		log.Infoln("Detect for", detectArgs.email)
	}
	if len(detectArgs.phone) != 0 {
		log.Infoln("Detect for", detectArgs.phone)
	}

	if Verbose {
		log.Infoln("Debug Mode")
		log.SetLevel(log.DebugLevel)
	}

	go utils.WriteToFile(detectArgs.output, writeContent, writeDone)

	log.Debugln(detectArgs)
	siteData, err := ioutil.ReadFile(detectArgs.file)
	if err != nil {
		log.Fatalln(err)
	}
	r := gjson.ParseBytes(siteData)

	siteDataMap := r.Map()
	if len(detectArgs.site) != 0 {
		log.Infoln("Detect to ", detectArgs.site)
		siteDataMap = make(map[string]gjson.Result)
		for _, site := range detectArgs.site {
			if siteData := r.Get(strings.ToLower(site)); siteData.Exists() {
				siteDataMap[site] = r.Get(strings.ToLower(site))
			} else {
				log.Infof(nonSiteData, site)
			}
		}
	}

	for _, v := range detectArgs.name {
		whoisMap.Store(v, make(map[string]int))
		detectResultMap.Store(v, make([][]string, 0))
	}
	for _, v := range detectArgs.email {
		whoisMap.Store(v, make(map[string]int))
		detectResultMap.Store(v, make([][]string, 0))
	}
	for _, v := range detectArgs.phone {
		whoisMap.Store(v, make(map[string]int))
		detectResultMap.Store(v, make([][]string, 0))
	}

	// set delay for each requests to site
	for site, siteBody := range siteDataMap {
		sleepMap[site] = 0
		if sleep := siteBody.Get("sleep"); sleep.Exists() && sleep.Value() != nil {
			sleepMap[site] = sleep.Int()
			sleepChannel[site] = make(chan bool, 1)
			log.Debugf("Delay %ds before each request to the %s\n", sleep.Int(), site)
		}
	}

	if detectArgs.check {
		log.Infoln("self check")
	} else {
		// detect by name
		for _, name := range detectArgs.name {
			for site, siteBody := range siteDataMap {
				go detectSite(name, site, "username", siteBody)
				wg.Add(1)
			}
		}

		// detect by email
		for _, email := range detectArgs.email {
			for site, siteBody := range siteDataMap {
				go detectSite(email, site, "email", siteBody)
				wg.Add(1)
			}
		}

		// detect by phone
		for _, phone := range detectArgs.phone {
			for site, siteBody := range siteDataMap {
				go detectSite(phone, site, "phone", siteBody)
				wg.Add(1)
			}
		}
	}

	wg.Wait()
	writeDone <- true

	fmt.Println()

	for _, name := range detectArgs.name {

		type registrantCountry struct {
			registrantCountry string
			count             int
		}

		tmpMap, _ := whoisMap.Load(name)
		userWhoisMap, _ := tmpMap.(map[string]int)

		var country []registrantCountry
		for k, v := range userWhoisMap {
			country = append(country, registrantCountry{k, v})
		}

		sort.Slice(country, func(i, j int) bool {
			return country[i].count > country[j].count // 降序
		})
		log.Infof("[+] %s Registrant Country: %v", name, country)

	}
	//tmpList, _ := detectResultMap.Load("piaolin")
	//userResultList, _ := tmpList.([][]string)
	//log.Infoln(userResultList)

	if detectArgs.token != "" {

		for _, name := range detectArgs.name {
			tmpList, _ := detectResultMap.Load(name)
			userResultList, _ := tmpList.([][]string)
			urls := ""
			for _, v := range userResultList {
				urls += v[2]
				urls += ", "
			}
			userLabel, err := utils.ChatUserLabel(detectArgs.token, urls)
			if err != nil {
				log.Debugln("[-] ", err)
			} else {
				log.Infof(chatgptUserLabel, name, userLabel.Choices[0].Message.Content)
			}
		}
	}

	log.Infof(detectCompletedInfo, detectArgs.output)
}

func detectSite(name, site, nameType string, siteBody gjson.Result) {
	defer wg.Done()

	// flag for precisely mode
	flag := false

	url := siteBody.Get("url").String()

	if siteBody.Get("isNSFW").Bool() && !detectArgs.isNSFW {
		log.Debugf(nsfwInfo, name, site, site)
		return
	}

	// Check name
	if nameCheck := siteBody.Get("nameCheck"); nameCheck.Exists() && len(nameCheck.Str) != 0 && nameType == "username" {
		match, err := regexp.MatchString(nameCheck.String(), name)
		if err != nil {
			log.Fatalln(err)
		}

		if !match {
			log.Debugf(nonExistInfo, name, site)
			return
		}
	}

	whoisData := siteBody.Get("whois")

	detectReq := siteBody.Get("detect").Array()

	detectCount := len(detectReq) - 1
	for index, detectData := range detectReq {

		// check status
		if status := detectData.Get("status"); status.Exists() && !status.Bool() {
			log.Debugf(disableSiteInfo, name, site, site)
			break
		}
		retryTimes := 0

		if search := detectData.Get("search"); search.Exists() {
			// google options
			if detectArgs.google {
				searchString := search.String()
				if strings.Contains(searchString, "%s") {
					searchString = fmt.Sprintf(searchString, name)
				}
				searchUrl := fmt.Sprintf(detectData.Get("searchUrl").String(), searchString)
				log.Infof(searchInfo, name, site, searchUrl)
				break
			}
		} else if strings.Contains(detectData.Get("type").String(), nameType) && detectUser(name, site, url, index, retryTimes, detectCount, &flag, detectData, whoisData) {
			continue
		} else {
			break
		}
	}
}

func detectUser(name, site, originalUrl string, requestTimes, retryTimes, detectCount int, flag *bool, detectData gjson.Result, whoisData gjson.Result) bool {

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
		time.Sleep(time.Duration(sleepMap[site]) * time.Second)
		<-ch
	}

	var rep *resty.Response
	for retryTimes < detectArgs.retry {
		var err error
		rep, err = utils.Requests(url, body, detectArgs.proxy, header, detectArgs.timeout)
		if err != nil {
			retryTimes += 1
			log.Debugf(reqErrorInfo, name, site, url, retryTimes, detectArgs.retry)
			time.Sleep(time.Duration(sleepMap[site]) * time.Second)
			continue
		}
		break
	}

	if retryTimes == detectArgs.retry {
		//log.Infof(reqErrorInfo, name, site, url, retryTimes, detectArgs.retry)
		return false
	}

	//log.Infoln(rep.Request.Body)
	//log.Debugln(rep.Status())
	//log.Debugln(rep.Proto())
	//
	//log.Debugln(rep.Request.Header)
	//
	//log.Debugln(rep.Status(), rep.String())

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
		*flag = true
	}

	userPage := detectData.Get("userPage").String()
	if strings.Contains(userPage, "%s") {
		userPage = fmt.Sprintf(userPage, name)
	}

	// precisely mode
	if !*flag {
		log.Debugf(nonExistInfo, name, site)
		return false
	} else if !detectArgs.precisely {
		// flag=true && precisely=false
		log.Infof(existInfo, name, site, userPage)
		writeWhois(name, site, whoisData)
		writeResult(name, site, originalUrl)
		writeContent <- fmt.Sprintf(existOutputInfo, name, site, userPage)
		return false
	} else if requestTimes == detectCount {
		// flag=true && precisely=true && last request
		log.Infof(existInfo, name, site, userPage)
		writeWhois(name, site, whoisData)
		writeResult(name, site, originalUrl)
		writeContent <- fmt.Sprintf(existOutputInfo, name, site, userPage)
		return true
	} else {
		return true
	}
}

func writeResult(name, site, url string) {
	tmpList, _ := detectResultMap.Load(name)
	userResultList, _ := tmpList.([][]string)
	userResultList = append(userResultList, []string{name, site, url})
	detectResultMap.Store(name, userResultList)
}

func writeWhois(name, site string, whoisData gjson.Result) {
	tmpMap, _ := whoisMap.Load(name)
	userWhoisMap, _ := tmpMap.(map[string]int)

	registrantCountry := whoisData.Get("RegistrantCountry").String()
	if _, ok := userWhoisMap[registrantCountry]; ok {
		userWhoisMap[registrantCountry] += 1
	} else {
		userWhoisMap[registrantCountry] = 1
	}

}
