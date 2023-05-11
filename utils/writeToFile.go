package utils

import (
	"fmt"
	"os"
)

func WriteToFile(filename string, content chan string, done chan bool) {
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
