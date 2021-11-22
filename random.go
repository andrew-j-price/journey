package main

import (
	"fmt"
	"reflect"
)

func randomLogMessages() {
	LoggerInfo.Println("Informational only")
	LoggerWarn.Println("Warning message")
	LoggerError.Println("Error! message")
	// NOTE: `Fatal` instead of `Println`
	// LoggerFatal.Fatal("Quit program")
}

func randomTypesAndKind() {
	b := true
	s := ""
	n := 1
	f := 1.0
	a := []string{"foo", "bar", "baz"}

	fmt.Println("### TypeOf ###")
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(reflect.TypeOf(n))
	fmt.Println(reflect.TypeOf(f))
	fmt.Println(reflect.TypeOf(a))

	fmt.Println("### Kind ###")
	fmt.Println(reflect.ValueOf(b).Kind())
	fmt.Println(reflect.ValueOf(s).Kind())
	fmt.Println(reflect.ValueOf(n).Kind())
	fmt.Println(reflect.ValueOf(f).Kind())
	fmt.Println(reflect.ValueOf(a).Kind())
	fmt.Println(reflect.ValueOf(a).Index(0).Kind()) // For slices and strings
}
