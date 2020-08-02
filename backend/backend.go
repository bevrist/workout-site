package main

import (
	"math"
	"os"

	structs "../common"

	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host     = "anthonyhanna.com"
	port     = 5432
	user     = "theuser"
	password = "cyber@jb122"
	dbname   = "workoutsite"
) //FIXME move this to either env vars or config file

// TODO add carbs fat and protein /w ratio
type Calculations struct {
	LowDay         float64
	NormalDay      float64
	HighDay        float64
	NFatRatio      float64
	NCarbRatio     float64
	NProteinRatio  float64
	HFatRatio      float64
	HCarbRatio     float64
	HProteinRatio  float64
	LFatRatio      float64
	LCarbRatio     float64
	LProteinRatio  float64
	NFatAmount     float64
	NCarbAmount    float64
	NProteinAmount float64
	HFatAmount     float64
	HCarbAmount    float64
	HProteinAmount float64
	LFatAmount     float64
	LCarbAmount    float64
	LProteinAmount float64
}

// Global Variables
var apiVersion string = "1.0" //the api version this service implements
// env
var listenAddress string //listen address of this service

var db *sql.DB

//GetUserInfoHandler returns json object of user data
func GetUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	//extract UID from URL
	vars := mux.Vars(r)
	UID := vars["UID"]

	//load user data from database by their UID
	sqlStatement := `SELECT first_name, last_name, weight, waistcirc, heightinches, leanbodymass, age, gender FROM client WHERE uid=$1;`
	var firstName, lastName, gender string
	var weight, waistCirc, heightInches, leanBodyMass, age int

	row := db.QueryRow(sqlStatement, UID)
	switch err := row.Scan(&firstName, &lastName, &weight, &waistCirc, &heightInches, &leanBodyMass, &age, &gender); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		//shit worked
	default:
		log.Fatal("GetUserInfoHandler: " + err.Error())
	}

	var bmr float64
	// TODO Change float math to round
	if gender == "male" {
		bmr = 66 + (6.3 * float64(weight)) + (12.9 * float64(heightInches)) - (6.8 * float64(age))
	} else {
		bmr = 655 + (4.3 * float64(weight)) + (4.7 * float64(heightInches)) - (4.7 * float64(age))
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

	NFatAmount := math.Round((normalday * NFatRatio) / 9)
	NCarbAmount := math.Round((normalday * NCarbRatio) / 4)
	NProteinAmount := math.Round((normalday * NProteinRatio) / 4)
	HFatAmount := math.Round((highday * HFatRatio) / 9)
	HCarbAmount := math.Round((highday * HCarbRatio) / 4)
	HProteinAmount := math.Round((highday * HProteinRatio) / 4)
	LFatAmount := math.Round((lowday * LFatRatio) / 9)
	LCarbAmount := math.Round((lowday * LCarbRatio) / 4)
	LProteinAmount := math.Round((lowday * LProteinRatio) / 4)

	//log.Println(bmr, lowday, normalday, highday, NFatRatio, NCarbRatio, NProteinRatio, HFatRatio, HCarbRatio, HProteinRatio, LFatRatio, LCarbRatio, LProteinRatio,
	//NFatAmount, NCarbAmount, NProteinAmount, HFatAmount, HCarbAmount, HProteinAmount, LFatAmount, LCarbAmount, LProteinAmount)
	//respond with JSON object
	response := structs.UserInfo{FirstName: firstName, LastName: lastName, Weight: weight, WaistCirc: waistCirc, HeightInches: heightInches, LeanBodyMass: leanBodyMass, Age: age, Gender: gender}

	calc := Calculations{LowDay: lowday, NormalDay: normalday, HighDay: highday, NFatRatio: NFatRatio, NCarbRatio: NCarbRatio, NProteinRatio: NProteinRatio, HFatRatio: HFatRatio, HCarbRatio: HCarbRatio,
		HProteinRatio: HProteinRatio, LFatRatio: LFatRatio, LCarbRatio: LCarbRatio, LProteinRatio: LProteinRatio,
		NFatAmount: NFatAmount, NCarbAmount: NCarbAmount, NProteinAmount: NProteinAmount, HFatAmount: HFatAmount, HCarbAmount: HCarbAmount, HProteinAmount: HProteinAmount,
		LFatAmount: LFatAmount, LCarbAmount: LCarbAmount, LProteinAmount: LProteinAmount}

	json.NewEncoder(w).Encode(response)
	json.NewEncoder(w).Encode(calc)
}

