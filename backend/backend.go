package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Global Variables
var listenAddress = flag.String("l", ":8090", "Webserver listen address") //holds http server listen address

// Declaring Struct for user information. This will be pulled from the database I think.
type Client struct {
	FirstName    string `json:"FirstName"`
	LastName     string `json:"LastName"`
	Weight       int    `json:"Weight"`
	WaistCirc    int    `json:"WaistCirc"`
	HeightInches int    `json:"HeightInches"`
	LeanBodyMass int    `json:"LeanBodyMass"`
}

type ClientList []Client

// Clients handles user input requests and responses from database
func Clients(w http.ResponseWriter, r *http.Request) {
	//handling POST data upload
	if r.Method == "POST" {
		// unmarshal the body of POST request as a Client struct
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		var client Client
		err = json.Unmarshal(reqBody, &client)
		if err != nil {
			log.Fatal(err)
		}

		//TODO move data in client struct into backend database

		//respond to POST with server status
		fmt.Fprintf(w, "SUCCESS")

		//print received data to output //FIXME remove
		out, _ := json.Marshal(client)
		log.Printf("%v\n", string(out))
	} else if r.Method == "GET" { //respond to GET requests
		clients := ClientList{
			Client{FirstName: "Anthony", LastName: "Hanna", Weight: 215, WaistCirc: 33, HeightInches: 75, LeanBodyMass: 180},
			Client{FirstName: "Ray", LastName: "Hanna", Weight: 202, WaistCirc: 29, HeightInches: 68, LeanBodyMass: 159},
		}

		//respond with JSON object
		json.NewEncoder(w).Encode(clients)

		//TODO Make a template function to local and populate the html file
		//TODO load data from a database
	}
}

func main() {
	//parse command line flags (declared with global variables)
	flag.Parse()

	//start http server
	http.HandleFunc("/", Clients)
	log.Printf("Backend listening at address " + *listenAddress)
	if err := http.ListenAndServe(*listenAddress, nil); err != nil {
		log.Fatal(err)
	}
}
