package main

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetHtmlFiles(t *testing.T) {
	URL := "http://localhost"
	req, err := http.Get(URL)
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.Fail()
	}
	if req.StatusCode != 200 {
		t.Errorf("Bad Response (not 200 OK) %v - %v on URL: %v", req.StatusCode, http.StatusText(req.StatusCode), URL)
		t.Fail()
	}

	URL = "http://localhost/auth"
	req, err = http.Get(URL)
	if req.StatusCode != 200 {
		t.Errorf("Bad Response (not 200 OK) %v - %v on URL: %v", req.StatusCode, http.StatusText(req.StatusCode), URL)
		t.Fail()
	}

	URL = "http://localhost/baseline"
	req, err = http.Get(URL)
	if req.StatusCode != 200 {
		t.Errorf("Bad Response (not 200 OK) %v - %v on URL: %v", req.StatusCode, http.StatusText(req.StatusCode), URL)
		t.Fail()
	}

	URL = "http://localhost/profile"
	req, err = http.Get(URL)
	if req.StatusCode != 200 {
		t.Errorf("Bad Response (not 200 OK) %v - %v on URL: %v", req.StatusCode, http.StatusText(req.StatusCode), URL)
		t.Fail()
	}

	URL = "http://localhost/weekly-tracking"
	req, err = http.Get(URL)
	if req.StatusCode != 200 {
		t.Errorf("Bad Response (not 200 OK) %v - %v on URL: %v", req.StatusCode, http.StatusText(req.StatusCode), URL)
		t.Fail()
	}
}

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
