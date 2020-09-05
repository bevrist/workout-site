package main

// this test requires mongoDB container to be running

import (
	"bytes"
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
	reqBody, _ := ioutil.ReadAll(req.Body)
	if strings.TrimSpace(string(reqBody)) != strings.TrimSpace(EXPECTED) {
		t.Errorf("Auth returned unexpected body: got %v \nwant %v", string(reqBody), EXPECTED)
		t.Fail()
	}
}

func TestGetUserInfo(t *testing.T) {
	req, err := http.Get("http://localhost/userInfo/testUID")
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.Fail()
	}
	// Check the response body is what we expect.
	EXPECTED := `{"UID":"testUID","FirstName":"Anthony","LastName":"Hanna","Weight":215,"WaistCirc":35.5,"HeightInches":75,"LeanBodyMass":15,"Age":20,"Gender":"male","Week":[{"Day":[{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalorie":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":40,"Carbs":40,"Protein":40,"TotalCalories":34,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":100,"Carbs":100,"Protein":100,"TotalCalories":300,"DayCalorie":"normal","Weight":321,"Cardio":"missed","WeightTraining":"no"}]},{"Day":[{"Fat":11,"Carbs":11,"Protein":11,"TotalCalories":31,"DayCalorie":"normal","Weight":222,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"}]},{"Day":[{"Fat":110,"Carbs":110,"Protein":110,"TotalCalories":310,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"}]}]}`
	reqBody, _ := ioutil.ReadAll(req.Body)
	if strings.TrimSpace(string(reqBody)) != strings.TrimSpace(EXPECTED) {
		t.Errorf("Auth returned unexpected body: got %v \nwant %v", string(reqBody), EXPECTED)
		t.Fail()
	}
}

func TestGetEmptyUserInfo(t *testing.T) {
	req, err := http.Get("http://localhost/userInfo/testUID2")
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.Fail()
	}
	// Check the response body is what we expect.
	EXPECTED := `null`
	reqBody, _ := ioutil.ReadAll(req.Body)
	if strings.TrimSpace(string(reqBody)) != strings.TrimSpace(EXPECTED) {
		t.Errorf("Auth returned unexpected body: got %v \nwant %v", string(reqBody), EXPECTED)
		t.Fail()
	}
}

func TestNewUserInfo(t *testing.T) {
	//do post request and verify success
	postBody := []byte(`{"UID":"testUID2","FirstName":"Test","LastName":"User","Weight":111,"WaistCirc":111.1,"HeightInches":111,"LeanBodyMass":111,"Age":111,"Gender":"female","Week":[{"Day":[{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalorie":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":40,"Carbs":40,"Protein":40,"TotalCalories":34,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":100,"Carbs":100,"Protein":100,"TotalCalories":300,"DayCalorie":"normal","Weight":321,"Cardio":"missed","WeightTraining":"no"}]}]}`)
	// post some data
	resp, err := http.Post("http://localhost/userInfo/testUID2", "application/json; charset=UTF-8", bytes.NewBuffer(postBody))
	if err != nil {
		t.Errorf("Post failed: %v", err)
		t.Fail()
	}
	respBody, _ := ioutil.ReadAll(resp.Body)
	if strings.TrimSpace(string(respBody)) != "ok" {
		t.Errorf("Auth returned unexpected body: got %v \nwant %v", string(respBody), "ok")
		t.Fail()
	}

	//validate post request
	req, err := http.Get("http://localhost/userInfo/testUID2")
	if err != nil {
		t.Errorf("Request failed: %v", err)
		t.Fail()
	}
	// Check the response body is what we expect.
	EXPECTED := `{"UID":"testUID2","FirstName":"Test","LastName":"User","Weight":111,"WaistCirc":111.1,"HeightInches":111,"LeanBodyMass":111,"Age":111,"Gender":"female","Week":[{"Day":[{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalorie":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":40,"Carbs":40,"Protein":40,"TotalCalories":34,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":100,"Carbs":100,"Protein":100,"TotalCalories":300,"DayCalorie":"normal","Weight":321,"Cardio":"missed","WeightTraining":"no"}]},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null}]}`
	reqBody, _ := ioutil.ReadAll(req.Body)
	if strings.TrimSpace(string(reqBody)) != strings.TrimSpace(EXPECTED) {
		t.Errorf("Auth returned unexpected body: got %v \nwant %v", string(reqBody), EXPECTED)
		t.Fail()
	}
}

