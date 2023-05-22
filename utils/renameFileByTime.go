package utils

import (
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
	"time"
)

func RenameFileByTime(filename string) (string, error) {
	if _, err := os.Stat(filename); err == nil {
		splitStr := strings.Split(filename, ".")
		newFileName := splitStr[0] + "_" + time.Now().Format("20060102_150405") + "." + splitStr[1]
		log.Debugln(newFileName)
		err := os.Rename(filename, newFileName)
		return newFileName, err
	} else {
		return filename, nil
	}
}

func RestoreFilename(filename string) error {
	if _, err := os.Stat(filename); err == nil {
		splitStr := strings.Split(filename, "_")
		fileExt := strings.Split(filename, ".")
		newFileName := splitStr[0] + "." + fileExt[len(fileExt)-1]
		log.Debugln(newFileName)
		return os.Rename(filename, newFileName)
	} else {
		return nil
	}
}
