package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

// Clients handles user input requests and responses from database
func Clients(w http.ResponseWriter, r *http.Request) {
	//handling POST data upload
	if r.Method == "POST" {
		// unmarshal the body of POST request as a Client struct
		reqBody, _ := ioutil.ReadAll(r.Body)
		var client Client
		json.Unmarshal(reqBody, &client)

		//TODO move data into backend database

		fmt.Fprintf(w, "SUCCESS")
		out, _ := json.Marshal(client)
		log.Printf("%v\n", string(out))
	} else if r.Method == "GET" { //respond to GET requests
		clients := ClientList{
			Client{FirstName: "Anthony", LastName: "Hanna", Weight: 215, WaistCirc: 33, HeightInches: 75, LeanBodyMass: 180},
			Client{FirstName: "Ray", LastName: "Hanna", Weight: 202, WaistCirc: 29, HeightInches: 68, LeanBodyMass: 159},
		}

		json.NewEncoder(w).Encode(clients)
		// TODO Make a template function to local and populate the html file
		//TODO load data from a database
		// fmt.Fprintf(w, "Backend Ready")
	}
}

func main() {
	http.HandleFunc("/", Clients)

	log.Printf("Backend listening at address localhost:8090\n")
	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatal(err)
	}
}
