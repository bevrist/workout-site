package main

import (
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
)

// Global Variables
var apiVersion string = "1.0" //the api version this service implements

// TODO setup listen address and other variables through env

var db *sql.DB

//GetUserInfoHandler returns json object of user data
func GetUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	//extract UID from URL
	vars := mux.Vars(r)
	UID := vars["UID"]

	//load user data from database by their UID
	sqlStatement := `SELECT first_name, last_name, weight, waistcirc, heightinches, leanbodymass FROM client WHERE uid=$1;`
	var firstName, lastName string
	var weight, waistCirc, heightInches, leanBodyMass int

	row := *db.QueryRow(sqlStatement, UID)
	switch err := row.Scan(&firstName, &lastName, &weight, &waistCirc, &heightInches, &leanBodyMass); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		//shit worked
	default:
		log.Fatal(err)
	}

	//respond with JSON object
	response := structs.UserInfo{FirstName: firstName, LastName: lastName, Weight: weight, WaistCirc: waistCirc, HeightInches: heightInches, LeanBodyMass: leanBodyMass}
	json.NewEncoder(w).Encode(response)
}

//UpdateUserInfoHandler updates user info from POST data
func UpdateUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	// extract UID from URL
	vars := mux.Vars(r)
	UID := vars["UID"]

	// unmarshal the body of POST request as a Client struct
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var userInfo structs.UserInfo
	err = json.Unmarshal(reqBody, &userInfo)
	if err != nil {
		log.Fatal(err)
	}

	//TODO write commentz
	sqlStatement := `SELECT uid FROM client WHERE uid=$1;`
	var uid string

	row := *db.QueryRow(sqlStatement, UID)
	switch err := row.Scan(&uid); err {
	case sql.ErrNoRows:
		//MAKE NEW USER (sql insert)
	case nil:
		//UPDATE USER INFO (sql update)
	default:
		log.Fatal(err)
	}

	//respond to POST and redirect to previous page
	fmt.Fprintf(w, "<!DOCTYPE html>SUCCESS <script>window.history.back();</script>")

	//print received data to output //FIXME remove
	out, _ := json.Marshal(userInfo)
	log.Printf("%v\n", string(out))
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully connected to DB")

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
