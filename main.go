package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	runApi := flag.Bool("api", false, "start api")
	useColor := flag.Bool("color", false, "display colorized output")
	flag.Parse()

	if *runApi {
		api_v1_main()
		os.Exit(3)
	}
	if *useColor {
		colorize(ColorYellow, "Hello world!")
		return
	}
	fmt.Println("Hello world!")
}
