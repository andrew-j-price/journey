package main

import (
	"flag"
	"os"
)

var DebugFlow bool

func init() {
	SimpleLogger()
}

func main() {
	enableDebug := flag.Bool("debug", false, "run debug commands")
	runApi := flag.Bool("api", false, "start api")
	runMath := flag.Bool("math", false, "drive -math add 5 7")
	runRandom := flag.Bool("random", false, "just testing things")
	useColor := flag.Bool("color", false, "display colorized output")
	flag.Parse()

	DebugFlow = *enableDebug
	if DebugFlow {
		LoggerDebug.Println("Debug mode is set to:", DebugFlow)
	}

	if *runApi {
		api_v1_main()
		os.Exit(3)
	}
	if *runMath {
		args := flag.Args()
		math_main(args)
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
