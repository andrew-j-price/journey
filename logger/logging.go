package logger

import (
	"io"
	"log"
	"os"
)

var (
	Debug *log.Logger
	Info  *log.Logger
	Warn  *log.Logger
	Error *log.Logger
	Fatal *log.Logger
)

func PackageLogger() {
	logFile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		/*
			panic(err)     // exit code 2 with traceback
			log.Fatal(err) // exit code 1 with log message
		*/
		log.Panic(err) // exit code 2 with log message and traceback

	}

	// NOTABLE OPTIONS: `log.LUTC|`, `log.Lmicroseconds`
	Debug = log.New(logFile, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(logFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warn = log.New(logFile, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(logFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	Fatal = log.New(logFile, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)

	mw := io.MultiWriter(os.Stdout, logFile)
	Debug.SetOutput(mw)
	Info.SetOutput(mw)
	Warn.SetOutput(mw)
	Error.SetOutput(mw)
	Fatal.SetOutput(mw)
}
