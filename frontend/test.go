package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Global Variables
var apiVersion string = "1.0" //the api version this service implements

// TODO setup listen address and other variables (backend, auth) through env

// GetUserDataHandler returns user data
func GetUserDataHandler(w http.ResponseWriter, r *http.Request) {
	// validate session token
	sessionToken := r.Header.Get("Session-Token")
	if sessionToken == "" {
		http.Error(w, "428 Precondition Required - Session-Token header Missing.", http.StatusPreconditionRequired)
		log.Println("Token Missing")
		return
	}
	// get UID from session token (auth service)
	resp, err := http.Get("http://localhost:8070/getUID/" + sessionToken) //FIXME sub host with variable AUTH
	if err != nil {
		http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
		log.Println("ERROR - Auth: " + err.Error())
		return
	}

	// Extract Auth struct from response body
	type Auth struct {
		IsValid bool
		UID     string
	}
	var auth Auth
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &auth)

	// respond with 401 if Auth response says token is invalid
	if auth.IsValid == false {
		http.Error(w, "401 unauthorized.", http.StatusUnauthorized)
		log.Println("401 unauthorized.")
		return
	}

	//request user data from backend
	backResp, err := http.Get("http://localhost:8090/userInfo/" + auth.UID) //FIXME sub host with variable BACKEND
	if err != nil {
		http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
		log.Println("ERROR - Backend: " + err.Error())
		return
	}

	//return backend response
	backBody, _ := ioutil.ReadAll(backResp.Body)
	fmt.Fprintf(w, string(backBody))
}

// SubmitFormHandler posts user information to backend
func SubmitFormHandler(w http.ResponseWriter, r *http.Request) {
	// validate session token
	sessionToken := r.Header.Get("Session-Token")
	if sessionToken == "" {
		http.Error(w, "428 Precondition Required - Session-Token header Missing.", http.StatusPreconditionRequired)
		return
	}
	// get UID from session token (auth service)
	// TODO validate auth and return 401 on error

	// parse form data
	if err := r.ParseForm(); err != nil {
		log.Printf("ParseForm() err: %v", err)
		http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
		return
	}
	//extract form data
	weight, _ := strconv.Atoi(r.FormValue("weight"))
	height, _ := strconv.Atoi(r.FormValue("height"))
	waist, _ := strconv.Atoi(r.FormValue("waist"))
	//TODO add session-token as parameter
	//format form data in JSON
	values := map[string]int{"Weight": weight, "HeightInches": height, "WaistCirc": waist}
	jsonValue, _ := json.Marshal(values)

	log.Println("POST:" + string(jsonValue)) //TODO: fix
}

func main() {
	//specify routes and start http server
	r := mux.NewRouter()
	r.HandleFunc("/apiVersion", func(w http.ResponseWriter, _ *http.Request) { fmt.Fprint(w, "{\"apiVersion\":"+apiVersion+"}") })
	r.HandleFunc("/getUserData", GetUserDataHandler).Methods(http.MethodGet, http.MethodHead)
	r.HandleFunc("/submitForm", SubmitFormHandler).Methods(http.MethodPost)
	r.HandleFunc("/_healthz", func(w http.ResponseWriter, _ *http.Request) { fmt.Fprint(w, "ok") })
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./www/")))
	var handlers http.Handler = r
	log.Printf("Frontend listening at address :8080")
	log.Fatal(http.ListenAndServe(":8080", handlers))
}
