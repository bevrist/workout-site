package main

import (
	"os"

	structs "../common"

	"bytes"
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
// env
var listenAddress, backendAddress, authAddress string

//getUID gets user id from Session Token (auth service)
func getUID(sessionToken string) (structs.Auth, error) {
	resp, err := http.Get("http://" + authAddress + "/getUID/" + sessionToken)
	if err != nil {
		return structs.Auth{IsValid: false, UID: ""}, err
	}
	//extract auth data from response body
	respBody, _ := ioutil.ReadAll(resp.Body)
	var auth structs.Auth
	json.Unmarshal(respBody, &auth)
	return auth, nil
}

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
	auth, err := getUID(sessionToken)
	if err != nil {
		http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
		log.Println("ERROR - Auth: " + err.Error())
		return
	}
	// respond with 401 if Auth response says token is invalid
	if auth.IsValid == false {
		http.Error(w, "401 unauthorized.", http.StatusUnauthorized)
		log.Println("401 unauthorized.")
		return
	}

	//request user data from backend
	backResp, err := http.Get("http://" + backendAddress + "/userInfo/" + auth.UID)
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
	// parse form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
		log.Printf("ParseForm() err: %v", err)
		return
	}
	//extract form data
	weight, _ := strconv.Atoi(r.FormValue("weight"))
	height, _ := strconv.Atoi(r.FormValue("height"))
	waist, _ := strconv.Atoi(r.FormValue("waist"))
	//FIXME complete this form
	sessionToken := r.FormValue("Session-Token")

	// get UID from session token (auth service)
	auth, err := getUID(sessionToken)
	if err != nil {
		http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
		log.Println("ERROR - Auth: " + err.Error())
		return
	}
	// respond with 401 if Auth response says token is invalid
	if auth.IsValid == false {
		http.Error(w, "401 unauthorized.", http.StatusUnauthorized)
		log.Println("401 unauthorized.")
		return
	}

	//format form data in JSON UserInfo format
	userInfo := structs.UserInfo{FirstName: "Test", LastName: "User", Weight: weight, WaistCirc: waist, HeightInches: height, LeanBodyMass: 10} //FIXME complete this form submission
	userInfoJSON, err := json.Marshal(userInfo)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(userInfoJSON))

	//post user data to backend
	backResp, err := http.Post("http://"+backendAddress+"/userInfo/"+auth.UID, "application/json", bytes.NewBuffer(userInfoJSON))
	if err != nil {
		http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
		log.Println("ERROR - Backend: " + err.Error())
		return
	}

	//return backend response
	backBody, _ := ioutil.ReadAll(backResp.Body)
	fmt.Fprintf(w, string(backBody))
}

func main() {
	//populate environment variables
	listenAddress = os.Getenv("FRONTEND_LISTEN_ADDRESS")
	backendAddress = os.Getenv("BACKEND_ADDRESS")
	authAddress = os.Getenv("AUTH_ADDRESS")
	//set default environment variables
	if listenAddress == "" {
		listenAddress = "localhost:8080"
	}
	if backendAddress == "" {
		backendAddress = "localhost:8090"
	}
	if authAddress == "" {
		authAddress = "localhost:8070"
	}
	log.Println("Backend address: " + backendAddress)
	log.Println("Auth address: " + authAddress)

	//specify routes and start http server
	r := mux.NewRouter()
	r.HandleFunc("/apiVersion", func(w http.ResponseWriter, _ *http.Request) { fmt.Fprint(w, "{\"apiVersion\":"+apiVersion+"}") })
	r.HandleFunc("/getUserData", GetUserDataHandler).Methods(http.MethodGet, http.MethodHead)
	r.HandleFunc("/submitForm", SubmitFormHandler).Methods(http.MethodPost)
	r.HandleFunc("/_healthz", func(w http.ResponseWriter, _ *http.Request) { fmt.Fprint(w, "ok") })
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./www/")))
	var handlers http.Handler = r
	log.Println("Frontend listening at address " + listenAddress)
	log.Fatal(http.ListenAndServe(listenAddress, handlers))
}
