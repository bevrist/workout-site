package main

import (
	structs "../common"

	"encoding/json"
	"fmt"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"

	"google.golang.org/api/option"
)

// Global Variables
var apiVersion string = "1.0" //the api version this service implements
var client *auth.Client       //firebase app instance

// TODO setup listen address and other variables through env

// GetUIDHandler validates session token and returns UID
func GetUIDHandler(w http.ResponseWriter, r *http.Request) {
	//extract session token from URL
	vars := mux.Vars(r)
	sessionToken := vars["SessionToken"]

	//validate session token and get UID
	var IsValid bool = true
	var UID string = ""
	token, err := client.VerifyIDTokenAndCheckRevoked(context.Background(), sessionToken)
	if err != nil {
		if err.Error() == "ID token has been revoked" {
			// Token is revoked. Inform the user to reauthenticate or signOut() the user.
			IsValid = false
		} else {
			// Token is invalid
			IsValid = false
		}
	} else {
		UID = token.UID
	}

	//create auth struct
	auth := structs.Auth{IsValid, UID}

	//marshal auth struct and respond to request
	out, err := json.Marshal(auth)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, string(out))
	log.Println(string(out))
}

func main() {
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
	r.HandleFunc("/apiVersion", func(w http.ResponseWriter, _ *http.Request) { fmt.Fprint(w, "{\"apiVersion\":"+apiVersion+"}") })
	r.HandleFunc("/getUID/{SessionToken}", GetUIDHandler).Methods(http.MethodGet, http.MethodHead)
	r.HandleFunc("/_healthz", func(w http.ResponseWriter, _ *http.Request) { fmt.Fprint(w, "ok") })
	var handler http.Handler = r
	log.Printf("Auth listening at address :8070")
	log.Fatal(http.ListenAndServe(":8070", handler))
}
