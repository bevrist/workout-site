package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	structs "../common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/gorilla/mux"
)

// Global Variables
var apiVersion string = "1.0" //the api version this service implements
// env
var listenAddress string //listen address of this service

var uri string
var ctx context.Context
var client *mongo.Client
var workoutsitedb *mongo.Database
var clientsCollection *mongo.Collection

//getUserData helper function which gets user collection from database
func getUserCollection(UID string) []structs.Client {
	// find document with matching UID
	filter := bson.M{"uid": UID}
	filterCursor, err := clientsCollection.Find(ctx, filter)
	if err != nil {
		log.Fatal("ERROR: getUserCollection: " + err.Error())
	}
	// store found document in client struct
	var client []structs.Client
	if err = filterCursor.All(ctx, &client); err != nil {
		log.Fatal("ERROR: getUserCollection: " + err.Error())
	}
	return client
}

//GetUserInfoHandler returns json object of user data, or null if UID does not exist
func GetUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	//extract UID from URL
	vars := mux.Vars(r)
	UID := vars["UID"]
	//get user collection from database
	client := getUserCollection(UID)
	//return "null" if getUserCollection is empty
	if len(client) == 0 {
		fmt.Fprintf(w, "null")
		return
	}
	//return data to request
	json.NewEncoder(w).Encode(client[0])
}

//updateUserDocument helper function which uploads Client struct to database
func updateUserDocument(client structs.Client) {
	opts := options.Update().SetUpsert(true) //update or insert document
	filter := bson.M{"uid": client.UID}      //filter according to UID
	_, err := clientsCollection.UpdateOne(context.TODO(), filter, bson.D{{Key: "$set", Value: client}}, opts)
	if err != nil {
		log.Fatal("ERROR: updateUserDocument: " + err.Error())
	}
}

//UpdateUserProfileHandler updates user profile info from a Client object
func UpdateUserProfileHandler(w http.ResponseWriter, r *http.Request) {
	// extract UID from URL
	vars := mux.Vars(r)
	UID := vars["UID"]
	// unmarshal the body of POST request as a Client struct
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("ERROR: UpdateUserProfileHandler: " + err.Error())
	}
	var client structs.Client
	err = json.Unmarshal(reqBody, &client)
	if err != nil {
		log.Fatal("ERROR: UpdateUserProfileHandler: " + err.Error())
	}
	client.UID = UID

	//update profile information in server
	updateUserDocument(client)

	fmt.Fprint(w, "ok")
}

//UpdateUserBaselineHandler updates user profile info from a Week object
func UpdateUserBaselineHandler(w http.ResponseWriter, r *http.Request) {
	// extract UID and weekToUpdate from URL
	vars := mux.Vars(r)
	UID := vars["UID"]
	weekToUpdate, _ := strconv.Atoi(vars["week"])
	// unmarshal the body of POST request as a Week struct
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("ERROR: UpdateUserBaselineHandler: " + err.Error())
	}
	var clientWeek structs.Week
	err = json.Unmarshal(reqBody, &clientWeek)
	if err != nil {
		log.Fatal("ERROR: UpdateUserBaselineHandler: " + err.Error())
	}
	//get existing user data from db
	userData := getUserCollection(UID)
	if len(userData) == 0 {
		log.Println("ERROR: UpdateUserBaselineHandler: getUserCollection is empty")
		http.Error(w, "update failed, UID does not exist.", http.StatusNotAcceptable)
		return
	}

	//ensure client week array has correct capacity, fill with null week objects if necessary
	for len(userData[0].Week) < 24 {
		userData[0].Week = append(userData[0].Week, structs.Week{})
	}

	//update specific week in existing user data from database
	userData[0].Week[weekToUpdate] = clientWeek
	//update user information in server
	updateUserDocument(userData[0])

	fmt.Fprint(w, "ok")
}

func main() {
	//populate environment variables
	listenAddress = os.Getenv("DATABASE_LISTEN_ADDRESS")
	databaseAddress := os.Getenv("DATABASE_ADDRESS")
	databaseUsername := os.Getenv("DATABASE_USERNAME")
	databasePassword := os.Getenv("DATABASE_PASSWORD")
	//set default environment variables
	if listenAddress == "" {
		listenAddress = "0.0.0.0:8050"
	}
	if databaseAddress == "" {
		databaseAddress = "localhost:27017"
	}
	if databaseUsername == "" {
		databaseUsername = "adminz"
	}
	if databasePassword == "" {
		databasePassword = "cheeksbutt"
	}

	//connect to mongoDB
	uri = "mongodb://" + databaseUsername + ":" + databasePassword + "@" + databaseAddress + "/workoutsite/?authSource=admin"
	var cancel context.CancelFunc
	var err error
	ctx, cancel = context.WithTimeout(context.Background(), 999999999*time.Second)
	defer cancel()
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("ERROR: Creating MongoDB connection failed: " + err.Error())
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal("ERROR: Disconnect from MongoDB failed: " + err.Error())
		}
	}()
	// Ping the primary
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal("ERROR: Connection to MongoDB failed: " + err.Error())
	}
	log.Printf("Connected to Database at: " + databaseAddress)
	workoutsitedb = client.Database("workoutsite")
	clientsCollection = workoutsitedb.Collection("clients")

	//specify routes and start http server
	r := mux.NewRouter()
	r.HandleFunc("/apiVersion", func(w http.ResponseWriter, _ *http.Request) { fmt.Fprint(w, "{\"apiVersion\":"+apiVersion+"}") })
	r.HandleFunc("/userInfo/{UID}", GetUserInfoHandler).Methods(http.MethodGet, http.MethodHead)
	r.HandleFunc("/userInfo/{UID}", UpdateUserProfileHandler).Methods(http.MethodPost)
	r.HandleFunc("/userBaseline/{week}/{UID}", UpdateUserBaselineHandler).Methods(http.MethodPost)
	r.HandleFunc("/_healthz", func(w http.ResponseWriter, _ *http.Request) { fmt.Fprint(w, "ok") })
	var handler http.Handler = r
	log.Println("Database listening at address " + listenAddress)
	log.Fatal("fail: ", http.ListenAndServe(listenAddress, handler))
}
