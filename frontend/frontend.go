package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	structs "../common"

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

// GetUserProfileHandler returns user data
func GetUserProfileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"FirstName":"Anthony","LastName":"Hannah","Weight":215,"WaistCirc":11,"HeightInches":72,"LeanBodyMass":15,"Age":27,"Gender":"female"}`) //FIXME
	return
	//FIXME

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

// GetUserBaselineHandler returns user data
func GetUserBaselineHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"LowDay":2599,"NormalDay":2978,"HighDay":3357,"NFatRatio":0.25,"NCarbRatio":0.37,"NProteinRatio":0.38,"HFatRatio":0.3,"HCarbRatio":0.5,"HProteinRatio":0.2,"LFatRatio":0.41,"LCarbRatio":0.32,"LProteinRatio":0.27,"NFatAmount":83,"NCarbAmount":275,"NProteinAmount":283,"HFatAmount":112,"HCarbAmount":420,"HProteinAmount":168,"LFatAmount":118,"LCarbAmount":208,"LProteinAmount":175}
	`) //FIXME
	return
	//FIXME

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
	backResp, err := http.Get("http://" + backendAddress + "/userInfo/" + auth.UID + "/base")
	if err != nil {
		http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
		log.Println("ERROR - Backend: " + err.Error())
		return
	}

	//return backend response
	backBody, _ := ioutil.ReadAll(backResp.Body)
	fmt.Fprintf(w, string(backBody))
}

// SubmitProfileHandler posts user information to backend
func SubmitProfileHandler(w http.ResponseWriter, r *http.Request) {
	// parse form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
		log.Printf("ParseForm() err: %v", err)
		return
	}

	var userInfo structs.UserInfo
	//extract form data
	userInfo.FirstName = r.FormValue("firstName")
	userInfo.LastName = r.FormValue("lastName")
	userInfo.Weight, _ = strconv.Atoi(r.FormValue("weight"))
	userInfo.WaistCirc, _ = strconv.Atoi(r.FormValue("waistCirc"))
	userInfo.HeightInches, _ = strconv.Atoi(r.FormValue("heightInches"))
	userInfo.LeanBodyMass, _ = strconv.Atoi(r.FormValue("leanBodyMass"))
	userInfo.Age, _ = strconv.Atoi(r.FormValue("age"))
	userInfo.Gender = r.FormValue("gender")

	sessionToken := r.FormValue("Session-Token")
	log.Println("session: " + sessionToken)

	//format form data in JSON UserInfo format
	userInfoJSON, err := json.Marshal(userInfo)
	if err != nil {
		log.Fatal(err)
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

	//post user data to backend
	backResp, err := http.Post("http://"+backendAddress+"/userInfo/"+auth.UID, "application/json", bytes.NewBuffer(userInfoJSON))
	if err != nil {
		http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
		log.Println("ERROR - Backend: " + err.Error())
		return
	}

	//return backend Response 	//TODO: make this check backend response code and return based off (update backend first)
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
		listenAddress = "0.0.0.0:8080"
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
	r.HandleFunc("/getUserProfile", GetUserProfileHandler).Methods(http.MethodGet, http.MethodHead)
	r.HandleFunc("/getUserBaseline", GetUserBaselineHandler).Methods(http.MethodGet, http.MethodHead)
	r.HandleFunc("/submitProfile", SubmitProfileHandler).Methods(http.MethodPost)
	r.HandleFunc("/_healthz", func(w http.ResponseWriter, _ *http.Request) { fmt.Fprint(w, "ok") })
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./www/")))
	var handlers http.Handler = r
	log.Println("Frontend listening at address " + listenAddress)
	log.Fatal(http.ListenAndServe(listenAddress, handlers))
}
