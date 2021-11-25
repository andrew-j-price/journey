package main

import (
	"fmt"
	"os"
	"reflect"
)

// Add sums two integers
func add(x, y int) int {
	return x + y
}

// Subtract subtracts two integers
func subtract(x, y int) int {
	return x - y
}

func validateArgs(args []string) {
	if debugFlow {
		LoggerDebug.Printf("args: is of type: %v, with value: %v\n", reflect.TypeOf(args), args)
		for i, v := range args {
			fmt.Printf("Index: %v, Value: %v\n", i, v)
		}
	}
	if len(args) == 0 {
		LoggerError.Println("Missing arguments")
		// flag.PrintDefaults()
		os.Exit(1)
	}
	if len(args) != 3 {
		LoggerError.Printf("Incorrect amount of arguments. args: is of length: %v\n", len(args))
		os.Exit(1)
	}
	valid_options := []string{"add", "subtract"}
	valid := isStringInSlice(args[0], valid_options)
	if !valid {
		LoggerDebug.Printf("valid is: %v\n", valid)
		LoggerError.Printf("Math operations must be one of %v\n", valid_options)
		os.Exit(1)
	}
}

func mathMain(args []string) {
	validateArgs(args)
	int1 := StringToInt(args[1], true)
	int2 := StringToInt(args[2], true)
	var result int
	if args[0] == "add" {
		result = add(int1, int2)
	}
	if args[0] == "subtract" {
		result = subtract(int1, int2)
	}
	fmt.Printf("%v result: %v \n", args[0], result)
	os.Exit(0)
}
