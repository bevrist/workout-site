// Frontend provides a webserver for static site files and exposes an API for retrieving and updating application data
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// Global Variables
var apiVersion string = "1.0" //the api version this service implements
// env
var listenAddress string

//TODO: FIXME update this to send data to frontend-api, or update website to avoid this altogether
// // SubmitProfileHandler posts user information to backend
// func SubmitProfileHandler(w http.ResponseWriter, r *http.Request) {
// 	// parse form data
// 	if err := r.ParseForm(); err != nil {
// 		http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
// 		log.Printf("ERROR: SubmitProfileHandler() - ParseForm() err: %v", err)
// 		return
// 	}
// 	var userInfo structs.UserInfo
// 	//extract form data
// 	userInfo.FirstName = r.FormValue("firstName")
// 	userInfo.LastName = r.FormValue("lastName")
// 	userInfo.Weight, _ = strconv.Atoi(r.FormValue("weight"))
// 	userInfo.WaistCirc, _ = strconv.Atoi(r.FormValue("waistCirc"))
// 	userInfo.HeightInches, _ = strconv.Atoi(r.FormValue("heightInches"))
// 	userInfo.LeanBodyMass, _ = strconv.Atoi(r.FormValue("leanBodyMass"))
// 	userInfo.Age, _ = strconv.Atoi(r.FormValue("age"))
// 	userInfo.Gender = r.FormValue("gender")
// 	// validate user request
// 	sessionToken := r.FormValue("Session-Token")
// 	auth := isRequestValid(w, sessionToken, "SubmitProfileHandler()")
// 	if auth.IsValid == false {
// 		return
// 	}
// 	//format form data in JSON UserInfo format
// 	userInfoJSON, err := json.Marshal(userInfo)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	//post user data to backend
// 	backResp, err := http.Post("http://"+backendAddress+"/userInfo/"+auth.UID, "application/json", bytes.NewBuffer(userInfoJSON))
// 	if err != nil {
// 		http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
// 		log.Println("ERROR: SubmitProfileHandler() - Backend: " + err.Error())
// 		return
// 	}
// 	//return backend Response 	//TODO: make this check backend response code and return based off (update backend first)
// 	backBody, _ := ioutil.ReadAll(backResp.Body)
// 	fmt.Fprintf(w, string(backBody))
// }

func main() {
	//populate environment variables
	listenAddress = os.Getenv("FRONTEND_LISTEN_ADDRESS")
	//set default environment variables
	if listenAddress == "" {
		listenAddress = "0.0.0.0:8080"
	}
	log.Println("Frontend Website URL: " + os.Getenv("FRONTEND_WEBSITE_URL"))

	//specify routes and start http server
	r := mux.NewRouter()
	r.HandleFunc("/apiVersion", func(w http.ResponseWriter, _ *http.Request) { fmt.Fprint(w, "{\"apiVersion\":"+apiVersion+"}") })
	// r.HandleFunc("/submitProfile", SubmitProfileHandler).Methods(http.MethodPost)
	r.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) { fmt.Fprint(w, "ok") })
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./www/")))
	var handlers http.Handler = r
	log.Println("Frontend listening at address " + listenAddress)
	log.Fatal(http.ListenAndServe(listenAddress, handlers))
}
