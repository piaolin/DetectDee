package cmd

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"os"
)

type genArgsType struct {
	fileName string
	//proxy    string
}

var (
	genArgs genArgsType
)

func init() {
	genCmd.Flags().StringVarP(&genArgs.fileName, "filename", "f", "site.md", "Markdown filename to generate")
	//genCmd.Flags().StringVarP(&genArgs.proxy, "proxy", "p", "", "Make requests over a proxy. e.g. socks5://127.0.0.1:1080")
	rootCmd.AddCommand(genCmd)
}

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate markdown for site data in data.json",
	Long:  ``,
	Run:   gen,
}

func gen(_ *cobra.Command, _ []string) {
	if Verbose {
		log.Infoln("Debug Mode")
		log.SetLevel(log.DebugLevel)
	}

	siteData, err := ioutil.ReadFile("data.json")
	if err != nil {
		log.Fatalln(err)
	}
	r := gjson.ParseBytes(siteData)

	f, err := os.Create(genArgs.fileName)
	defer f.Close()

	count := 0
	siteTypeMap := make(map[string][][]string)
	for site, siteBody := range r.Map() {
		siteTypeTemp := []string{site, siteBody.Get("url").String()}
		siteTypeMap[siteBody.Get("type").String()] = append(siteTypeMap[siteBody.Get("type").String()], siteTypeTemp)
		count += 1
	}

	for siteType, siteList := range siteTypeMap {
		f.WriteString(fmt.Sprintf("## %s\n", siteType))
		for index, site := range siteList {
			f.WriteString(fmt.Sprintf("%d. ![](https://www.google.com/s2/favicons?domain=%s) [%s](%s)\n", index+1, site[1], site[0], site[1]))
		}
		f.WriteString("\n")
	}
	log.Infof("[+] Generate %d site to %s\n", count, genArgs.fileName)
}
