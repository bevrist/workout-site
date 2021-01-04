package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"time"

	structs "../common"

	"github.com/gorilla/mux"
)

// Global Variables
var apiVersion string = "1.0" //the api version this service implements
//env
var databaseAddress string //listen address of this service

func GetUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	UID := vars["UID"]
	resp, err := http.Get("http://" + databaseAddress + "/userInfo/" + UID)
	if err != nil {
		http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
		log.Println("GetUserInfoHandler: " + err.Error())
		return
	}

	reqBody, _ := ioutil.ReadAll(resp.Body)

	fmt.Fprintf(w, string(reqBody))
}

func ListUsersHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://" + databaseAddress + "/listUsers")
	if err != nil {
		http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
		log.Println("ListUsersHandler: " + err.Error())
		return
	}

	reqBody, _ := ioutil.ReadAll(resp.Body)

	fmt.Fprintf(w, string(reqBody))
}

func UpdateUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	UID := vars["UID"]
	//add UID to user data
	rBody, _ := ioutil.ReadAll(r.Body)
	var userInfo structs.Client
	err := json.Unmarshal(rBody, &userInfo)
	if err != nil {
		http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
		log.Println("ERROR: UpdateUserInfoHandler: " + err.Error())
		return
	}
	userInfo.UID = UID
	userInfoJSON, _ := json.Marshal(userInfo)
	//post user data to database
	resp, err := http.Post("http://"+databaseAddress+"/userInfo/"+UID, "application/json", bytes.NewBuffer(userInfoJSON))
	if err != nil {
		http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
		log.Println("ERROR: UpdateUserInfoHandler() - database: " + err.Error())
		return
	}
	//return database Response
	respBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Fprintf(w, string(respBody))
}

func generateUserBaselineHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	UID := vars["UID"]
	rBody, _ := ioutil.ReadAll(r.Body)
	//Unmarshal and add UID to userInfo struct
	var userInfo structs.Client
	err := json.Unmarshal(rBody, &userInfo)
	if err != nil {
		log.Println("ERROR: generateUserBaselineHandler: " + err.Error())
		http.Error(w, "400 - invalid syntax.", http.StatusBadRequest)
		return
	}
	userInfo.UID = UID

	var bmr float64

	if userInfo.Gender == "male" {
		bmr = 66 + (6.3 * float64(userInfo.Weight)) + (12.9 * float64(userInfo.HeightInches)) - (6.8 * float64(userInfo.Age))
	} else {
		bmr = 655 + (4.3 * float64(userInfo.Weight)) + (4.7 * float64(userInfo.HeightInches)) - (4.7 * float64(userInfo.Age))
	}

	lowday := math.Round(bmr * 1.2)
	normalday := math.Round(bmr * 1.375)
	highday := math.Round(bmr * 1.55)

	NormalDayFat := .25
	NormalDayCarb := .37
	NormalDayProtein := .38
	HighDayFat := .3
	HighDayCarb := .5
	HighDayProtein := .2
	LowDayFat := .41
	LowDayCarb := .32
	LowDayProtein := .27

	var newRecommendation structs.Recommendation

	loc, _ := time.LoadLocation("America/Los_Angeles")
	newRecommendation.ModifiedDate = time.Now().In(loc).Format("2006-01-02")
	newRecommendation.NormalDayFat = int(math.Round((normalday * NormalDayFat) / 9))
	newRecommendation.NormalDayCarb = int(math.Round((normalday * NormalDayCarb) / 4))
	newRecommendation.NormalDayProtein = int(math.Round((normalday * NormalDayProtein) / 4))
	newRecommendation.HighDayFat = int(math.Round((highday * HighDayFat) / 9))
	newRecommendation.HighDayCarb = int(math.Round((highday * HighDayCarb) / 4))
	newRecommendation.HighDayProtein = int(math.Round((highday * HighDayProtein) / 4))
	newRecommendation.LowDayFat = int(math.Round((lowday * LowDayFat) / 9))
	newRecommendation.LowDayCarb = int(math.Round((lowday * LowDayCarb) / 4))
	newRecommendation.LowDayProtein = int(math.Round((lowday * LowDayProtein) / 4))
	newRecommendation.HighDayCalories = int(math.Round(highday))
	newRecommendation.NormalDayCalories = int(math.Round(normalday))
	newRecommendation.LowDayCalories = int(math.Round(lowday))

	userInfo.Recommendation = append(userInfo.Recommendation, newRecommendation)

	userProfile, err := json.Marshal(userInfo)
	if err != nil {
		http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
		log.Println("ERROR: generateUserBaselineHandler() - database: " + err.Error())
	}

	//post user data to database
	resp, err := http.Post("http://"+databaseAddress+"/userInfo/"+UID, "application/json", bytes.NewBuffer(userProfile))
	if err != nil {
		http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
		log.Println("ERROR: generateUserBaselineHandler() - database: " + err.Error())
		return
	}
	//return database Response
	respBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Fprintf(w, string(respBody))
}

func UpdateUserRecommendationsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	UID := vars["UID"]
	WEEK := vars["week"]
	rBody, _ := ioutil.ReadAll(r.Body)
	//add ModifiedDate to post data
	var recommendation structs.Recommendation
	err := json.Unmarshal(rBody, &recommendation)
	loc, _ := time.LoadLocation("America/Los_Angeles")
	recommendation.ModifiedDate = time.Now().In(loc).Format("2006-01-02")

	recommendationJSON, _ := json.Marshal(recommendation)

	//post user data to database
	resp, err := http.Post("http://"+databaseAddress+"/userRecommendation/"+WEEK+"/"+UID, "application/json", bytes.NewBuffer(recommendationJSON))
	if err != nil {
		http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
		log.Println("ERROR: PostUserProfileHandler() - database: " + err.Error())
		return
	}
	//return database Response
	respBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Fprintf(w, string(respBody))
}

func UpdateUserDailyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	UID := vars["UID"]
	WEEK := vars["week"]
	DAY := vars["day"]
	rBody, _ := ioutil.ReadAll(r.Body)

	//post user data to database
	resp, err := http.Post("http://"+databaseAddress+"/userDaily/"+WEEK+"/"+DAY+"/"+UID, "application/json", bytes.NewBuffer(rBody))
	if err != nil {
		http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
		log.Println("ERROR: PostUserProfileHandler() - database: " + err.Error())
		return
	}
	//return database Response
	respBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Fprintf(w, string(respBody))
}

func UpdateUserWeeklylineHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	UID := vars["UID"]
	WEEK := vars["week"]
	rBody, _ := ioutil.ReadAll(r.Body)

	//post user data to database
	resp, err := http.Post("http://"+databaseAddress+"/userWeekly/"+WEEK+"/"+UID, "application/json", bytes.NewBuffer(rBody))
	if err != nil {
		http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
		log.Println("ERROR: PostUserProfileHandler() - database: " + err.Error())
		return
	}
	//return database Response
	respBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Fprintf(w, string(respBody))
}

func main() {
	//populate environment variables
	listenAddress := os.Getenv("BACKEND_LISTEN_ADDRESS")
	databaseAddress = os.Getenv("DATABASE_ADDRESS")
	//set default environment variables
	if listenAddress == "" {
		listenAddress = "0.0.0.0:8090"
	}
	if databaseAddress == "" {
		databaseAddress = "localhost:8050"
	}
	log.Println("Database address: " + databaseAddress)

	//specify routes and start http server
	r := mux.NewRouter()
	r.HandleFunc("/apiVersion", func(w http.ResponseWriter, _ *http.Request) { fmt.Fprint(w, "{\"apiVersion\":"+apiVersion+"}") })
	r.HandleFunc("/userInfo/{UID}", GetUserInfoHandler).Methods(http.MethodGet, http.MethodHead)
	r.HandleFunc("/listUsers", ListUsersHandler).Methods(http.MethodGet, http.MethodHead)
	r.HandleFunc("/userInfo/{UID}", UpdateUserInfoHandler).Methods(http.MethodPost)
	r.HandleFunc("/userWeekly/{week}/{UID}", UpdateUserWeeklylineHandler).Methods(http.MethodPost)
	r.HandleFunc("/userDaily/{week}/{day}/{UID}", UpdateUserDailyHandler).Methods(http.MethodPost)
	r.HandleFunc("/generateUserBaseline/{UID}", generateUserBaselineHandler).Methods(http.MethodPost)
	r.HandleFunc("/userRecommendation/{week}/{UID}", UpdateUserRecommendationsHandler).Methods(http.MethodPost)
	//TODO: implement /listUsers passthrough
	r.HandleFunc("/_healthz", func(w http.ResponseWriter, _ *http.Request) { fmt.Fprint(w, "ok") })
	var handler http.Handler = r
	log.Printf("Backend listening at address " + listenAddress)
	log.Fatal(http.ListenAndServe(listenAddress, handler))
}
