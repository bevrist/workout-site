// Auth provides an api for handling validation of session tokens and retrieving User ID's (UID)

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"

	structs "local/common"
)

// Global Variables
var apiVersion string = "1.0" //the api version this service implements
// list of admin UID's
var Admins []string = []string{"testUID", "test3", "ADMIN-UIDS-HERE"}
var test = false

// var client *auth.Client //firebase app instance
// var useFirebase bool    //debug flag for using firebase
var ctx = context.Background()
var rdb = &redis.Client{}

// GetUIDHandler validates api session token and returns UID
func GetUIDHandler(res http.ResponseWriter, req *http.Request) {
	//extract session token from URL
	vars := mux.Vars(req)
	sessionToken := vars["SessionToken"]

	//validate sessionToken and get UID
	// uid := getUID(sessionToken)
	uid := validateSessionToken(sessionToken)
	isValid := true
	if uid == "" {
		isValid = false
	}
	//check if user is on admins list
	isAdmin := false
	for _, admin := range Admins {
		if admin == uid {
			isAdmin = true
			break
		}
	}

	//create auth struct
	auth := structs.Auth{IsValid: isValid, UID: uid, IsAdmin: isAdmin}

	//marshal auth struct and respond to request
	out, err := json.Marshal(auth)
	if err != nil {
		log.Println("ERROR: Invalid syntax: " + err.Error())
		http.Error(res, "400 - invalid syntax.", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(res, string(out))
}

// returns redis uid of valid token, else returns empty string
func validateSessionToken(token string) string {
	//TEST hardcoded reply for testing purposes
	if token == "test" {
		return "testUID"
	} else if token == "testfail" {
		return ""
	} else if test == true {
		return token
	}
	// regex confirm token is valid (no funny or extra characters)
	reg := regexp.MustCompile(`^[a-zA-Z0-9]{10,}$`)
	if reg.MatchString(token) == false {
		return ""
	}
	// if token doesnt exist in db, this will return ""
	val, err := rdb.Get(ctx, token).Result()
	if err != nil {
		log.Println("WARN: get token fail: " + err.Error())
	}
	return val
}

// extract user info from oauth, generate session token, store in redis, store cookie, redirect to /daily-update
func completeLogin(gothUser goth.User, res http.ResponseWriter) {
	newSessionToken := randSeq(64)
	// set token which expires in one day
	err := rdb.Set(ctx, newSessionToken, gothUser.Email, time.Hour*24).Err()
	if err != nil {
		log.Println("ERROR: Setting Session Token Failed: " + gothUser.Email + "; " + err.Error())
	}
	res.Header().Set("Set-Cookie", "Authorization="+newSessionToken+"; HttpOnly; SameSite=Lax")
	res.Header().Set("Location", "/daily-update")
	res.WriteHeader(http.StatusTemporaryRedirect)
	fmt.Fprint(res, `<h2>Login Successful!</h2><br><a href="/daily-update">Go to Daily Update</a>`)
}

// handles logging out of application, delete session token, delete cookie, redirect home
func authLogoutHandler(res http.ResponseWriter, req *http.Request) {
	gothic.Logout(res, req)
	err := rdb.Del(ctx, req.Header.Get("Authorization")).Err()
	if err != nil {
		log.Println("ERROR: Deleting Session Token Failed: " + err.Error())
	}
	res.Header().Set("Set-Cookie", "Authorization=0; HttpOnly; expires=Thu, 01 Jan 1970 00:00:00 GMT")
	res.Header().Set("Location", "/")
	res.WriteHeader(http.StatusTemporaryRedirect)
}

//get user info from login provider callback request
func authCallbackHandler(res http.ResponseWriter, req *http.Request) {
	gothUser, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		fmt.Fprintln(res, err)
		return
	}
	completeLogin(gothUser, res)
}

// realistically this just redirects user to oauth callback
func authProviderHandler(res http.ResponseWriter, req *http.Request) {
	// try to get the user without re-authenticating (fails, no goth session exists)
	if gothUser, err := gothic.CompleteUserAuth(res, req); err == nil {
		completeLogin(gothUser, res)
	} else {
		gothic.BeginAuthHandler(res, req)
	}
}

//show login page
func authLoginHandler(res http.ResponseWriter, req *http.Request) {
	// if user auth token is invalid, delete token from browser
	if validateSessionToken(req.Header.Get("Authorization")) == "" {
		res.Header().Set("Set-Cookie", "Authorization=0; HttpOnly; expires=Thu, 01 Jan 1970 00:00:00 GMT")
	} else { //redirect to /daily-update if valid token
		res.Header().Set("Location", "/")
		res.WriteHeader(http.StatusTemporaryRedirect)
	}
	// TODO: swap this "page" with html file embedded
	fmt.Fprint(res, `<h1>Login</h1><a href="/auth/github">login with github</a><br><a href="/auth/google">login with google</a>`)
}

//generate random string of n length
func randSeq(n int) string {
	var characters = []rune(`abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789`)
	b := make([]rune, n)
	for i := range b {
		b[i] = characters[rand.Intn(len(characters))]
	}
	return string(b)
}

func main() {
	//populate environment variables
	listenAddress := os.Getenv("AUTH_LISTEN_ADDRESS")
	redisConnectionString := os.Getenv("REDIS_CONNECTION_STRING")
	testEnv := os.Getenv("TEST")
	//set default values for missing env vars
	if listenAddress == "" {
		listenAddress = "0.0.0.0:8070"
	}
	if redisConnectionString == "" {
		// redisConnectionString = "redis://user:pass:@localhost:6379/0"
		redisConnectionString = "redis://localhost:6379/0"
	}
	if strings.ToLower(testEnv) == "true" || testEnv == "1" {
		log.Println("WARN: TEST MODE ENABLED!")
		test = true
	}

	//die if provider key and secret not provided, instantiate goth providers
	if test != true {
		if os.Getenv("PROVIDER_KEY") == "" {
			log.Fatalln("FATAL: PROVIDER_KEY not set.")
		}
		if os.Getenv("PROVIDER_SECRET") == "" {
			log.Fatalln("FATAL: PROVIDER_SECRET not set.")
		}
		// instantiate goth providers
		goth.UseProviders(
			// google.New(os.Getenv("PROVIDER_KEY"), os.Getenv("PROVIDER_SECRET"), "http://localhost:3000/auth/google/callback"),
			github.New(os.Getenv("PROVIDER_KEY"), os.Getenv("PROVIDER_SECRET"), "http://localhost:3000/auth/github/callback"),
		)
	}

	//create new redis client
	opt, err := redis.ParseURL(redisConnectionString)
	if err != nil {
		panic(err)
	}
	rdb = redis.NewClient(opt)

	//specify routes and start http server
	var router = mux.NewRouter()
	var handler http.Handler = router
	var server = http.Server{Addr: listenAddress, Handler: handler}
	router.HandleFunc("/apiVersion", func(res http.ResponseWriter, _ *http.Request) { fmt.Fprint(res, `{"apiVersion":`+apiVersion+"}") })
	router.HandleFunc("/getUID/{SessionToken}", GetUIDHandler).Methods(http.MethodGet, http.MethodHead)
	router.HandleFunc("/auth/{provider}/callback", authCallbackHandler).Methods(http.MethodGet, http.MethodHead)
	router.HandleFunc("/auth/logout/{provider}", authLogoutHandler).Methods(http.MethodGet, http.MethodHead)
	router.HandleFunc("/auth/logout", func(res http.ResponseWriter, _ *http.Request) {
		res.Header().Set("Location", "/auth/logout/github")
		res.WriteHeader(http.StatusTemporaryRedirect)
	}).Methods(http.MethodGet, http.MethodHead)
	router.HandleFunc("/auth/{provider}", authProviderHandler).Methods(http.MethodGet, http.MethodHead)
	router.HandleFunc("/auth", authLoginHandler).Methods(http.MethodGet, http.MethodHead)
	router.HandleFunc("/healthz", func(res http.ResponseWriter, _ *http.Request) { fmt.Fprint(res, "ok") })
	log.Printf("Auth listening at address " + listenAddress)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
