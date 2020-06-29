package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Declaring Struct for user information. This will be pulled from the database I think.
type Client struct {
	FirstName    string `json:"FirstName"`
	LastName     string `json:"LastName"`
	Weight       int    `json:"Weight"`
	WaistCirc    int    `json:"WaistCirc"`
	HeightInches int    `json:"Height"`
	LeanBodyMass int    `json:"LeanBodyMass"`
}

type ClientList []Client

// Inputting info into the array and encoding it into the json format
func Clients(w http.ResponseWriter, r *http.Request) {
	clients := ClientList{
		Client{FirstName: "Anthony", LastName: "Hanna", Weight: 215, WaistCirc: 33, HeightInches: 75, LeanBodyMass: 180},
		Client{FirstName: "Ray", LastName: "Hanna", Weight: 202, WaistCirc: 29, HeightInches: 68, LeanBodyMass: 159},
	}

	json.NewEncoder(w).Encode(clients)
	// TODO Make a template function to local and populate the html file
}

func handleRequests() {
	http.HandleFunc("/", Clients)
	log.Fatal(http.ListenAndServe(":80", nil))

}

func main() {
	handleRequests()
}
