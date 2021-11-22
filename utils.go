package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// Params (string) to convert, (bool) if conversion fails, should exit
// Returns (int) of string
func stringToInt(theString string, pass bool) int {
	theInt, err := strconv.Atoi(theString)
	if err != nil {
		fmt.Println(theInt, err, reflect.TypeOf(theInt))
		if !pass {
			fmt.Println("Conversion failed. Exiting.")
			os.Exit(1)
		}
	}
	return theInt
}
