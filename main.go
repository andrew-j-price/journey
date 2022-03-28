/*
Package: journey

This is an my educational `journey` repo to `drive` my learnings on Golang
*/
package main

import (
	"flag"
	"os"

	"github.com/andrew-j-price/journey/logger"
	"github.com/andrew-j-price/journey/random"
)

var debugFlow bool

func init() {
	// SimpleLogger()  // references "./keep/logging.go.keep" file
	logger.PackageLogger()
}

func main() {
	enableDebug := flag.Bool("debug", false, "include debug output")
	runApi := flag.Bool("api", false, "start api")
	runMath := flag.Bool("math", false, "drive -math add 5 7")
	runRandom := flag.Bool("random", false, "just testing things")
	useColor := flag.Bool("color", false, "display colorized output")
	flag.Parse()

	debugFlow = *enableDebug
	if debugFlow {
		logger.Debug.Println("Debug mode is set to:", debugFlow)
	}

	if *runApi {
		apiMain()
		os.Exit(3)
	}
	if *runMath {
		args := flag.Args()
		mathMain(args)
		os.Exit(0)
	}
	if *runRandom {
		// random.RandomGreetingMain()
		// random.RandomLoopMain()
		// random.RandomTypesAndKind()
		// random.RandomLogMessages()
		// random.FakeDataMain(debugFlow)
		// random.JsonDataMain()
		// random.AttestMain()
		random.GithubGraphqlMain()
		os.Exit(0)
	}
	if *useColor {
		colorize(ColorYellow, "Hello world!")
		os.Exit(0)
	}
}
