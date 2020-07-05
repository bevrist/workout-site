package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// Global Variables
var backendURL = flag.String("b", "http://localhost:8090", "URL of backend endpoint") //holds backend server address (can be modified by command line flag e.g.'-b http://otherhost.com:5000')
var listenAddress = flag.String("l", ":8080", "Weberver listen address")              //holds http server listen address
var templates = template.Must(template.ParseFiles("frontend.html"))                   //holds all parsed templates

// TemplateVars declares all information to be passed to the HTML template
type TemplateVars struct {
	BackendURL string
}

//serveTemplate processes template and returns final html file
func serveTemplate(w http.ResponseWriter, r *http.Request) {
	vars := &TemplateVars{BackendURL: *backendURL} //set template variables
	templates.Execute(w, vars)
}

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
		response, _ := http.Post(*backendURL, "application/json", bytes.NewBuffer(jsonValue))
		responseBody, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("POST:" + string(jsonValue))
		log.Println("Backend Response: " + string(responseBody))

		serveTemplate(w, r)
	} else if r.Method == "GET" { //serve get request
		serveTemplate(w, r)
	} else {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
}

func main() {
	//parse command line flags (declared with global variables)
	flag.Parse()

	log.Println("hiiiii" + *backendURL + ":" + *listenAddress)

	//check for backend server avalibility
	log.Println("Checking for backend server at " + *backendURL)
	backendResponse, err := http.Get(*backendURL)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("BackendStatus:" + backendResponse.Status)

	//start http server
	http.HandleFunc("/", ShowFrontend)
	log.Println("Frontend listening at address " + *listenAddress)
	if err := http.ListenAndServe(*listenAddress, nil); err != nil {
		log.Fatal(err)
	}
}
