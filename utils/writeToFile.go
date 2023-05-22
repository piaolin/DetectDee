package utils

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

func WriteToFile(filename string, content chan string, done chan bool) {

	_, err := RenameFileByTime(filename)
	if err != nil {
		log.Fatal(err)
	}

	file, _ := os.Create(filename)
	defer file.Close()

	for {
		select {
		case line := <-content:
			_, _ = fmt.Fprint(file, line)
		case <-done:
			return
		}
	}
}
