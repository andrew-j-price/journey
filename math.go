package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

// Add is our function that sums two integers
func Add(x, y int) (res int) {
	return x + y
}

// Subtract subtracts two integers
func Subtract(x, y int) (res int) {
	return x - y
}

// NOTE: a lot of optimization to do
func math_main(args []string) {
	LoggerInfo.Printf("args: is of type: %v, with value: %v\n", reflect.TypeOf(args), args)

	if len(args) == 0 {
		LoggerError.Println("Missing arguments")
		// flag.PrintDefaults()
		os.Exit(1)
	}
	if len(args) != 3 {
		LoggerError.Printf("Incorrect amount of arguments. args: is of length: %v\n", len(args))
		os.Exit(1)
	}
	for i, v := range args {
		fmt.Printf("Index: %v, Value: %v\n", i, v)
	}
	valid_options := []string{"add", "subtract"}
	valid := stringInSlice(args[0], valid_options)
	fmt.Printf("valid is: %v\n", valid)
	if !valid {
		LoggerError.Printf("Math operations must be one of %v\n", valid_options)
		os.Exit(1)
	}
	int1, err1 := strconv.Atoi(args[1])
	fmt.Println(int1, err1, reflect.TypeOf(int1))
	int2, err2 := strconv.Atoi(args[2])
	fmt.Println(int2, err2, reflect.TypeOf(int2))
	if args[0] == "add" {
		result := Add(int1, int2)
		fmt.Printf("Add result: %v \n", result)
		os.Exit(0)
	}
}
