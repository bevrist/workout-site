package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var backendURL string //global variable for backend server address

// ShowFrontend respond to get requests with webpage, forward POST form to backend
func ShowFrontend(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" { //handle form post and forward responses to backend
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		//extract form data
		weight, _ := strconv.Atoi(r.FormValue("weight"))
		height, _ := strconv.Atoi(r.FormValue("height"))
		waist, _ := strconv.Atoi(r.FormValue("waist"))
		//format form data in JSON
		values := map[string]int{"Weight": weight, "HeightInches": height, "WaistCirc": waist}
		jsonValue, _ := json.Marshal(values)

		//forward form responses to backend server
		response, _ := http.Post(backendURL, "application/json", bytes.NewBuffer(jsonValue))
		responseBody, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("POST:" + string(jsonValue))
		log.Println("Backend Response: " + string(responseBody))

		http.ServeFile(w, r, "index.html")
	} else if r.Method == "GET" { //serve get request
		http.ServeFile(w, r, "index.html") //TODO use templates to have javascript pull from proper backend
	} else {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
}

func main() {
	//take in command line flags
	backendURL = *flag.String("b", "http://localhost:8090", "URL of backend endpoint")
	flag.Parse()

	//check for backend server avalibility
	log.Println("Checking for backend server at " + backendURL)
	backendResponse, err := http.Get(backendURL)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("BackendStatus:" + backendResponse.Status)

	//start server
	http.HandleFunc("/", ShowFrontend)
	log.Println("Frontend listening at address localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
