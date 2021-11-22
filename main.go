package main

import (
	"flag"
	"os"
)

var debugFlow bool

func init() {
	simpleLogger()
}

func main() {
	enableDebug := flag.Bool("debug", false, "run debug commands")
	runApi := flag.Bool("api", false, "start api")
	runMath := flag.Bool("math", false, "drive -math add 5 7")
	runRandom := flag.Bool("random", false, "just testing things")
	useColor := flag.Bool("color", false, "display colorized output")
	flag.Parse()

	debugFlow = *enableDebug
	if debugFlow {
		LoggerDebug.Println("Debug mode is set to:", debugFlow)
	}

	if *runApi {
		apiMainV1()
		os.Exit(3)
	}
	if *runMath {
		args := flag.Args()
		mathMain(args)
		os.Exit(0)
	}
	if *runRandom {
		// randomTypesAndKind()
		// randomLogMessages()
		os.Exit(0)
	}
	if *useColor {
		colorize(ColorYellow, "Hello world!")
		os.Exit(0)
	}
}
