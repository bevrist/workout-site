package main

import (
	"context"
	"os"
	"time"

	structs "../common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Global Variables
var apiVersion string = "1.0" //the api version this service implements
// env
var listenAddress string //listen address of this service

var db *sql.DB

var uri string
var ctx context.Context
var client *mongo.Client
var workoutsitedb *mongo.Database
var clientsCollection *mongo.Collection

//GetUserInfoHandler returns json object of user data, or null if UID does not exist
func GetUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	//extract UID from URL
	vars := mux.Vars(r)
	UID := vars["UID"]

	// find document with matching UID
	filter := bson.M{"uid": UID}
	filterCursor, err := clientsCollection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	// store found document in client struct
	var client []structs.Client
	if err = filterCursor.All(ctx, &client); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(client)
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
	var client structs.Client
	err = json.Unmarshal(reqBody, &client)
	if err != nil {
		log.Fatal("UpdateUserInfoHandler: " + err.Error())
	}
	client.UID = UID

	//TODO: check if any user info in userInfo struct is nil and populate with existing user info

	//Check if UID exists in DB
	// Command to filter Document according to UID
	filter := bson.M{"uid": UID}

	// Query Documents that match uid from filter
	filterCursor, err := clientsCollection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	// Create splice to hold the data that was queried from the Find Command
	var finduid []structs.Client

	//Present ALL cases that match the filter and store them in client
	if err = filterCursor.All(ctx, &finduid); err != nil {
		log.Fatal(err)
	}

	//if no weekly user data exists, populate with baseline data
	if len(finduid) == 0 {
		log.Println("new doc created!")
		_, err := clientsCollection.InsertOne(ctx, bson.M{"uid": UID})
		if err != nil {
			log.Fatal(err)
		}

		//update filter with new doc with matching UID
		filterCursor, err = clientsCollection.Find(ctx, filter)
		if err != nil {
			log.Fatal(err)
		}
		var finduid []structs.Client
		if err = filterCursor.All(ctx, &finduid); err != nil {
			log.Fatal(err)
		}
	}

	// myBson, _ := bson.Marshal(&client)
	//yoBson := bson.D{
	//{Key: "$set", Value: client}}
	//}
	// := bson.M{"$set": bson.M{client}}
	//log.Println(a)
	// log.Println(string(a))
	//bson.MarshalAppend
	// clientsCollection.InsertOne(ctx, myBson)
	//options.Update().SetUpsert(true)
	result, _ := clientsCollection.UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: client}})
	log.Println(result)
	// json.NewEncoder(response).Encode(result)
	fmt.Fprint(w, "ok")
}

func main() {
	//populate environment variables
	listenAddress = os.Getenv("DATABASE_LISTEN_ADDRESS")
	databaseAddress := os.Getenv("DATABASE_ADDRESS")
	//set default environment variables
	if listenAddress == "" {
		listenAddress = "0.0.0.0:8050"
	}
	if databaseAddress == "" {
		databaseAddress = "34.94.73.0:27017"
	}

	uri = "mongodb://adminz:cheeksbutt@" + databaseAddress + "/workoutsite/?authSource=admin"

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
	r.HandleFunc("/userInfo/{UID}", UpdateUserInfoHandler).Methods(http.MethodPost)
	r.HandleFunc("/_healthz", func(w http.ResponseWriter, _ *http.Request) { fmt.Fprint(w, "ok") })
	var handler http.Handler = r
	log.Println("Backend listening at address " + listenAddress)
	log.Fatal(http.ListenAndServe(listenAddress, handler))
}
