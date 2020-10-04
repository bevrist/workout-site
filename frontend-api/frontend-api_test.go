package main

import (
	// "bytes"
	// "encoding/json"
	"io/ioutil"
	"net/http"
	// "strings"
	"testing"
	"os"

	// structs "../common"
	// "github.com/google/go-cmp/cmp"
)

var frontendApiAddress string
// get service address from env
func TestMain(m *testing.M) {
	frontendApiAddress = os.Getenv("FRONTEND_API_SERVICE_ADDRESS")
	if frontendApiAddress == "" {
		frontendApiAddress = "localhost:8050"
	}
    os.Exit(m.Run())
}

func TestAPIVersion(t *testing.T) {
	req, err := http.Get("http://"+frontendApiAddress+"/apiVersion")
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.Fail()
	}
	// Check the response body is what we expect.
	expected := `{"apiVersion":1.0}`
	respBody, _ := ioutil.ReadAll(req.Body)
	if string(respBody) != expected {
		t.Errorf("Frontend-api returned unexpected body: got %v \nwant %v", string(respBody), expected)
		t.Fail()
	}
}

// func TestGetUserProfile(t *testing.T) {
// 	req, err := http.NewRequest("GET", "http://localhost/userProfile", nil)
// 	// set session token header for request
// 	req.Header.Set("Session-Token", "test")
// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		t.Errorf("Connection failed: %v", err)
// 		t.Fail()
// 	}
// 	// Check the response body is what we expect.
// 	expected := `{"FirstName":"Anthony","LastName":"Hannah","Weight":215,"WaistCirc":11,"HeightInches":72,"LeanBodyMass":15,"Age":27,"Gender":"female"}`
// 	respBody, _ := ioutil.ReadAll(resp.Body)
// 	if string(respBody) != expected {
// 		t.Errorf("Frontend-api returned unexpected body: got %v \nwant %v", string(respBody), expected)
// 		t.Fail()
// 	}
// }

// func TestGetUserBaseline(t *testing.T) {
// 	req, err := http.NewRequest("GET", "http://localhost/userBaseline", nil)
// 	// set session token header for request
// 	req.Header.Set("Session-Token", "test")
// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		t.Errorf("Connection failed: %v", err)
// 		t.Fail()
// 	}
// 	// Check the response body is what we expect.
// 	expected := `{"LowDay":2599,"NormalDay":2978,"HighDay":3357,"NFatRatio":0.25,"NCarbRatio":0.37,"NProteinRatio":0.38,"HFatRatio":0.3,"HCarbRatio":0.5,"HProteinRatio":0.2,"LFatRatio":0.41,"LCarbRatio":0.32,"LProteinRatio":0.27,"NFatAmount":83,"NCarbAmount":275,"NProteinAmount":283,"HFatAmount":112,"HCarbAmount":420,"HProteinAmount":168,"LFatAmount":118,"LCarbAmount":208,"LProteinAmount":175}`
// 	respBody, _ := ioutil.ReadAll(resp.Body)
// 	if string(respBody) != expected {
// 		t.Errorf("Frontend-api returned unexpected body: got %v \nwant %v", string(respBody), expected)
// 		t.Fail()
// 	}
// }
