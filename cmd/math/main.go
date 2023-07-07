package main

import (
	"flag"
	"os"

	"github.com/andrew-j-price/journey/pkg/logger"
)

var debugFlow bool

func init() {
	// SimpleLogger()  // references "./keep/logging.go.keep" file
	logger.PackageLogger()
}

func main() {
	enableDebug := flag.Bool("debug", false, "include debug output")
	flag.Parse()

	debugFlow = *enableDebug
	if debugFlow {
		logger.Debug.Println("Debug mode is set to:", debugFlow)
	}

	args := flag.Args()
	mathMain(args)
	os.Exit(0)
}
