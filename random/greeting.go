package random

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func RandomGreetingMain() {
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
