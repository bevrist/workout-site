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
		log.Fatal("GetUserInfoHandler: " + err.Error())
	}

	reqBody, _ := ioutil.ReadAll(resp.Body)

	fmt.Fprintf(w, string(reqBody))
}

func CreateUserBaseline(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	UID := vars["UID"]
	WEEK := vars["Week"]
	// unmarshal the body of POST request as a Client struct
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("ERROR: UpdateUserProfileHandler: " + err.Error())
	}
	var userInfo structs.Client
	err = json.Unmarshal(reqBody, &userInfo)
	if err != nil {
		log.Println("ERROR: UpdateUserProfileHandler: " + err.Error())
		http.Error(w, "400 - invalid syntax.", http.StatusBadRequest)
		return
	}

	var bmr float64
	dbWeight := userInfo.Week[0].Day[0].Weight

	if userInfo.Gender == "male" {
		bmr = 66 + (6.3 * float64(dbWeight)) + (12.9 * float64(userInfo.HeightInches)) - (6.8 * float64(userInfo.Age))
	} else {
		bmr = 655 + (4.3 * float64(dbWeight)) + (4.7 * float64(userInfo.HeightInches)) - (4.7 * float64(userInfo.Age))
	}

	lowday := math.Round(bmr * 1.2)
	normalday := math.Round(bmr * 1.375)
	highday := math.Round(bmr * 1.55)

	NFatRatio := .25
	NCarbRatio := .37
	NProteinRatio := .38
	HFatRatio := .3
	HCarbRatio := .5
	HProteinRatio := .2
	LFatRatio := .41
	LCarbRatio := .32
	LProteinRatio := .27

	var newRecommendation structs.Recommendation

	newRecommendation.NormalDayFat = int(math.Round((normalday * NFatRatio) / 9))
	newRecommendation.NormalDayCarb = int(math.Round((normalday * NCarbRatio) / 4))
	newRecommendation.NormalDayProtein = int(math.Round((normalday * NProteinRatio) / 4))
	newRecommendation.HighDayFat = int(math.Round((highday * HFatRatio) / 9))
	newRecommendation.HighDayCarb = int(math.Round((highday * HCarbRatio) / 4))
	newRecommendation.HighDayProtein = int(math.Round((highday * HProteinRatio) / 4))
	newRecommendation.LowDayFat = int(math.Round((lowday * LFatRatio) / 9))
	newRecommendation.LowDayCarb = int(math.Round((lowday * LCarbRatio) / 4))
	newRecommendation.LowDayProtein = int(math.Round((lowday * LProteinRatio) / 4))

	//log.Println(bmr, lowday, normalday, highday, NFatRatio, NCarbRatio, NProteinRatio, HFatRatio, HCarbRatio, HProteinRatio, LFatRatio, LCarbRatio, LProteinRatio,
	//NFatAmount, NCarbAmount, NProteinAmount, HFatAmount, HCarbAmount, HProteinAmount, LFatAmount, LCarbAmount, LProteinAmount)
	//respond with JSON object
	RecommendationJSON, err := json.Marshal(newRecommendation)
	if err != nil {
		log.Fatal(err)
	}

	backResp, err := http.Post("http://"+databaseAddress+"/userRecommendation/"+WEEK+"/"+UID, "application/json", bytes.NewBuffer(RecommendationJSON))
	if err != nil {
		http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
		log.Println("ERROR: PostUserProfileHandler() - Backend: " + err.Error())
		return
	}

	json.NewEncoder(w).Encode(backResp)

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
	//r.HandleFunc("/userInfo/{UID}/base", CalculateUserInfoHandler).Methods(http.MethodGet, http.MethodHead)
	//r.HandleFunc("/userInfo/{UID}", UpdateUserInfoHandler).Methods(http.MethodPost)
	r.HandleFunc("/_healthz", func(w http.ResponseWriter, _ *http.Request) { fmt.Fprint(w, "ok") })
	var handler http.Handler = r
	log.Printf("Backend listening at address " + listenAddress)
	log.Fatal(http.ListenAndServe(listenAddress, handler))
}
