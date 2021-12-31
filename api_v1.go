package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/andrew-j-price/journey/logger"
)

type event struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type allEvents []event

var events = allEvents{
	{
		ID:          "1",
		Title:       "Introduction to Golang",
		Description: "Come join us for a chance to learn how golang works and get to eventually try it out",
	},
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent event
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newEvent)
	events = append(events, newEvent)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEvent)
}

func getOneEvent(w http.ResponseWriter, r *http.Request) {
	// eventID := mux.Vars(r)["id"]
	eventID := "0"

	for _, singleEvent := range events {
		if singleEvent.ID == eventID {
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message":"hello world!"}`))
	logger.Info.Printf("hello route received %v call and sent response %v\n", http.MethodGet, http.StatusOK)
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"ping":"pong"}`))
	logger.Info.Printf("hello route received %v call and sent response %v\n", http.MethodGet, http.StatusOK)
}

func apiMainV1() {
	listen_address := getEnv("LISTEN_ADDRESS", ":8080")
	logger.Info.Printf("Startup binding to %s\n", listen_address)
	http.HandleFunc("/", hello)
	http.HandleFunc("/home", homeLink)
	http.HandleFunc("/ping", ping)
	// http.HandleFunc("/event", createEvent).Methods("POST")
	// http.HandleFunc("/events/{id}", getOneEvent).Methods("GET")
	// http.HandleFunc("/events", getAllEvents).Methods("GET")
	logger.Info.Printf("Web server started")
	log.Fatal(http.ListenAndServe(listen_address, nil))
}