func TestUpdateUserInfo(t *testing.T) {
	//do post request and verify success
	postBody := []byte(`{"UID":"testUID2","FirstName":"TestUpdate","LastName":"UpdateUser","Weight":222,"WaistCirc":222.2,"HeightInches":222,"LeanBodyMass":222,"Age":222,"Gender":"male","Week":[{"Day":[{"Fat":222,"Carbs":222,"Protein":222,"TotalCalories":222,"DayCalorie":"low","Weight":222,"Cardio":"missed","WeightTraining":"no"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalorie":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":40,"Carbs":40,"Protein":40,"TotalCalories":34,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":100,"Carbs":100,"Protein":100,"TotalCalories":300,"DayCalorie":"normal","Weight":321,"Cardio":"missed","WeightTraining":"no"}]}]}`)
	// post some data
	resp, err := http.Post("http://localhost/userInfo/testUID2", "application/json; charset=UTF-8", bytes.NewBuffer(postBody))
	if err != nil {
		t.Errorf("Post failed: %v", err)
		t.Fail()
	}
	respBody, _ := ioutil.ReadAll(resp.Body)
	if strings.TrimSpace(string(respBody)) != "ok" {
		t.Errorf("Auth returned unexpected body: got %v \nwant %v", string(respBody), "ok")
		t.Fail()
	}

	//validate post request
	req, err := http.Get("http://localhost/userInfo/testUID2")
	if err != nil {
		t.Errorf("Request failed: %v", err)
		t.Fail()
	}
	// Check the response body is what we expect.
	EXPECTED := `{"UID":"testUID2","FirstName":"TestUpdate","LastName":"UpdateUser","Weight":222,"WaistCirc":222.2,"HeightInches":222,"LeanBodyMass":222,"Age":222,"Gender":"male","Week":[{"Day":[{"Fat":222,"Carbs":222,"Protein":222,"TotalCalories":222,"DayCalorie":"low","Weight":222,"Cardio":"missed","WeightTraining":"no"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalorie":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":40,"Carbs":40,"Protein":40,"TotalCalories":34,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":100,"Carbs":100,"Protein":100,"TotalCalories":300,"DayCalorie":"normal","Weight":321,"Cardio":"missed","WeightTraining":"no"}]},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null}]}`
	reqBody, _ := ioutil.ReadAll(req.Body)
	if strings.TrimSpace(string(reqBody)) != strings.TrimSpace(EXPECTED) {
		t.Errorf("Auth returned unexpected body: got %v \nwant %v", string(reqBody), EXPECTED)
		t.Fail()
	}
}

func TestUpdateUserBaseline(t *testing.T) {
	//do post request and verify success
	postBody := []byte(`{"Day":[{"Fat":999,"Carbs":999,"Protein":999,"TotalCalories":999,"DayCalorie":"normal","Weight":999,"Cardio":"missed","WeightTraining":"yes"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalorie":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":40,"Carbs":40,"Protein":40,"TotalCalories":34,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":123,"Carbs":123,"Protein":123,"TotalCalories":123,"DayCalorie":"low","Weight":123,"Cardio":"missed","WeightTraining":"yes"}]}`)
	// post some data
	resp, err := http.Post("http://localhost/userBaseline/2/testUID2", "application/json; charset=UTF-8", bytes.NewBuffer(postBody))
	if err != nil {
		t.Errorf("Post failed: %v", err)
		t.Fail()
	}
	respBody, _ := ioutil.ReadAll(resp.Body)
	if strings.TrimSpace(string(respBody)) != "ok" {
		t.Errorf("Auth returned unexpected body: got %v \nwant %v", string(respBody), "ok")
		t.Fail()
	}

	//validate post request
	req, err := http.Get("http://localhost/userInfo/testUID2")
	if err != nil {
		t.Errorf("Request failed: %v", err)
		t.Fail()
	}
	// Check the response body is what we expect.
	EXPECTED := `{"UID":"testUID2","FirstName":"TestUpdate","LastName":"UpdateUser","Weight":222,"WaistCirc":222.2,"HeightInches":222,"LeanBodyMass":222,"Age":222,"Gender":"male","Week":[{"Day":[{"Fat":222,"Carbs":222,"Protein":222,"TotalCalories":222,"DayCalorie":"low","Weight":222,"Cardio":"missed","WeightTraining":"no"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalorie":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":40,"Carbs":40,"Protein":40,"TotalCalories":34,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":100,"Carbs":100,"Protein":100,"TotalCalories":300,"DayCalorie":"normal","Weight":321,"Cardio":"missed","WeightTraining":"no"}]},{"Day":null},{"Day":[{"Fat":999,"Carbs":999,"Protein":999,"TotalCalories":999,"DayCalorie":"normal","Weight":999,"Cardio":"missed","WeightTraining":"yes"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalorie":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":40,"Carbs":40,"Protein":40,"TotalCalories":34,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":123,"Carbs":123,"Protein":123,"TotalCalories":123,"DayCalorie":"low","Weight":123,"Cardio":"missed","WeightTraining":"yes"}]},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null}]}`
	reqBody, _ := ioutil.ReadAll(req.Body)
	if strings.TrimSpace(string(reqBody)) != strings.TrimSpace(EXPECTED) {
		t.Errorf("Auth returned unexpected body: got %v \nwant %v", string(reqBody), EXPECTED)
		t.Fail()
	}
}

