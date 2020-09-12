package main

import (
	"io/ioutil"
	"net/http"
	"testing"
)

// func TestGetHtmlFiles(t *testing.T) {
// 	URL := "http://localhost"
// 	req, err := http.Get(URL)
// 	if err != nil {
// 		t.Errorf("Connection failed: %v", err)
// 		t.Fail()
// 	}
// 	if req.StatusCode != 200 {
// 		t.Errorf("Bad Response (not 200 OK) %v - %v on URL: %v", req.StatusCode, http.StatusText(req.StatusCode), URL)
// 		t.Fail()
// 	}

// 	URL = "http://localhost/auth"
// 	req, err = http.Get(URL)
// 	if req.StatusCode != 200 {
// 		t.Errorf("Bad Response (not 200 OK) %v - %v on URL: %v", req.StatusCode, http.StatusText(req.StatusCode), URL)
// 		t.Fail()
// 	}

// 	URL = "http://localhost/baseline"
// 	req, err = http.Get(URL)
// 	if req.StatusCode != 200 {
// 		t.Errorf("Bad Response (not 200 OK) %v - %v on URL: %v", req.StatusCode, http.StatusText(req.StatusCode), URL)
// 		t.Fail()
// 	}

// 	URL = "http://localhost/profile"
// 	req, err = http.Get(URL)
// 	if req.StatusCode != 200 {
// 		t.Errorf("Bad Response (not 200 OK) %v - %v on URL: %v", req.StatusCode, http.StatusText(req.StatusCode), URL)
// 		t.Fail()
// 	}

// 	URL = "http://localhost/weekly-tracking"
// 	req, err = http.Get(URL)
// 	if req.StatusCode != 200 {
// 		t.Errorf("Bad Response (not 200 OK) %v - %v on URL: %v", req.StatusCode, http.StatusText(req.StatusCode), URL)
// 		t.Fail()
// 	}
// }

func TestAPIVersion(t *testing.T) {
	req, err := http.Get("http://localhost/apiVersion")
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.Fail()
	}
	// Check the response body is what we expect.
	expected := `{"apiVersion":1.0}`
	respBody, _ := ioutil.ReadAll(req.Body)
	if string(respBody) != expected {
		t.Errorf("Auth returned unexpected body: got %v \nwant %v", string(respBody), expected)
		t.Fail()
	}
}

func TestGetUserProfile(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost/getUserProfile", nil)
	// set session token header for request
	req.Header.Set("Session-Token", "test")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.Fail()
	}
	// Check the response body is what we expect.
	expected := `{"FirstName":"Anthony","LastName":"Hannah","Weight":215,"WaistCirc":11,"HeightInches":72,"LeanBodyMass":15,"Age":27,"Gender":"female"}`
	respBody, _ := ioutil.ReadAll(resp.Body)
	if string(respBody) != expected {
		t.Errorf("Auth returned unexpected body: got %v \nwant %v", string(respBody), expected)
		t.Fail()
	}
}

func TestGetUserBaseline(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost/getUserBaseline", nil)
	// set session token header for request
	req.Header.Set("Session-Token", "test")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.Fail()
	}
	// Check the response body is what we expect.
	expected := `{"LowDay":2599,"NormalDay":2978,"HighDay":3357,"NFatRatio":0.25,"NCarbRatio":0.37,"NProteinRatio":0.38,"HFatRatio":0.3,"HCarbRatio":0.5,"HProteinRatio":0.2,"LFatRatio":0.41,"LCarbRatio":0.32,"LProteinRatio":0.27,"NFatAmount":83,"NCarbAmount":275,"NProteinAmount":283,"HFatAmount":112,"HCarbAmount":420,"HProteinAmount":168,"LFatAmount":118,"LCarbAmount":208,"LProteinAmount":175}`
	respBody, _ := ioutil.ReadAll(resp.Body)
	if string(respBody) != expected {
		t.Errorf("Auth returned unexpected body: got %v \nwant %v", string(respBody), expected)
		t.Fail()
	}
}
