package main

import (
	"log"
	"os"
)

var criticalLog, errorLog, warnLog, infoLog, debugLog *log.Logger

func init() {
	file, err := os.Create("log-file.txt")
	if err != nil {
		log.Fatal(err)
	}
	flags := log.Ldate | log.Ltime
	criticalLog = log.New(file, "CRITICAL: ", flags)
	errorLog = log.New(file, "ERROR: ", flags)
	warnLog = log.New(file, "WARNING: ", flags)
	infoLog = log.New(file, "INFO: ", flags)
	debugLog = log.New(file, "DEBUG: ", flags)
}

func main() {
	infoLog.Print("That is a milestone")
	errorLog.Print("Got an error here")
	debugLog.Print("Extra information for a debug")
	warnLog.Print("You should be warned about this")
}
