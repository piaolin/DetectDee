package utils

import (
	"bufio"
	"os"
	"strings"
)

func ParseResult(filename string) ([][]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var resultList [][]string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if scanner.Text() != "" {
			splitStr := strings.Split(scanner.Text(), ", ")
			resultList = append(resultList, splitStr)
		}
	}
	return resultList, err
}
