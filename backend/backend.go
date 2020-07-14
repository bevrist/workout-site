package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Global Variables
var apiVersion string = "1.0" //the api version this service implements

// TODO setup listen address and other variables through env

//UserInfo holds user information
type UserInfo struct {
	FirstName    string
	LastName     string
	Weight       int
	WaistCirc    int
	HeightInches int
	LeanBodyMass int
}

//GetUserInfoHandler returns json object of user data
func GetUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	// //extract UID from URL
	// vars := mux.Vars(r)
	// UID := vars["UID"]

	//FIXME remove
	response := UserInfo{FirstName: "Anthony", LastName: "Hanna", Weight: 202, WaistCirc: 29, HeightInches: 68, LeanBodyMass: 159}

	//TODO load user data from database by their UID

	//respond with JSON object
	json.NewEncoder(w).Encode(response)
}

//UpdateUserInfoHandler updates user info from POST data
func UpdateUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	// // extract UID from URL
	// vars := mux.Vars(r)
	// UID := vars["UID"]

	// unmarshal the body of POST request as a Client struct
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var userInfo UserInfo
	err = json.Unmarshal(reqBody, &userInfo)
	if err != nil {
		log.Fatal(err)
	}

	//TODO move data in client struct into backend database

	//respond to POST and redirect to previous page
	fmt.Fprintf(w, "<!DOCTYPE html>SUCCESS <script>window.history.back();</script>")

	//print received data to output //FIXME remove
	out, _ := json.Marshal(userInfo)
	log.Printf("%v\n", string(out))
}

func main() {
	//specify routes and start http server
	r := mux.NewRouter()
	r.HandleFunc("/apiVersion", func(w http.ResponseWriter, _ *http.Request) { fmt.Fprint(w, "{\"apiVersion\":"+apiVersion+"}") })
	r.HandleFunc("/userInfo/{UID}", GetUserInfoHandler).Methods(http.MethodGet, http.MethodHead)
	r.HandleFunc("/userInfo/{UID}", UpdateUserInfoHandler).Methods(http.MethodPost)
	r.HandleFunc("/_healthz", func(w http.ResponseWriter, _ *http.Request) { fmt.Fprint(w, "ok") })
	var handler http.Handler = r
	log.Printf("Auth listening at address :8090")
	log.Fatal(http.ListenAndServe(":8090", handler))
}
