package main

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestApiVersion(t *testing.T) {
	req, err := http.Get("http://localhost/apiVersion")
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
	req, err := http.Get("http://localhost/getUID/test")
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.Fail()
	}
	// Check the response body is what we expect.
	EXPECTED := `{"IsValid":true,"UID":"testUID"}`
	respBody, _ := ioutil.ReadAll(req.Body)
	if strings.TrimSpace(string(respBody)) != strings.TrimSpace(EXPECTED) {
		t.Errorf("Auth returned unexpected body: got %v \nwant %v", string(respBody), EXPECTED)
		t.Fail()
	}
}

func TestGetUIDFail(t *testing.T) {
	req, err := http.Get("http://localhost/getUID/testfail")
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.Fail()
	}
	// Check the response body is what we expect.
	EXPECTED := `{"IsValid":false,"UID":""}`
	respBody, _ := ioutil.ReadAll(req.Body)
	if strings.TrimSpace(string(respBody)) != strings.TrimSpace(EXPECTED) {
		t.Errorf("Auth returned unexpected body: got %v \nwant %v", string(respBody), EXPECTED)
		t.Fail()
	}
}
