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
		log.Fatal(err)
	}
	fmt.Fprintf(w, string(out))
	log.Println(string(out))
}

func main() {
	//populate environment variables
	listenAddress = os.Getenv("AUTH_LISTEN_ADDRESS")
	//set default environment variables
	if listenAddress == "" {
		listenAddress = "0.0.0.0:8070"
	}

	//initialize firebase app connection
	opt := option.WithCredentialsFile("./workout-app-8b023-firebase-adminsdk-jh1ev-bbfc733122.json") //load credentials file
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	client, err = app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	//specify routes and start http server
	r := mux.NewRouter()
	r.HandleFunc("/apiVersion", func(w http.ResponseWriter, _ *http.Request) { fmt.Fprint(w, `{"apiVersion":`+apiVersion+"}") })
	r.HandleFunc("/getUID/{SessionToken}", GetUIDHandler).Methods(http.MethodGet, http.MethodHead)
	r.HandleFunc("/_healthz", func(w http.ResponseWriter, _ *http.Request) { fmt.Fprint(w, "ok") })
	var handler http.Handler = r
	log.Printf("Auth listening at address " + listenAddress)
	log.Fatal(http.ListenAndServe(listenAddress, handler))
}
