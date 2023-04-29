package main

import (
	"DetectDee/cmd"
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors:               true,
		FullTimestamp:             true,
		PadLevelText:              true,
		DisableColors:             false,
		EnvironmentOverrideColors: true,
		TimestampFormat:           "2006-01-02 15:04:05",
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	cmd.Execute()
}
