package main

import (
	"fmt"
	"os"
	"reflect"
)

// Add sums two integers
func Add(x, y int) int {
	return x + y
}

// Subtract subtracts two integers
func Subtract(x, y int) int {
	return x - y
}

func math_main(args []string) {
	if debugFlow {
		LoggerDebug.Printf("args: is of type: %v, with value: %v\n", reflect.TypeOf(args), args)
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
	if debugFlow {
		for i, v := range args {
			fmt.Printf("Index: %v, Value: %v\n", i, v)
		}
	}
	valid_options := []string{"add", "subtract"}
	valid := isStringInSlice(args[0], valid_options)
	if !valid {
		LoggerDebug.Printf("valid is: %v\n", valid)
		LoggerError.Printf("Math operations must be one of %v\n", valid_options)
		os.Exit(1)
	}
	int1 := stringToInt(args[1], true)
	int2 := stringToInt(args[2], true)
	if args[0] == "add" {
		result := Add(int1, int2)
		fmt.Printf("%v result: %v \n", args[0], result)
	}
	os.Exit(0)
}
