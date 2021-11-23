package main

import (
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message":"hello world!"}`))
	LoggerInfo.Printf("hello route received %v call and sent response %v\n", http.MethodGet, http.StatusOK)
}

func apiMainV1() {
	listen_address := getEnv("LISTEN_ADDRESS", ":8080")
	LoggerInfo.Printf("Startup binding to %s\n", listen_address)
	http.HandleFunc("/", hello)
	LoggerInfo.Printf("Web server started")
	log.Fatal(http.ListenAndServe(listen_address, nil))
}