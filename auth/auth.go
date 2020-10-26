// Auth provides an api for handling validation of session tokens and retrieving User ID's (UID)
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/gorilla/mux"
	"google.golang.org/api/option"

	structs "../common"
)

// Global Variables
var apiVersion string = "1.0" //the api version this service implements
var client *auth.Client       //firebase app instance
var useFirebase bool		  //debug flag for using firebase
// env
var listenAddress string //listen address of this service

//helper function for actually retrieving UID
func getUID(sessionToken string) string {
	//TEST hardcoded reply for testing purposes
	if sessionToken == "test" {
		return "testUID"
	} else if sessionToken == "testfail" {
		return ""
	}
	if useFirebase == false {
		return sessionToken
	}

	//validate session token and return UID, failure will return empty string
	token, err := client.VerifyIDTokenAndCheckRevoked(context.Background(), sessionToken)
	if err != nil {
		if err.Error() == "ID token has been revoked" {
			// Token is revoked. Inform the user to re-authenticate or signOut() the user.
			log.Println("Token revoked: " + sessionToken)
			return ""
		}
		// Token is invalid
		log.Println("Token Invalid: " + sessionToken)
		return ""
	}
	return token.UID
}

// GetUIDHandler validates session token and returns UID
func GetUIDHandler(w http.ResponseWriter, r *http.Request) {
	//extract session token from URL
	vars := mux.Vars(r)
	sessionToken := vars["SessionToken"]

	//validate sessionToken and get UID
	uid := getUID(sessionToken)
	isValid := true
	if uid == "" {
		isValid = false
	}

	//create auth struct
	auth := structs.Auth{IsValid: isValid, UID: uid}

	//marshal auth struct and respond to request
	out, err := json.Marshal(auth)
	if err != nil {
		log.Println("ERROR: Invalid syntax: " + err.Error())
		http.Error(w, "400 - invalid syntax.", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, string(out))
}

func main() {
	//populate environment variables
	listenAddress = os.Getenv("AUTH_LISTEN_ADDRESS")
	firebaseCredentials := os.Getenv("AUTH_FIREBASE_CREDENTIALS")
	//set default environment variables
	if listenAddress == "" {
		listenAddress = "0.0.0.0:8070"
	}

	useFirebase = true
	//initialize firebase app connection
	var opt option.ClientOption
	log.Println("AUTH_FIREBASE_CREDENTIALS = '" + firebaseCredentials + "'")
	if firebaseCredentials == "{}" {
		log.Println("Env AUTH_FIREBASE_CREDENTIALS empty, attempting to load from file...")
		opt = option.WithCredentialsFile("./workout-app-8b023-firebase-adminsdk-jh1ev-bbfc733122.json") //load credentials file
	} else if firebaseCredentials == "{test}" {
		useFirebase = false
	} else {
		opt = option.WithCredentialsJSON([]byte(firebaseCredentials))
	}

	if useFirebase == true {
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			log.Fatalf("error initializing firebase app: %v\n", err)
		}
		client, err = app.Auth(context.Background())
		if err != nil {
			log.Fatalf("error getting firebase Auth client: %v\n", err)
		}
	} else {
		log.Println("WARNING: Auth_Service not using firebase, all replies will be mirrored...")
	}

	//specify routes and start http server
	r := mux.NewRouter()
	r.HandleFunc("/apiVersion", func(w http.ResponseWriter, _ *http.Request) { fmt.Fprint(w, `{"apiVersion":`+apiVersion+"}") })
	r.HandleFunc("/getUID/{SessionToken}", GetUIDHandler).Methods(http.MethodGet, http.MethodHead)
	r.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) { fmt.Fprint(w, "ok") })
	var handler http.Handler = r
	log.Printf("Auth listening at address " + listenAddress)
	log.Fatal(http.ListenAndServe(listenAddress, handler))
}
