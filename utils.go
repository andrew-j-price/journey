package main

import (
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

func isStringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// Params (string) to convert, (bool) if conversion fails, should program exit
// Returns (int) of string
func stringToInt(theString string, quit bool) int {
	theInt, err := strconv.Atoi(theString)
	if err != nil {
		LoggerError.Println(err)
		LoggerDebug.Printf("During processing of: %v which returned: %v with type: %v", theString, theInt, reflect.TypeOf(theInt))
		if quit {
			LoggerFatal.Panic("Conversion failed. Exiting.")
		}
	}
	return theInt
}
