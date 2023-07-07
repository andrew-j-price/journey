package main

import (
	"log"
	"net/http"

	"github.com/andrew-j-price/journey/pkg/helpers"
)

func main() {
	listen_address := helpers.GetEnv("LISTEN_ADDRESS", ":8080")
	log.Printf("Startup binding to %s\n", listen_address)
	mux := http.NewServeMux()
	mux.HandleFunc("/", identityRoot)
	mux.HandleFunc("/healthz", helpers.HttpRoutePing)
	mux.HandleFunc("/ping", helpers.HttpRoutePing)
	log.Printf("Web server starting\n")
	log.Fatal(http.ListenAndServe(listen_address, mux))
}

func identityRoot(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		helpers.HttpRouteDefault(w, r)
	} else if r.URL.Path == "/default" {
		helpers.HttpRouteDefault(w, r)
	} else {
		helpers.HttpRoute404(w, r)
	}
}
