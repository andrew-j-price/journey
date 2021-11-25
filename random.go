package main

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"time"
)

func randomGreetingMain() {
	hourOfDay := time.Now().Hour()
	greeting, err := randomGetGreeting(hourOfDay)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(greeting)
	fmt.Println("")
}

func randomGetGreeting(hour int) (string, error) {
	var message string
	if hour < 7 {
		err := errors.New("too early for greetings :(")
		return message, err
	} else if hour < 12 {
		message = "Good Morning"
	} else if hour < 18 {
		message = "Good Afternoon"
	} else {
		message = "Good Evening"
	}
	return message, nil
}

func randomLoopMain() {
	randomLoop1()
	randomLoop2()
	randomArray()
	randomSliceAppended()
	randomSliceLiteral()
}

func randomLoop1() {
	fmt.Println("### randomLoop1")
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
	fmt.Println("")
}

func randomLoop2() {
	fmt.Println("### randomLoop2")
	i := 0
	isLessThanFive := true
	for isLessThanFive {
		if i >= 3 {
			isLessThanFive = false
		}
		fmt.Println(i)
		i++
	}
	fmt.Println("")
}

func randomArray() {
	fmt.Println("### randomArray")
	// NOTE: arrays are fixed length
	var langs [3]string
	langs[0] = "Go"
	langs[1] = "Python"
	langs[2] = "Ruby"
	fmt.Println(langs)
	fmt.Println(langs[0])
	fmt.Println("")
}

func randomSliceAppended() {
	fmt.Println("### randomSliceAppended")
	var langs []string
	langs = append(langs, "Go")
	langs = append(langs, "Python")
	langs = append(langs, "Ruby")
	langs = append(langs, "Java")
	fmt.Println(langs)
	fmt.Println(langs[0])
	fmt.Println("")
}

func randomSliceLiteral() {
	fmt.Println("### randomSliceLiteral")
	langs := []string{"Go", "Python", "Ruby"}
	fmt.Println(langs)
	fmt.Println(langs[0])
	fmt.Println("")
}

func randomLogMessages() {
	LoggerInfo.Println("Informational only")
	LoggerWarn.Println("Warning message")
	LoggerError.Println("Error! message")
	// NOTE: `Fatal` or `Panic` instead of `Println`
	// LoggerFatal.Fatal("Quit program")
	// LoggerFatal.Fatal("Quit program with traceback")
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
