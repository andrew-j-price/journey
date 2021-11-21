package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message":"hello world!"}`))
}

func api_v1_main() {
	listen_address := getEnv("LISTEN_ADDRESS", ":8080")
	fmt.Printf("Server listening %s ...\n", listen_address)
	http.HandleFunc("/", hello)
	fmt.Println("Server started")
	log.Fatal(http.ListenAndServe(listen_address, nil))
}
