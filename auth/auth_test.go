package main

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestApiVersion(t *testing.T) {
	req, err := http.Get("http://localhost/apiVersion")
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.Fail()
	}
	// Check the response body is what we expect.
	expected := `{"apiVersion":1.0}`
	respBody, _ := ioutil.ReadAll(req.Body)
	if string(respBody) != expected {
		t.Errorf("Auth returned unexpected body: got %v want %v", string(respBody), expected)
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
	expected := `{"IsValid":true,"UID":"testUID"}`
	respBody, _ := ioutil.ReadAll(req.Body)
	if string(respBody) != expected {
		t.Errorf("Auth returned unexpected body: got %v want %v", string(respBody), expected)
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
	expected := `{"IsValid":false,"UID":""}`
	respBody, _ := ioutil.ReadAll(req.Body)
	if string(respBody) != expected {
		t.Errorf("Auth returned unexpected body: got %v want %v", string(respBody), expected)
		t.Fail()
	}
}
