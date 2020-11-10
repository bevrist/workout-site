package main

import (
	"log"
	"net/http"
	"os"
	"testing"
)

var frontendWebAddress string

// get service address from env
func TestMain(m *testing.M) {
	frontendWebAddress = os.Getenv("FRONTEND_WEB_ADDRESS")
	if frontendWebAddress == "" {
		frontendWebAddress = "http://localhost"
	}
	log.Println("Testing Frontend-Web at address: " + frontendWebAddress)
	os.Exit(m.Run())
}

func TestGetHtmlFiles(t *testing.T) {
	URL := frontendWebAddress
	req, err := http.Get(URL)
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.Fail()
	}
	if req.StatusCode != 200 {
		t.Errorf("Bad Response (not 200 OK) %v - %v on URL: %v", req.StatusCode, http.StatusText(req.StatusCode), URL)
		t.Fail()
	}

	URL = frontendWebAddress + "/auth"
	req, err = http.Get(URL)
	if req.StatusCode != 200 {
		t.Errorf("Bad Response (not 200 OK) %v - %v on URL: %v", req.StatusCode, http.StatusText(req.StatusCode), URL)
		t.Fail()
	}

	URL = frontendWebAddress + "/history"
	req, err = http.Get(URL)
	if req.StatusCode != 200 {
		t.Errorf("Bad Response (not 200 OK) %v - %v on URL: %v", req.StatusCode, http.StatusText(req.StatusCode), URL)
		t.Fail()
	}

	URL = frontendWebAddress + "/profile"
	req, err = http.Get(URL)
	if req.StatusCode != 200 {
		t.Errorf("Bad Response (not 200 OK) %v - %v on URL: %v", req.StatusCode, http.StatusText(req.StatusCode), URL)
		t.Fail()
	}

	URL = frontendWebAddress + "/daily-update"
	req, err = http.Get(URL)
	if req.StatusCode != 200 {
		t.Errorf("Bad Response (not 200 OK) %v - %v on URL: %v", req.StatusCode, http.StatusText(req.StatusCode), URL)
		t.Fail()
	}
}
