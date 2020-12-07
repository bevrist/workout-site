package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"log"
	"testing"
	"time"

	structs "../common"

	"github.com/google/go-cmp/cmp"
)

var authAddress string
// get service address from env
func TestMain(m *testing.M) {
	authAddress = os.Getenv("AUTH_SERVICE_ADDRESS")
	if authAddress == "" {
		authAddress = "localhost:8070"
	}
	if (authAddress == "localhost:8070") {
		log.Println("Launching Auth for Local Test...")
		go main()
		time.Sleep(time.Second)
	}
	log.Println("Testing Auth at address: " + authAddress)
    os.Exit(m.Run())
}

func TestApiVersion(t *testing.T) {
	req, err := http.Get("http://"+authAddress+"/apiVersion")
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.FailNow()
	}
	// Check the response body is what we expect.
	EXPECTED := `{"apiVersion":1.0}`
	respBody, _ := ioutil.ReadAll(req.Body)
	if string(respBody) != EXPECTED {
		t.Errorf("Auth returned unexpected body: got %v \nwant %v", string(respBody), EXPECTED)
		t.FailNow()
	}
}

func TestGetUID(t *testing.T) {
	req, err := http.Get("http://"+authAddress+"/getUID/test")
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.FailNow()
	}
	//unmarshal response into struct
	reqBody, _ := ioutil.ReadAll(req.Body)
	var reqAuth, expectedAuth structs.Auth
	json.Unmarshal(reqBody, &reqAuth)
	//compare received struct with expected struct
	EXPECTED := []byte(`{"IsValid":true,"UID":"testUID","IsAdmin":true}`)
	json.Unmarshal(EXPECTED, &expectedAuth)
	if !cmp.Equal(reqAuth, expectedAuth) {
		t.Errorf("Database returned unexpected body: \ngot -: %+v \nwant -: %+v", string(reqBody), string(EXPECTED))
		t.FailNow()
	}
}

func TestGetUIDFail(t *testing.T) {
	req, err := http.Get("http://"+authAddress+"/getUID/testfail")
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.FailNow()
	}
	//unmarshal response into struct
	reqBody, _ := ioutil.ReadAll(req.Body)
	var reqAuth, expectedAuth structs.Auth
	json.Unmarshal(reqBody, &reqAuth)
	//compare received struct with expected struct
	EXPECTED := []byte(`{"IsValid":false,"UID":"","IsAdmin":false}`)
	json.Unmarshal(EXPECTED, &expectedAuth)
	if !cmp.Equal(reqAuth, expectedAuth) {
		t.Errorf("Database returned unexpected body: \ngot -: %+v \nwant -: %+v", string(reqBody), string(EXPECTED))
		t.FailNow()
	}
}

// test for user that does not have admin status
func TestGetNonAdminStatus(t *testing.T) {
	req, err := http.Get("http://"+authAddress+"/getUID/test2")
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.FailNow()
	}
	//unmarshal response into struct
	reqBody, _ := ioutil.ReadAll(req.Body)
	var reqAuth, expectedAuth structs.Auth
	json.Unmarshal(reqBody, &reqAuth)
	//compare received struct with expected struct
	EXPECTED := []byte(`{"IsValid":true,"UID":"test2","IsAdmin":false}`)
	json.Unmarshal(EXPECTED, &expectedAuth)
	if !cmp.Equal(reqAuth, expectedAuth) {
		t.Errorf("Database returned unexpected body: \ngot -: %+v \nwant -: %+v", string(reqBody), string(EXPECTED))
		t.FailNow()
	}
}