//UpdateUserInfoHandler updates user info from POST data
func UpdateUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	// extract UID from URL
	vars := mux.Vars(r)
	UID := vars["UID"]

	// unmarshal the body of POST request as a Client struct
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("UpdateUserInfoHandler: " + err.Error())
	}
	var userInfo structs.UserInfo
	err = json.Unmarshal(reqBody, &userInfo)
	if err != nil {
		log.Fatal("UpdateUserInfoHandler: " + err.Error())
	}

	//TODO: check if any user info in userInfo struct is nil and populate with existing user info

	//Check if UID exists in DB
	sqlStatement := `SELECT uid FROM client WHERE uid=$1;`
	var uid string
	row := db.QueryRow(sqlStatement, UID)
	switch err := row.Scan(&uid); err {
	case sql.ErrNoRows:
		//if UID not found: MAKE NEW USER (sql insert)
		sqlInsertStatement := `INSERT INTO client ("uid", "first_name", "last_name", "weight", "waistcirc", "heightinches", "leanbodymass", "age", "gender" ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);`
		_, err := db.Exec(sqlInsertStatement, UID, userInfo.FirstName, userInfo.LastName, userInfo.Weight, userInfo.WaistCirc, userInfo.HeightInches, userInfo.LeanBodyMass, userInfo.Age, userInfo.Gender)
		if err != nil {
			log.Fatal("UpdateUserInfoHandler: " + err.Error())
		}
	case nil:
		//if UID found: UPDATE USER INFO (sql update)
		sqlInsertStatement := `UPDATE client SET  "first_name" = COALESCE(NULLIF($2,''), first_name) , "last_name" = COALESCE(NULLIF($3,''), last_name), "weight" = COALESCE(NULLIF($4,0), weight), 
		"waistcirc" = COALESCE(NULLIF($5,0), waistcirc), "heightinches" = COALESCE(NULLIF($6,0), heightinches), "leanbodymass" = COALESCE(NULLIF($7,0), leanbodymass),
		 "age" = COALESCE(NULLIF($8,0), age, "gender" = COALESCE(NULLIF($9,''), gender) WHERE uid=$1;`
		_, err := db.Exec(sqlInsertStatement, UID, userInfo.FirstName, userInfo.LastName, userInfo.Weight, userInfo.WaistCirc, userInfo.HeightInches, userInfo.LeanBodyMass, userInfo.Age, userInfo.Gender)
		if err != nil {
			log.Fatal("UpdateUserInfoHandler: " + err.Error())
		}
	default:
		log.Fatal("UpdateUserInfoHandler: " + err.Error())
	}

	//respond to POST and redirect to previous page
	fmt.Fprintf(w, "<!DOCTYPE html>SUCCESS <script>window.history.back();</script>")
	//FIXME: change this to respond with 200 success code

	//print received data to output //FIXME remove
	out, _ := json.Marshal(userInfo)
	log.Printf("%v\n", string(out))
}

func main() {
	//populate environment variables
	listenAddress = os.Getenv("BACKEND_LISTEN_ADDRESS")
	//set default environment variables
	if listenAddress == "" {
		listenAddress = "0.0.0.0:8090"
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("main: " + err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("main: " + err.Error())
	}
	log.Println("Successfully connected to DB")

	//specify routes and start http server
	r := mux.NewRouter()
	r.HandleFunc("/apiVersion", func(w http.ResponseWriter, _ *http.Request) { fmt.Fprint(w, "{\"apiVersion\":"+apiVersion+"}") })
	r.HandleFunc("/userInfo/{UID}", GetUserInfoHandler).Methods(http.MethodGet, http.MethodHead)
	r.HandleFunc("/userInfo/{UID}", UpdateUserInfoHandler).Methods(http.MethodPost)
	r.HandleFunc("/_healthz", func(w http.ResponseWriter, _ *http.Request) { fmt.Fprint(w, "ok") })
	var handler http.Handler = r
	log.Printf("Backend listening at address " + listenAddress)
	log.Fatal(http.ListenAndServe(listenAddress, handler))
}
