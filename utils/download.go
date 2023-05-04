package utils

import (
	"github.com/go-resty/resty/v2"
)

func DownloadFile(filepath string, url string) error {

	client := resty.New()

	_, err := client.R().
		SetOutput(filepath).
		Get(url)
	return err
}
