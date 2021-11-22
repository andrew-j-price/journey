package main

import (
	"io"
	"log"
	"os"
)

var (
	LoggerInfo  *log.Logger
	LoggerWarn  *log.Logger
	LoggerError *log.Logger
	LoggerFatal *log.Logger
)

func SimpleLogger() {
	logFile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
		// TODO: analyze panic
		// panic(err)
	}

	// NOTABLE OPTIONS: `log.LUTC|`, `log.Lmicroseconds`
	LoggerInfo = log.New(logFile, "INFO:  ", log.Ldate|log.Ltime|log.Lshortfile)
	LoggerWarn = log.New(logFile, "WARN:  ", log.Ldate|log.Ltime|log.Lshortfile)
	LoggerError = log.New(logFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	LoggerFatal = log.New(logFile, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)

	mw := io.MultiWriter(os.Stdout, logFile)
	LoggerInfo.SetOutput(mw)
	LoggerWarn.SetOutput(mw)
	LoggerError.SetOutput(mw)
	LoggerFatal.SetOutput(mw)
}
