/*
Package: journey

This is an my educational `journey` repo to `drive` my learnings on Golang
*/
package main

import (
	"flag"
	"os"

	"github.com/andrew-j-price/journey/pkg/logger"
	"github.com/andrew-j-price/journey/pkg/random"
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

	// random.RandomGreetingMain()
	// random.RandomLoopMain()
	// random.RandomTypesAndKind()
	// random.RandomLogMessages()
	random.FakeDataMain(debugFlow)
	// random.JsonDataMain()
	// random.AttestMain()
	// random.GithubGraphqlMain()
	// rps.LocalTesting()
	os.Exit(0)

}
