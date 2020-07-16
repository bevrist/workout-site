package main

import (
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

// TODO setup listen address and other variables (backend, auth) through env

//Auth contains authentication data
// type Auth struct {
// 	IsValid bool
// 	UID     string
// }

//getUID gets user id from Session Token (auth service)
func getUID(sessionToken string) (structs.Auth, error) {
	resp, err := http.Get("http://localhost:8070/getUID/" + sessionToken) //FIXME sub host with variable AUTH
	if err != nil {
		return structs.Auth{false, ""}, err
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
	userInfo := structs.UserInfo{"Test", "User", weight, waist, height, 10}
	userInfoJSON, err := json.Marshal(userInfo)
	if err != nil {
		log.Fatal(err)
	}

	//post user data to backend
	backResp, err := http.Post("http://localhost:8090/userInfo/"+auth.UID, "application/json", bytes.NewBuffer(userInfoJSON)) //FIXME sub host with variable BACKEND
	if err != nil {
		http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
		log.Println("ERROR - Backend: " + err.Error())
		return
	}

	//return backend response
	backBody, _ := ioutil.ReadAll(backResp.Body)
	fmt.Fprintf(w, string(backBody))
	log.Println("POST:" + string(userInfoJSON)) //FIXME: remove
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
