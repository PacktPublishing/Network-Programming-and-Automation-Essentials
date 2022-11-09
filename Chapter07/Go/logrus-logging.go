package main

import (
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	log.SetLevel(log.ErrorLevel)
}

func main() {

	log.Debug("Debug is supressed in error level")
	log.Info("This info won't show in error level")
	log.Error("Got an error here")
}
