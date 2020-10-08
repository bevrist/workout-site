package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"log"
	"testing"

	structs "../common"
)

var authAddress string
// get service address from env
func TestMain(m *testing.M) {
	authAddress = os.Getenv("AUTH_SERVICE_ADDRESS")
	if authAddress == "" {
		authAddress = "localhost:8070"
	}
	log.Println("Testing Auth at address: " + authAddress)
    os.Exit(m.Run())
}

func TestApiVersion(t *testing.T) {
	req, err := http.Get("http://"+authAddress+"/apiVersion")
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.Fail()
	}
	// Check the response body is what we expect.
	EXPECTED := `{"apiVersion":1.0}`
	respBody, _ := ioutil.ReadAll(req.Body)
	if string(respBody) != EXPECTED {
		t.Errorf("Auth returned unexpected body: got %v \nwant %v", string(respBody), EXPECTED)
		t.Fail()
	}
}

func TestGetUID(t *testing.T) {
	req, err := http.Get("http://"+authAddress+"/getUID/test")
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.Fail()
	}
	//unmarshal response into struct
	reqBody, _ := ioutil.ReadAll(req.Body)
	var reqAuth, expectedAuth structs.Auth
	json.Unmarshal(reqBody, &reqAuth)
	//compare received struct with expected struct
	EXPECTED := []byte(`{"IsValid":true, "UID":"testUID"}`)
	json.Unmarshal(EXPECTED, &expectedAuth)
	if reqAuth != expectedAuth {
		t.Errorf("Auth returned unexpected body: \ngot  %+v \nwant %+v", reqAuth, expectedAuth)
		t.Fail()
	}
}

func TestGetUIDFail(t *testing.T) {
	req, err := http.Get("http://"+authAddress+"/getUID/testfail")
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.Fail()
	}
	//unmarshal response into struct
	reqBody, _ := ioutil.ReadAll(req.Body)
	var reqAuth, expectedAuth structs.Auth
	json.Unmarshal(reqBody, &reqAuth)
	//compare received struct with expected struct
	EXPECTED := []byte(`{"IsValid":false,"UID":""}`)
	json.Unmarshal(EXPECTED, &expectedAuth)
	if reqAuth != expectedAuth {
		t.Errorf("Auth returned unexpected body: \ngot  %+v \nwant %+v", reqAuth, expectedAuth)
		t.Fail()
	}
}
