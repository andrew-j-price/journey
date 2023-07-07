package main

import (
	"flag"
	"fmt"
	"os"
	"reflect"

	helper "github.com/andrew-j-price/journey/pkg/helpers"
	logger "github.com/andrew-j-price/journey/pkg/logger"
	"github.com/andrew-j-price/journey/pkg/math"
)

var debugFlow bool

func init() {
	logger.Logger()
}

func main() {
	enableDebug := flag.Bool("debug", false, "include debug output")
	flag.Parse()

	debugFlow = *enableDebug
	if debugFlow {
		logger.Debug.Println("Debug mode is set to:", debugFlow)
	}

	args := flag.Args()
	validateArgs(args)
	math.Main(args)
	os.Exit(0)
}

func validateArgs(args []string) {
	if debugFlow {
		logger.Debug.Printf("args: is of type: %v, with value: %v\n", reflect.TypeOf(args), args)
		for i, v := range args {
			fmt.Printf("Index: %v, Value: %v\n", i, v)
		}
	}
	if len(args) == 0 {
		logger.Error.Println("Missing arguments")
		logger.Info.Println(("Usage: math add 9 6"))
		// flag.PrintDefaults()
		os.Exit(1)
	}
	if len(args) != 3 {
		logger.Error.Printf("Incorrect amount of arguments. args: is of length: %v\n", len(args))
		os.Exit(1)
	}
	valid_options := []string{"add", "subtract"}
	valid := helper.IsStringInSlice(args[0], valid_options)
	if !valid {
		logger.Debug.Printf("valid is: %v\n", valid)
		logger.Error.Printf("Math operations must be one of %v\n", valid_options)
		os.Exit(1)
	}
}
