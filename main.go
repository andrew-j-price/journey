package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Color string

const (
	ColorBlack  Color = "\u001b[30m"
	ColorRed          = "\u001b[31m"
	ColorGreen        = "\u001b[32m"
	ColorYellow       = "\u001b[33m"
	ColorBlue         = "\u001b[34m"
	ColorReset        = "\u001b[0m"
)

func colorize(color Color, message string) {
	fmt.Println(string(color), message, string(ColorReset))
}

func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message":"hello world!"}`))
}

func apiMain() {
	listen_address := getEnv("LISTEN_ADDRESS", ":8080")
	fmt.Printf("Server listening %s ...", listen_address)
	http.HandleFunc("/", hello)
	fmt.Println("Server started")
	log.Fatal(http.ListenAndServe(listen_address, nil))
}

func main() {
	runApi := flag.Bool("api", false, "start api")
	useColor := flag.Bool("color", false, "display colorized output")
	flag.Parse()

	if *runApi {
		apiMain()
		return
	}
	if *useColor {
		colorize(ColorYellow, "Hello world!")
		return
	}
	fmt.Println("Hello world!")
}
