package main

// this test requires mongoDB container to be running

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

func TestGetUserInfo(t *testing.T) {
	req, err := http.Get("http://localhost/userInfo/testUID")
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.Fail()
	}
	// Check the response body is what we expect.
	EXPECTED := `[{"UID":"testUID","FirstName":"Anthony","LastName":"Hanna","Weight":215,"WaistCirc":35.5,"HeightInches":75,"LeanBodyMass":15,"Age":20,"Gender":"male","Week":[{"Day":[{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalorie":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":40,"Carbs":40,"Protein":40,"TotalCalories":34,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":100,"Carbs":100,"Protein":100,"TotalCalories":300,"DayCalorie":"normal","Weight":321,"Cardio":"missed","WeightTraining":"no"}]},{"Day":[{"Fat":11,"Carbs":11,"Protein":11,"TotalCalories":31,"DayCalorie":"normal","Weight":222,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"}]},{"Day":[{"Fat":110,"Carbs":110,"Protein":110,"TotalCalories":310,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"}]}]}]`
	respBody, _ := ioutil.ReadAll(req.Body)
	if strings.TrimSpace(string(respBody)) != strings.TrimSpace(EXPECTED) {
		t.Errorf("Auth returned unexpected body: got %v \nwant %v", string(respBody), EXPECTED)
		t.Fail()
	}
}

// func TestGetUIDFail(t *testing.T) {
// 	req, err := http.Get("http://localhost/getUID/testfail")
// 	if err != nil {
// 		t.Errorf("Connection failed: %v", err)
// 		t.Fail()
// 	}
// 	// Check the response body is what we expect.
// 	EXPECTED := `{"IsValid":false,"UID":""}`
// 	respBody, _ := ioutil.ReadAll(req.Body)
// 	if string(respBody) != EXPECTED {
// 		t.Errorf("Auth returned unexpected body: got %v \nwant %v", string(respBody), EXPECTED)
// 		t.Fail()
// 	}
// }
