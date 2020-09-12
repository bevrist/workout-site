package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	structs "../common"
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
	req, err := http.Get("http://localhost/getUID/testfail")
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
