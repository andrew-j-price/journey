/*
Package: journey

This is an my educational `journey` repo to `drive` my learnings on Golang
*/
package main

import (
	"flag"
	"os"

	"github.com/andrew-j-price/journey/boondocks/boonclient"
	"github.com/andrew-j-price/journey/boondocks/boonserver"
	"github.com/andrew-j-price/journey/logger"
	"github.com/andrew-j-price/journey/random"
	"github.com/andrew-j-price/journey/rps"
)

var debugFlow bool

func init() {
	// SimpleLogger()  // references "./keep/logging.go.keep" file
	logger.PackageLogger()
}

func main() {
	enableDebug := flag.Bool("debug", false, "include debug output")
	runApi := flag.Bool("api", false, "start api")
	runGrpcClient := flag.Bool("boondocks-client", false, "start GRPC client")
	runGrpcServer := flag.Bool("boondocks-server", false, "start GRPC server")
	runIdentity := flag.Bool("identity", false, "start identity app")
	runMath := flag.Bool("math", false, "drive -math add 5 7")
	runRandom := flag.Bool("random", false, "just testing things")
	runRps := flag.Bool("rps", false, "Game of RPS")
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
	if *runGrpcClient {
		boonclient.Main()
		os.Exit(3)
	}
	if *runGrpcServer {
		boonserver.Main()
		os.Exit(3)
	}
	if *runIdentity {
		identityMain()
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
	if *runRps {
		rps.Main()
		os.Exit(3)
	}
	if *useColor {
		colorize(ColorYellow, "Hello world!")
		os.Exit(0)
	}
}