func TestUpdateInvalidUserBaseline(t *testing.T) {
	//do post request and verify success
	postBody := []byte(`{"Day":[{"Fat":321,"Carbs":321,"Protein":321,"TotalCalories":321,"DayCalorie":"high","Weight":312,"Cardio":"missed","WeightTraining":"no"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalorie":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":40,"Carbs":40,"Protein":40,"TotalCalories":34,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":888,"Carbs":888,"Protein":888,"TotalCalories":888,"DayCalorie":"high","Weight":888,"Cardio":"missed","WeightTraining":"no"}]}`)
	// post some data
	resp, err := http.Post("http://localhost/userBaseline/23/testUID", "application/json; charset=UTF-8", bytes.NewBuffer(postBody))
	if err != nil {
		t.Errorf("Post failed: %v", err)
		t.Fail()
	}
	respBody, _ := ioutil.ReadAll(resp.Body)
	if strings.TrimSpace(string(respBody)) != "ok" {
		t.Errorf("Auth returned unexpected body: got %v \nwant %v", string(respBody), "ok")
		t.Fail()
	}

	//validate post request
	req, err := http.Get("http://localhost/userInfo/testUID")
	if err != nil {
		t.Errorf("Request failed: %v", err)
		t.Fail()
	}
	// Check the response body is what we expect.
	EXPECTED := `{"UID":"testUID","FirstName":"Anthony","LastName":"Hanna","Weight":215,"WaistCirc":35.5,"HeightInches":75,"LeanBodyMass":15,"Age":20,"Gender":"male","Week":[{"Day":[{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalorie":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":40,"Carbs":40,"Protein":40,"TotalCalories":34,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":100,"Carbs":100,"Protein":100,"TotalCalories":300,"DayCalorie":"normal","Weight":321,"Cardio":"missed","WeightTraining":"no"}]},{"Day":[{"Fat":11,"Carbs":11,"Protein":11,"TotalCalories":31,"DayCalorie":"normal","Weight":222,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"}]},{"Day":[{"Fat":110,"Carbs":110,"Protein":110,"TotalCalories":310,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"}]},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":[{"Fat":321,"Carbs":321,"Protein":321,"TotalCalories":321,"DayCalorie":"high","Weight":312,"Cardio":"missed","WeightTraining":"no"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalorie":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":40,"Carbs":40,"Protein":40,"TotalCalories":34,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":888,"Carbs":888,"Protein":888,"TotalCalories":888,"DayCalorie":"high","Weight":888,"Cardio":"missed","WeightTraining":"no"}]}]}`
	reqBody, _ := ioutil.ReadAll(req.Body)
	if strings.TrimSpace(string(reqBody)) != strings.TrimSpace(EXPECTED) {
		t.Errorf("Auth returned unexpected body: got %v \nwant %v", string(reqBody), EXPECTED)
		t.Fail()
	}
}
