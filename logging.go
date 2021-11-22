package main

import (
	"io"
	"log"
	"os"
)

var (
	LoggerDebug *log.Logger
	LoggerInfo  *log.Logger
	LoggerWarn  *log.Logger
	LoggerError *log.Logger
	LoggerFatal *log.Logger
)

func SimpleLogger() {
	logFile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		/*
			panic(err)     // exit code 2 with traceback
			log.Fatal(err) // exit code 1 with log message
		*/
		log.Panic(err) // exit code 2 with log message and traceback

	}

	// NOTABLE OPTIONS: `log.LUTC|`, `log.Lmicroseconds`
	LoggerDebug = log.New(logFile, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	LoggerInfo = log.New(logFile, "INFO:  ", log.Ldate|log.Ltime|log.Lshortfile)
	LoggerWarn = log.New(logFile, "WARN:  ", log.Ldate|log.Ltime|log.Lshortfile)
	LoggerError = log.New(logFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	LoggerFatal = log.New(logFile, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)

	mw := io.MultiWriter(os.Stdout, logFile)
	LoggerDebug.SetOutput(mw)
	LoggerInfo.SetOutput(mw)
	LoggerWarn.SetOutput(mw)
	LoggerError.SetOutput(mw)
	LoggerFatal.SetOutput(mw)
}
