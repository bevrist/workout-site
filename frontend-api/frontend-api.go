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
func validateSessionToken(w http.ResponseWriter, sessionToken string, isAdmin bool) string {
	authResp, err := http.Get("http://" + authAddress + "/getUID/" + sessionToken)
	if err != nil {
		http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
		log.Println("ERROR: - Auth: " + err.Error())
		return ""
	}
	authBody, _ := ioutil.ReadAll(authResp.Body)
	var authInfo structs.Auth
	json.Unmarshal(authBody, &authInfo)
	if authInfo.IsValid == false {
		return ""
	}
	if isAdmin && authInfo.IsAdmin == false {
		return ""
	}
	return authInfo.UID
}

//GetUserInfoHandler returns the users data
func GetUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	authCookie, err := r.Cookie("Authorization")
	//check auth token exists
	if err != nil {
		http.Error(w, "401 - unauthorized", http.StatusUnauthorized)
		return
	}
	//validate session token
	UID := validateSessionToken(w, authCookie.Value, false)
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
	authCookie, err := r.Cookie("Authorization")
	//check auth token exists
	if err != nil {
		http.Error(w, "401 - unauthorized", http.StatusUnauthorized)
		return
	}
	//validate session token
	UID := validateSessionToken(w, authCookie.Value, false)
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
	authCookie, err := r.Cookie("Authorization")
	//check auth token exists
	if err != nil {
		http.Error(w, "401 - unauthorized", http.StatusUnauthorized)
		return
	}
	//validate session token
	UID := validateSessionToken(w, authCookie.Value, false)
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
	authCookie, err := r.Cookie("Authorization")
	//check auth token exists
	if err != nil {
		http.Error(w, "401 - unauthorized", http.StatusUnauthorized)
		return
	}
	//validate session token
	UID := validateSessionToken(w, authCookie.Value, false)
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
	authCookie, err := r.Cookie("Authorization")
	//check auth token exists
	if err != nil {
		http.Error(w, "401 - unauthorized", http.StatusUnauthorized)
		return
	}
	//validate session token
	UID := validateSessionToken(w, authCookie.Value, false)
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

//AdminGetUserInfoHandler returns a users data for an admin request
func AdminGetUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	userUID := r.Header.Get("User-UID")
	authCookie, err := r.Cookie("Authorization")
	//check auth token exists
	if err != nil {
		http.Error(w, "401 - unauthorized", http.StatusUnauthorized)
		return
	}
	//validate session token and verify Admin
	UID := validateSessionToken(w, authCookie.Value, true)
	if UID == "" {
		http.Error(w, "403 Forbidden.", http.StatusForbidden)
		return
	}
	//get user data to return to Admin
	resp, err := http.Get("http://" + backendAddress + "/userInfo/" + userUID)
	if err != nil {
		http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
		log.Println("ERROR: AdminGetUserInfoHandler - Backend: " + err.Error())
	}
	//strip UID field from response
	reqBody, _ := ioutil.ReadAll(resp.Body)
	var userInfo structs.Client
	err = json.Unmarshal(reqBody, &userInfo)
	if err != nil {
		http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
		log.Println("ERROR: AdminGetUserInfoHandler - Backend: " + err.Error())
		return
	}
	userInfo.UID = ""
	userInfoJSON, _ := json.Marshal(userInfo)
	fmt.Fprintf(w, string(userInfoJSON))

}

//AdminGetUserInfoHandler returns a users data for an admin request
func ListUsersHandler(w http.ResponseWriter, r *http.Request) {
	authCookie, err := r.Cookie("Authorization")
	//check auth token exists
	if err != nil {
		http.Error(w, "401 - unauthorized", http.StatusUnauthorized)
		return
	}
	//validate session token and verify Admin
	UID := validateSessionToken(w, authCookie.Value, true)
	if UID == "" {
		http.Error(w, "403 Forbidden.", http.StatusForbidden)
		return
	}
	//get user data to return to Admin
	resp, err := http.Get("http://" + backendAddress + "/listUsers")
	if err != nil {
		http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
		log.Println("ERROR: ListUsersHandler - Backend: " + err.Error())
	}
	reqBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Fprintf(w, string(reqBody))
}

//AdminUpdateUserRecHandler update the user profile
func AdminUpdateUserRecHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	WEEK := vars["week"]
	userUID := r.Header.Get("User-UID")
	authCookie, err := r.Cookie("Authorization")
	//check auth token exists
	if err != nil {
		http.Error(w, "401 - unauthorized", http.StatusUnauthorized)
		return
	}
	//validate session token and verify Admin
	UID := validateSessionToken(w, authCookie.Value, true)
	if UID == "" {
		http.Error(w, "403 Forbidden.", http.StatusForbidden)
		return
	}
	//post userInfo data to backend
	rBody, _ := ioutil.ReadAll(r.Body)
	resp, err := http.Post("http://"+backendAddress+"/userRecommendation/"+WEEK+"/"+userUID, "application/json", bytes.NewBuffer(rBody))
	if err != nil {
		http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
		log.Println("ERROR: AdminUpdateUserRecHandler - Backend: " + err.Error())
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
	// Admin handlers
	r.HandleFunc("/admin/listUsers", ListUsersHandler).Methods(http.MethodGet, http.MethodHead)
	r.HandleFunc("/admin/userInfo", AdminGetUserInfoHandler).Methods(http.MethodGet, http.MethodHead)
	r.HandleFunc("/admin/userRecommendation/{week}", AdminUpdateUserRecHandler).Methods(http.MethodPost)
	r.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) { fmt.Fprint(w, "ok") })
	var handlers http.Handler = r
	log.Println("Frontend-api listening at address " + listenAddress)
	log.Fatal(http.ListenAndServe(listenAddress, handlers))
}
