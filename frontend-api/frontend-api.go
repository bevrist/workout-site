// Frontend-api provides a public API for clients interacting with the app
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	structs "../common"

	"github.com/gorilla/mux"
)

// Global Variables
var apiVersion string = "1.0" //the api version this service implements
// env vars
var listenAddress, backendAddress, authAddress string

//validateRequest validates session tokens and returns UID
func validateSessionToken(w http.ResponseWriter, sessionToken string) string {
	authResp, err := http.Get("http://" + authAddress + "/getUID/" + sessionToken)
	if err != nil {
		http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
		log.Println("ERROR: - Auth: " + err.Error())
	}
	authBody, _ := ioutil.ReadAll(authResp.Body)
	var authInfo structs.Auth
	json.Unmarshal(authBody, &authInfo)
	if authInfo.IsValid == false {
		return ""
	}
	return authInfo.UID
}

//GetUserInfoHandler returns the users data
func GetUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	sessionToken := r.Header.Get("Session-Token")
	if sessionToken == "" {
		http.Error(w, "428 Precondition Required - missing Session-Token.", http.StatusPreconditionRequired)
		return
	}
	//validate session token
	UID := validateSessionToken(w, sessionToken)
	if UID == "" {
		http.Error(w, "401 Unauthorized.", http.StatusUnauthorized)
		return
	}
	//get user data to return to user
	resp, err := http.Get("http://" + backendAddress + "/userInfo/" + UID)
	if err != nil {
		http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
		log.Println("ERROR: GetUserInfoHandler - Backend: " + err.Error())
	}
	//strip UID field from response
	reqBody, _ := ioutil.ReadAll(resp.Body)
	var userInfo structs.Client
	err = json.Unmarshal(reqBody, &userInfo)
	if err != nil {
		http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
		log.Println("ERROR: GetUserInfoHandler - Backend: " + err.Error())
		return
	}
	userInfo.UID = ""
	userInfoJSON, _ := json.Marshal(userInfo)
	fmt.Fprintf(w, string(userInfoJSON))
}

//UpdateUserInfoHandler update the user profile
func UpdateUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	sessionToken := r.Header.Get("Session-Token")
	if sessionToken == "" {
		http.Error(w, "428 Precondition Required - missing Session-Token.", http.StatusPreconditionRequired)
		return
	}
	//validate session token
	UID := validateSessionToken(w, sessionToken)
	if UID == "" {
		http.Error(w, "401 Unauthorized.", http.StatusUnauthorized)
		return
	}
	//post userInfo data to backend
	rBody, _ := ioutil.ReadAll(r.Body)
	resp, err := http.Post("http://"+backendAddress+"/userInfo/"+UID, "application/json", bytes.NewBuffer(rBody))
	if err != nil {
		http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
		log.Println("ERROR: UpdateUserInfoHandler - Backend: " + err.Error())
		return
	}
	//return backend Response
	respBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Fprintf(w, string(respBody))
}

//UpdateUserWeeklyHandler update the user profile
func UpdateUserWeeklyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	WEEK := vars["week"]
	sessionToken := r.Header.Get("Session-Token")
	if sessionToken == "" {
		http.Error(w, "428 Precondition Required - missing Session-Token.", http.StatusPreconditionRequired)
		return
	}
	//validate session token
	UID := validateSessionToken(w, sessionToken)
	if UID == "" {
		http.Error(w, "401 Unauthorized.", http.StatusUnauthorized)
		return
	}
	//post userInfo data to backend
	rBody, _ := ioutil.ReadAll(r.Body)
	resp, err := http.Post("http://"+backendAddress+"/userWeekly/"+WEEK+"/"+UID, "application/json", bytes.NewBuffer(rBody))
	if err != nil {
		http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
		log.Println("ERROR: UpdateUserWeeklyHandler - Backend: " + err.Error())
		return
	}
	//return backend Response
	respBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Fprintf(w, string(respBody))
}

