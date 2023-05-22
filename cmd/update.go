package cmd

import (
	"DetectDee/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type updateArgsType struct {
	proxy string
	url   string
	file  string
}

var (
	updateArgs updateArgsType
)

func init() {
	updateCmd.Flags().StringVarP(&updateArgs.proxy, "proxy", "p", "", "Make requests over a proxy. e.g. socks5://127.0.0.1:1080")
	updateCmd.Flags().StringVarP(&updateArgs.url, "url", "u", "https://raw.githubusercontent.com/piaolin/DetectDee/main/data.json", "Update data.json using url")
	updateCmd.Flags().StringVarP(&updateArgs.file, "file", "f", "data.json", "Save to file")
	rootCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update data.json using github",
	Long:  ``,
	Run:   update,
}

func update(_ *cobra.Command, _ []string) {
	if Verbose {
		log.Infoln("Debug Mode")
		log.SetLevel(log.DebugLevel)
	}

	newFileName, err := utils.RenameFileByTime(updateArgs.file)
	if err != nil {
		log.Fatal(err)
	}

	err = utils.DownloadFile(updateArgs.file, updateArgs.url)
	if err != nil {
		// if download error, restore filename
		if newFileName != updateArgs.file {
			_ = utils.RestoreFilename(newFileName)
		}
		log.Fatalln("Error downloading file: ", err)
	}

	log.Infof("[+] download file to %s\n", updateArgs.file)
}