//UpdateUserDailyHandler update the user profile
func UpdateUserDailyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	WEEK := vars["week"]
	DAY := vars["day"]
	sessionToken := r.Header.Get("Session-Token")
	if sessionToken == "" {
		http.Error(w, "428 Precondition Required - missing Session-Token.", http.StatusPreconditionRequired)
		return
	}
	//validate session token
	UID := validateSessionToken(w, sessionToken)
	if UID == "" {
		http.Error(w, "401 Unauthorized.", http.StatusUnauthorized)
		return
	}
	//post userInfo data to backend
	rBody, _ := ioutil.ReadAll(r.Body)
	resp, err := http.Post("http://"+backendAddress+"/userDaily/"+WEEK+"/"+DAY+"/"+UID, "application/json", bytes.NewBuffer(rBody))
	if err != nil {
		http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
		log.Println("ERROR: UpdateUserDailyHandler - Backend: " + err.Error())
		return
	}
	//return backend Response
	respBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Fprintf(w, string(respBody))
}

//GenerateUserBaselineHandler update the user profile
func GenerateUserBaselineHandler(w http.ResponseWriter, r *http.Request) {
	sessionToken := r.Header.Get("Session-Token")
	if sessionToken == "" {
		http.Error(w, "428 Precondition Required - missing Session-Token.", http.StatusPreconditionRequired)
		return
	}
	//validate session token
	UID := validateSessionToken(w, sessionToken)
	if UID == "" {
		http.Error(w, "401 Unauthorized.", http.StatusUnauthorized)
		return
	}
	//post data to backend
	rBody, _ := ioutil.ReadAll(r.Body)
	resp, err := http.Post("http://"+backendAddress+"/generateUserBaseline/"+UID, "application/json", bytes.NewBuffer(rBody))
	if err != nil {
		http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
		log.Println("ERROR: GenerateUserBaselineHandler - Backend: " + err.Error())
		return
	}
	//return backend Response
	respBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Fprintf(w, string(respBody))
}

//UpdateUserRecommendationsHandler update the user profile
func UpdateUserRecommendationsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	WEEK := vars["week"]
	sessionToken := r.Header.Get("Session-Token")
	if sessionToken == "" {
		http.Error(w, "428 Precondition Required - missing Session-Token.", http.StatusPreconditionRequired)
		return
	}
	//validate session token
	UID := validateSessionToken(w, sessionToken)
	if UID == "" {
		http.Error(w, "401 Unauthorized.", http.StatusUnauthorized)
		return
	}
	//post userInfo data to backend
	rBody, _ := ioutil.ReadAll(r.Body)
	resp, err := http.Post("http://"+backendAddress+"/userRecommendation/"+WEEK+"/"+UID, "application/json", bytes.NewBuffer(rBody))
	if err != nil {
		http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
		log.Println("ERROR: UpdateUserRecommendationsHandler - Backend: " + err.Error())
		return
	}
	//return backend Response
	respBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Fprintf(w, string(respBody))
}

func main() {
	//populate environment variables
	listenAddress = os.Getenv("FRONTEND_API_LISTEN_ADDRESS")
	backendAddress = os.Getenv("BACKEND_ADDRESS")
	authAddress = os.Getenv("AUTH_ADDRESS")
	//set default environment variables
	if listenAddress == "" {
		listenAddress = "0.0.0.0:8888"
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
	r.HandleFunc("/userInfo", GetUserInfoHandler).Methods(http.MethodGet, http.MethodHead)
	r.HandleFunc("/userInfo", UpdateUserInfoHandler).Methods(http.MethodPost)
	r.HandleFunc("/userWeekly/{week}", UpdateUserWeeklyHandler).Methods(http.MethodPost)
	r.HandleFunc("/userDaily/{week}/{day}", UpdateUserDailyHandler).Methods(http.MethodPost)
	r.HandleFunc("/generateUserBaseline", GenerateUserBaselineHandler).Methods(http.MethodPost)
	r.HandleFunc("/userRecommendations/{week}", UpdateUserRecommendationsHandler).Methods(http.MethodPost)
	r.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) { fmt.Fprint(w, "ok") })
	var handlers http.Handler = r
	log.Println("Frontend-api listening at address " + listenAddress)
	log.Fatal(http.ListenAndServe(listenAddress, handlers))
}
