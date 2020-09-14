package main

// this test requires mongoDB container to be running

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	structs "../common"
	"github.com/google/go-cmp/cmp"
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

//get testUID user info from database
func TestGetUserInfo(t *testing.T) {
	req, err := http.Get("http://localhost/userInfo/testUID")
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.Fail()
	}
	//unmarshal response into struct
	reqBody, _ := ioutil.ReadAll(req.Body)
	var reqAuth, expectedAuth structs.Client
	json.Unmarshal(reqBody, &reqAuth)
	//compare received struct with expected struct
	EXPECTED := []byte(`{"UID":"testUID","FirstName":"Anthony","LastName":"Hanna","Weight":215,"WaistCirc":35.5,"HeightInches":75,"LeanBodyMass":15,"Age":20,"Gender":"male","Week":[{"Day":[{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalorie":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":40,"Carbs":40,"Protein":40,"TotalCalories":34,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":100,"Carbs":100,"Protein":100,"TotalCalories":300,"DayCalorie":"normal","Weight":321,"Cardio":"missed","WeightTraining":"no"}]},{"Day":[{"Fat":11,"Carbs":11,"Protein":11,"TotalCalories":31,"DayCalorie":"normal","Weight":222,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"}]},{"Day":[{"Fat":110,"Carbs":110,"Protein":110,"TotalCalories":310,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"}]}],"Recommendation":[{"HighDayProtein":10,"HighDayCarb":11,"HighDayFat":12,"HighDayCalories":13,"NormalDayProtein":14,"NormalDayCarb":15,"NormalDayFat":16,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":24,"HIITChangeCarioIntervals":25,"Week":1,"ModifiedDate":"2020-09-13"}]}`)
	json.Unmarshal(EXPECTED, &expectedAuth)
	if !cmp.Equal(reqAuth, expectedAuth) {
		t.Errorf("Auth returned unexpected body: \ngot  %+v \nwant %+v", reqAuth, expectedAuth)
		t.Fail()
	}
}

//get "null" from a non-existant user in DB
func TestGetEmptyUserInfo(t *testing.T) {
	req, err := http.Get("http://localhost/userInfo/testUIDNoExists")
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

//create a new user and populate their data
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
	//unmarshal response into struct
	reqBody, _ := ioutil.ReadAll(req.Body)
	var reqAuth, expectedAuth structs.Client
	json.Unmarshal(reqBody, &reqAuth)
	//compare received struct with expected struct
	EXPECTED := []byte(`{"UID":"testUID2","FirstName":"Test","LastName":"User","Weight":111,"WaistCirc":111.1,"HeightInches":111,"LeanBodyMass":111,"Age":111,"Gender":"female","Week":[{"Day":[{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalorie":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":40,"Carbs":40,"Protein":40,"TotalCalories":34,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":100,"Carbs":100,"Protein":100,"TotalCalories":300,"DayCalorie":"normal","Weight":321,"Cardio":"missed","WeightTraining":"no"}]}]}`)
	json.Unmarshal(EXPECTED, &expectedAuth)
	if !cmp.Equal(reqAuth, expectedAuth) {
		t.Errorf("Auth returned unexpected body: \ngot  %+v \nwant %+v", reqAuth, expectedAuth)
		t.Fail()
	}
}

//update existing User info
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
	//unmarshal response into struct
	reqBody, _ := ioutil.ReadAll(req.Body)
	var reqAuth, expectedAuth structs.Client
	json.Unmarshal(reqBody, &reqAuth)
	//compare received struct with expected struct
	EXPECTED := []byte(`{"UID":"testUID2","FirstName":"TestUpdate","LastName":"UpdateUser","Weight":222,"WaistCirc":222.2,"HeightInches":222,"LeanBodyMass":222,"Age":222,"Gender":"male","Week":[{"Day":[{"Fat":222,"Carbs":222,"Protein":222,"TotalCalories":222,"DayCalorie":"low","Weight":222,"Cardio":"missed","WeightTraining":"no"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalorie":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":40,"Carbs":40,"Protein":40,"TotalCalories":34,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":100,"Carbs":100,"Protein":100,"TotalCalories":300,"DayCalorie":"normal","Weight":321,"Cardio":"missed","WeightTraining":"no"}]}]}`)
	json.Unmarshal(EXPECTED, &expectedAuth)
	if !cmp.Equal(reqAuth, expectedAuth) {
		t.Errorf("Auth returned unexpected body: \ngot  %+v \nwant %+v", reqAuth, expectedAuth)
		t.Fail()
	}
}

//update only user profile data and verify week data untouched
func TestUpdateUserProfile(t *testing.T) {
	//create control user in DB
	postBody := []byte(`{"UID":"testUID3","FirstName":"TestUpdate","LastName":"UpdateUser","Weight":222,"WaistCirc":222.2,"HeightInches":222,"LeanBodyMass":222,"Age":222,"Gender":"male","Week":[{"Day":[{"Fat":222,"Carbs":222,"Protein":222,"TotalCalories":222,"DayCalorie":"low","Weight":222,"Cardio":"missed","WeightTraining":"no"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalorie":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":40,"Carbs":40,"Protein":40,"TotalCalories":34,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":100,"Carbs":100,"Protein":100,"TotalCalories":300,"DayCalorie":"normal","Weight":321,"Cardio":"missed","WeightTraining":"no"}]}]}`)
	resp, err := http.Post("http://localhost/userInfo/testUID3", "application/json; charset=UTF-8", bytes.NewBuffer(postBody))
	if err != nil {
		t.Errorf("Post failed: %v", err)
		t.Fail()
	}
	respBody, _ := ioutil.ReadAll(resp.Body)
	if strings.TrimSpace(string(respBody)) != "ok" {
		t.Errorf("Auth returned unexpected body: got %v \nwant %v", string(respBody), "ok")
		t.Fail()
	}

	//do post request and verify success
	postBody = []byte(`{"UID":"testUID3","FirstName":"ProfileOnlyFirst","LastName":"ProfileOnlyLast","Weight":333,"WaistCirc":333,"HeightInches":333,"LeanBodyMass":333,"Age":333,"Gender":"female"}`)
	// post some data
	resp, err = http.Post("http://localhost/userInfo/testUID3", "application/json; charset=UTF-8", bytes.NewBuffer(postBody))
	if err != nil {
		t.Errorf("Post failed: %v", err)
		t.Fail()
	}
	respBody, _ = ioutil.ReadAll(resp.Body)
	if strings.TrimSpace(string(respBody)) != "ok" {
		t.Errorf("Auth returned unexpected body: got %v \nwant %v", string(respBody), "ok")
		t.Fail()
	}

	//validate post request
	req, err := http.Get("http://localhost/userInfo/testUID3")
	if err != nil {
		t.Errorf("Request failed: %v", err)
		t.Fail()
	}
	//unmarshal response into struct
	reqBody, _ := ioutil.ReadAll(req.Body)
	var reqAuth, expectedAuth structs.Client
	json.Unmarshal(reqBody, &reqAuth)
	//compare received struct with expected struct
	EXPECTED := []byte(`{"UID":"testUID3","FirstName":"ProfileOnlyFirst","LastName":"ProfileOnlyLast","Weight":333,"WaistCirc":333,"HeightInches":333,"LeanBodyMass":333,"Age":333,"Gender":"female","Week":[{"Day":[{"Fat":222,"Carbs":222,"Protein":222,"TotalCalories":222,"DayCalorie":"low","Weight":222,"Cardio":"missed","WeightTraining":"no"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalorie":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":40,"Carbs":40,"Protein":40,"TotalCalories":34,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":100,"Carbs":100,"Protein":100,"TotalCalories":300,"DayCalorie":"normal","Weight":321,"Cardio":"missed","WeightTraining":"no"}]}]}`)
	json.Unmarshal(EXPECTED, &expectedAuth)
	if !cmp.Equal(reqAuth, expectedAuth) {
		t.Errorf("Auth returned unexpected body: \ngot  %+v \nwant %+v", reqAuth, expectedAuth)
		t.Fail()
	}
}

// update user baseline data and verify profile data untouched
func TestUpdateUserBaseline(t *testing.T) {
	//create control user in DB
	postBody := []byte(`{"UID":"testUID4","FirstName":"Test4","LastName":"Update4","Weight":222,"WaistCirc":222.2,"HeightInches":222,"LeanBodyMass":222,"Age":222,"Gender":"male","Week":[{"Day":[{"Fat":222,"Carbs":222,"Protein":222,"TotalCalories":222,"DayCalorie":"low","Weight":222,"Cardio":"missed","WeightTraining":"no"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalorie":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":40,"Carbs":40,"Protein":40,"TotalCalories":34,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":100,"Carbs":100,"Protein":100,"TotalCalories":300,"DayCalorie":"normal","Weight":321,"Cardio":"missed","WeightTraining":"no"}]}]}`)
	resp, err := http.Post("http://localhost/userInfo/testUID4", "application/json; charset=UTF-8", bytes.NewBuffer(postBody))
	if err != nil {
		t.Errorf("Post failed: %v", err)
		t.Fail()
	}
	respBody, _ := ioutil.ReadAll(resp.Body)
	if strings.TrimSpace(string(respBody)) != "ok" {
		t.Errorf("Auth returned unexpected body: got %v \nwant %v", string(respBody), "ok")
		t.Fail()
	}

	//do post request and verify success
	postBody = []byte(`{"Day":[{"Fat":777,"Carbs":777,"Protein":777,"TotalCalories":999,"DayCalorie":"normal","Weight":999,"Cardio":"missed","WeightTraining":"yes"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalorie":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":40,"Carbs":40,"Protein":40,"TotalCalories":34,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":123,"Carbs":123,"Protein":123,"TotalCalories":123,"DayCalorie":"low","Weight":123,"Cardio":"missed","WeightTraining":"yes"}]}`)
	// post some data
	resp, err = http.Post("http://localhost/userBaseline/2/testUID4", "application/json; charset=UTF-8", bytes.NewBuffer(postBody))
	if err != nil {
		t.Errorf("Post failed: %v", err)
		t.Fail()
	}
	respBody, _ = ioutil.ReadAll(resp.Body)
	if strings.TrimSpace(string(respBody)) != "ok" {
		t.Errorf("Auth returned unexpected body: got %v \nwant %v", string(respBody), "ok")
		t.Fail()
	}

	//validate post request
	req, err := http.Get("http://localhost/userInfo/testUID4")
	if err != nil {
		t.Errorf("Request failed: %v", err)
		t.Fail()
	}
	//unmarshal response into struct
	reqBody, _ := ioutil.ReadAll(req.Body)
	var reqAuth, expectedAuth structs.Client
	json.Unmarshal(reqBody, &reqAuth)
	//compare received struct with expected struct
	EXPECTED := []byte(`{"UID":"testUID4","FirstName":"Test4","LastName":"Update4","Weight":222,"WaistCirc":222.2,"HeightInches":222,"LeanBodyMass":222,"Age":222,"Gender":"male","Week":[{"Day":[{"Fat":222,"Carbs":222,"Protein":222,"TotalCalories":222,"DayCalorie":"low","Weight":222,"Cardio":"missed","WeightTraining":"no"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalorie":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":40,"Carbs":40,"Protein":40,"TotalCalories":34,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":100,"Carbs":100,"Protein":100,"TotalCalories":300,"DayCalorie":"normal","Weight":321,"Cardio":"missed","WeightTraining":"no"}]},{"Day":null},{"Day":[{"Fat":777,"Carbs":777,"Protein":777,"TotalCalories":999,"DayCalorie":"normal","Weight":999,"Cardio":"missed","WeightTraining":"yes"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalorie":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":40,"Carbs":40,"Protein":40,"TotalCalories":34,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":123,"Carbs":123,"Protein":123,"TotalCalories":123,"DayCalorie":"low","Weight":123,"Cardio":"missed","WeightTraining":"yes"}]},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null}]}`)
	json.Unmarshal(EXPECTED, &expectedAuth)
	if !cmp.Equal(reqAuth, expectedAuth) {
		t.Errorf("Auth returned unexpected body: \ngot  %+v \nwant %+v", reqAuth, expectedAuth)
		t.Fail()
	}
}

// update user baseline data for user with broken week data
func TestUpdateInvalidUserBaseline(t *testing.T) {
	//create control user in DB
	postBody := []byte(`{"UID":"testUID5","FirstName":"Test5","LastName":"Update5","Weight":555,"WaistCirc":555.2,"HeightInches":555,"LeanBodyMass":555,"Age":555,"Gender":"female","Week":[{"Day":[{"Fat":555,"Carbs":555,"Protein":555,"TotalCalories":555,"DayCalorie":"high","Weight":555,"Cardio":"missed","WeightTraining":"yes"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalorie":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":40,"Carbs":40,"Protein":40,"TotalCalories":34,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":5555,"Carbs":5555,"Protein":5555,"TotalCalories":5555,"DayCalorie":"normal","Weight":5555,"Cardio":"missed","WeightTraining":"no"}]}]}`)
	resp, err := http.Post("http://localhost/userInfo/testUID5", "application/json; charset=UTF-8", bytes.NewBuffer(postBody))
	if err != nil {
		t.Errorf("Post failed: %v", err)
		t.Fail()
	}
	respBody, _ := ioutil.ReadAll(resp.Body)
	if strings.TrimSpace(string(respBody)) != "ok" {
		t.Errorf("Auth returned unexpected body: got %v \nwant %v", string(respBody), "ok")
		t.Fail()
	}

	//do post request and verify success
	postBody = []byte(`{"Day":[{"Fat":321,"Carbs":321,"Protein":321,"TotalCalories":321,"DayCalorie":"high","Weight":312,"Cardio":"missed","WeightTraining":"no"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalorie":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":40,"Carbs":40,"Protein":40,"TotalCalories":34,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":888,"Carbs":888,"Protein":888,"TotalCalories":888,"DayCalorie":"high","Weight":888,"Cardio":"missed","WeightTraining":"testSuccess"}]}`)
	// post some data
	resp, err = http.Post("http://localhost/userBaseline/23/testUID5", "application/json; charset=UTF-8", bytes.NewBuffer(postBody))
	if err != nil {
		t.Errorf("Post failed: %v", err)
		t.Fail()
	}
	respBody, _ = ioutil.ReadAll(resp.Body)
	if strings.TrimSpace(string(respBody)) != "ok" {
		t.Errorf("Auth returned unexpected body: got %v \nwant %v", string(respBody), "ok")
		t.Fail()
	}

	//validate post request
	req, err := http.Get("http://localhost/userInfo/testUID5")
	if err != nil {
		t.Errorf("Request failed: %v", err)
		t.Fail()
	}
	//unmarshal response into struct
	reqBody, _ := ioutil.ReadAll(req.Body)
	var reqAuth, expectedAuth structs.Client
	json.Unmarshal(reqBody, &reqAuth)
	//compare received struct with expected struct
	EXPECTED := []byte(`{"UID":"testUID5","FirstName":"Test5","LastName":"Update5","Weight":555,"WaistCirc":555.2,"HeightInches":555,"LeanBodyMass":555,"Age":555,"Gender":"female","Week":[{"Day":[{"Fat":555,"Carbs":555,"Protein":555,"TotalCalories":555,"DayCalorie":"high","Weight":555,"Cardio":"missed","WeightTraining":"yes"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalorie":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":40,"Carbs":40,"Protein":40,"TotalCalories":34,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":5555,"Carbs":5555,"Protein":5555,"TotalCalories":5555,"DayCalorie":"normal","Weight":5555,"Cardio":"missed","WeightTraining":"no"}]},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":null},{"Day":[{"Fat":321,"Carbs":321,"Protein":321,"TotalCalories":321,"DayCalorie":"high","Weight":312,"Cardio":"missed","WeightTraining":"no"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalorie":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":40,"Carbs":40,"Protein":40,"TotalCalories":34,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":888,"Carbs":888,"Protein":888,"TotalCalories":888,"DayCalorie":"high","Weight":888,"Cardio":"missed","WeightTraining":"testSuccess"}]}]}`)
	json.Unmarshal(EXPECTED, &expectedAuth)
	if !cmp.Equal(reqAuth, expectedAuth) {
		t.Errorf("Auth returned unexpected body: \ngot  %+v \nwant %+v", reqAuth, expectedAuth)
		t.Fail()
	}
}

// add user Recommendation object
func TestAddUserRecommendation(t *testing.T) {
	//create control user in DB
	postBody := []byte(`{"UID":"testUID6","FirstName":"Test6","LastName":"Update6","Weight":666,"WaistCirc":666.2,"HeightInches":666,"LeanBodyMass":666,"Age":666,"Gender":"female","Week":[{"Day":[{"Fat":666,"Carbs":666,"Protein":666,"TotalCalories":666,"DayCalorie":"high","Weight":666,"Cardio":"missed","WeightTraining":"yes"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalorie":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":40,"Carbs":40,"Protein":40,"TotalCalories":34,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":6665,"Carbs":6665,"Protein":6665,"TotalCalories":6665,"DayCalorie":"normal","Weight":6665,"Cardio":"missed","WeightTraining":"no"}]}]}`)
	resp, err := http.Post("http://localhost/userInfo/testUID6", "application/json; charset=UTF-8", bytes.NewBuffer(postBody))
	if err != nil {
		t.Errorf("Post failed: %v", err)
		t.Fail()
	}
	respBody, _ := ioutil.ReadAll(resp.Body)
	if strings.TrimSpace(string(respBody)) != "ok" {
		t.Errorf("Auth returned unexpected body: got %v \nwant %v", string(respBody), "ok")
		t.Fail()
	}

	//do post request and verify success
	postBody = []byte(`{"Recommendation":[{"HighDayProtein":10,"HighDayCarb":11,"HighDayFat":12,"HighDayCalories":13,"NormalDayProtein":14,"NormalDayCarb":15,"NormalDayFat":16,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":24,"HIITChangeCarioIntervals":25,"Week":1,"ModifiedDate":"2020-09-13"}]}`)
	// post some data
	resp, err = http.Post("http://localhost/userInfo/testUID6", "application/json; charset=UTF-8", bytes.NewBuffer(postBody))
	if err != nil {
		t.Errorf("Post failed: %v", err)
		t.Fail()
	}
	respBody, _ = ioutil.ReadAll(resp.Body)
	if strings.TrimSpace(string(respBody)) != "ok" {
		t.Errorf("Auth returned unexpected body: got %v \nwant %v", string(respBody), "ok")
		t.Fail()
	}

	//validate post request
	req, err := http.Get("http://localhost/userInfo/testUID6")
	if err != nil {
		t.Errorf("Request failed: %v", err)
		t.Fail()
	}
	//unmarshal response into struct
	reqBody, _ := ioutil.ReadAll(req.Body)
	var reqAuth, expectedAuth structs.Client
	json.Unmarshal(reqBody, &reqAuth)
	//compare received struct with expected struct
	EXPECTED := []byte(`{"UID":"testUID6","FirstName":"Test6","LastName":"Update6","Weight":666,"WaistCirc":666.2,"HeightInches":666,"LeanBodyMass":666,"Age":666,"Gender":"female","Week":[{"Day":[{"Fat":666,"Carbs":666,"Protein":666,"TotalCalories":666,"DayCalorie":"high","Weight":666,"Cardio":"missed","WeightTraining":"yes"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalorie":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":40,"Carbs":40,"Protein":40,"TotalCalories":34,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":6665,"Carbs":6665,"Protein":6665,"TotalCalories":6665,"DayCalorie":"normal","Weight":6665,"Cardio":"missed","WeightTraining":"no"}]}],"Recommendation":[{"HighDayProtein":10,"HighDayCarb":11,"HighDayFat":12,"HighDayCalories":13,"NormalDayProtein":14,"NormalDayCarb":15,"NormalDayFat":16,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":24,"HIITChangeCarioIntervals":25,"Week":1,"ModifiedDate":"2020-09-13"}]}`)
	json.Unmarshal(EXPECTED, &expectedAuth)
	if !cmp.Equal(reqAuth, expectedAuth) {
		t.Errorf("Auth returned unexpected body: \ngot  %+v \nwant %+v", reqAuth, expectedAuth)
		t.Fail()
	}
}

// update user Recommendation object
func TestUpdateUserRecommendation(t *testing.T) {
	//create control user in DB
	postBody := []byte(`{"UID":"testUID7","FirstName":"Test7","LastName":"Update7","Weight":777,"WaistCirc":777.2,"HeightInches":777,"LeanBodyMass":777,"Age":777,"Gender":"female","Recommendation":[{"HighDayProtein":10,"HighDayCarb":11,"HighDayFat":12,"HighDayCalories":13,"NormalDayProtein":14,"NormalDayCarb":15,"NormalDayFat":16,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":24,"HIITChangeCarioIntervals":25,"Week":1,"ModifiedDate":"2020-09-13"}],"Week":[{"Day":[{"Fat":777,"Carbs":777,"Protein":777,"TotalCalories":777,"DayCalorie":"high","Weight":777,"Cardio":"missed","WeightTraining":"yes"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalorie":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":40,"Carbs":40,"Protein":40,"TotalCalories":34,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":7775,"Carbs":7775,"Protein":7775,"TotalCalories":7775,"DayCalorie":"normal","Weight":7775,"Cardio":"missed","WeightTraining":"no"}]}]}`)
	resp, err := http.Post("http://localhost/userInfo/testUID7", "application/json; charset=UTF-8", bytes.NewBuffer(postBody))
	if err != nil {
		t.Errorf("Post failed: %v", err)
		t.Fail()
	}
	respBody, _ := ioutil.ReadAll(resp.Body)
	if strings.TrimSpace(string(respBody)) != "ok" {
		t.Errorf("Auth returned unexpected body: got %v \nwant %v", string(respBody), "ok")
		t.Fail()
	}

	//do post request and verify success
	postBody = []byte(`{"HighDayProtein":777,"HighDayCarb":777,"HighDayFat":777,"HighDayCalories":777,"NormalDayProtein":141,"NormalDayCarb":151,"NormalDayFat":161,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":241,"HIITChangeCarioIntervals":251,"Week":3,"ModifiedDate":"2020-09-77"}`)
	// post some data
	resp, err = http.Post("http://localhost/userRecommendation/2/testUID7", "application/json; charset=UTF-8", bytes.NewBuffer(postBody))
	if err != nil {
		t.Errorf("Post failed: %v", err)
		t.Fail()
	}
	respBody, _ = ioutil.ReadAll(resp.Body)
	if strings.TrimSpace(string(respBody)) != "ok" {
		t.Errorf("Auth returned unexpected body: got %v \nwant %v", string(respBody), "ok")
		t.Fail()
	}

	//validate post request
	req, err := http.Get("http://localhost/userInfo/testUID7")
	if err != nil {
		t.Errorf("Request failed: %v", err)
		t.Fail()
	}
	//unmarshal response into struct
	reqBody, _ := ioutil.ReadAll(req.Body)
	var reqAuth, expectedAuth structs.Client
	json.Unmarshal(reqBody, &reqAuth)
	//compare received struct with expected struct
	EXPECTED := []byte(`{"UID":"testUID7","FirstName":"Test7","LastName":"Update7","Weight":777,"WaistCirc":777.2,"HeightInches":777,"LeanBodyMass":777,"Age":777,"Gender":"female","Week":[{"Day":[{"Fat":777,"Carbs":777,"Protein":777,"TotalCalories":777,"DayCalorie":"high","Weight":777,"Cardio":"missed","WeightTraining":"yes"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalorie":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":40,"Carbs":40,"Protein":40,"TotalCalories":34,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":7775,"Carbs":7775,"Protein":7775,"TotalCalories":7775,"DayCalorie":"normal","Weight":7775,"Cardio":"missed","WeightTraining":"no"}]}],"Recommendation":[{"HighDayProtein":10,"HighDayCarb":11,"HighDayFat":12,"HighDayCalories":13,"NormalDayProtein":14,"NormalDayCarb":15,"NormalDayFat":16,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":24,"HIITChangeCarioIntervals":25,"Week":1,"ModifiedDate":"2020-09-13"},{},{"HighDayProtein":777,"HighDayCarb":777,"HighDayFat":777,"HighDayCalories":777,"NormalDayProtein":141,"NormalDayCarb":151,"NormalDayFat":161,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":241,"HIITChangeCarioIntervals":251,"Week":3,"ModifiedDate":"2020-09-77"},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}]}`)
	json.Unmarshal(EXPECTED, &expectedAuth)
	if !cmp.Equal(reqAuth, expectedAuth) {
		t.Errorf("Auth returned unexpected body: \ngot  %+v \nwant %+v", reqAuth, expectedAuth)
		t.Fail()
	}
}

// Another Update user Recommendation object
func TestUpdateUserRecommendation2(t *testing.T) {
	//create control user in DB
	postBody := []byte(`{"UID":"testUID8","FirstName":"Test8","LastName":"Update8","Weight":888,"WaistCirc":888.2,"HeightInches":888,"LeanBodyMass":888,"Age":888,"Gender":"female","Week":[{"Day":[{"Fat":888,"Carbs":888,"Protein":888,"TotalCalories":888,"DayCalorie":"high","Weight":888,"Cardio":"missed","WeightTraining":"yes"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"}]}],"Recommendation":[{"HighDayProtein":10,"HighDayCarb":11,"HighDayFat":12,"HighDayCalories":13,"NormalDayProtein":14,"NormalDayCarb":15,"NormalDayFat":16,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":24,"HIITChangeCarioIntervals":25,"Week":1,"ModifiedDate":"2020-09-13"},{},{"HighDayProtein":888,"HighDayCarb":888,"HighDayFat":888,"HighDayCalories":888,"NormalDayProtein":141,"NormalDayCarb":151,"NormalDayFat":161,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":241,"HIITChangeCarioIntervals":251,"Week":3,"ModifiedDate":"2020-09-00"},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}]}`)
	resp, err := http.Post("http://localhost/userInfo/testUID8", "application/json; charset=UTF-8", bytes.NewBuffer(postBody))
	if err != nil {
		t.Errorf("Post failed: %v", err)
		t.Fail()
	}
	respBody, _ := ioutil.ReadAll(resp.Body)
	if strings.TrimSpace(string(respBody)) != "ok" {
		t.Errorf("Auth returned unexpected body: got %v \nwant %v", string(respBody), "ok")
		t.Fail()
	}

	//do post request and verify success
	postBody = []byte(`{"HighDayProtein":888,"HighDayCarb":888,"HighDayFat":888,"HighDayCalories":888,"NormalDayProtein":141,"NormalDayCarb":151,"NormalDayFat":161,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":241,"HIITChangeCarioIntervals":251,"Week":3,"ModifiedDate":"2020-09-88"}`)
	// post some data
	resp, err = http.Post("http://localhost/userRecommendation/0/testUID8", "application/json; charset=UTF-8", bytes.NewBuffer(postBody))
	if err != nil {
		t.Errorf("Post failed: %v", err)
		t.Fail()
	}
	respBody, _ = ioutil.ReadAll(resp.Body)
	if strings.TrimSpace(string(respBody)) != "ok" {
		t.Errorf("Auth returned unexpected body: got %v \nwant %v", string(respBody), "ok")
		t.Fail()
	}

	//validate post request
	req, err := http.Get("http://localhost/userInfo/testUID8")
	if err != nil {
		t.Errorf("Request failed: %v", err)
		t.Fail()
	}
	//unmarshal response into struct
	reqBody, _ := ioutil.ReadAll(req.Body)
	var reqAuth, expectedAuth structs.Client
	json.Unmarshal(reqBody, &reqAuth)
	//compare received struct with expected struct
	EXPECTED := []byte(`{"UID":"testUID8","FirstName":"Test8","LastName":"Update8","Weight":888,"WaistCirc":888.2,"HeightInches":888,"LeanBodyMass":888,"Age":888,"Gender":"female","Week":[{"Day":[{"Fat":888,"Carbs":888,"Protein":888,"TotalCalories":888,"DayCalorie":"high","Weight":888,"Cardio":"missed","WeightTraining":"yes"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalorie":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"}]}],"Recommendation":[{"HighDayProtein":888,"HighDayCarb":888,"HighDayFat":888,"HighDayCalories":888,"NormalDayProtein":141,"NormalDayCarb":151,"NormalDayFat":161,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":241,"HIITChangeCarioIntervals":251,"Week":3,"ModifiedDate":"2020-09-88"},{},{"HighDayProtein":888,"HighDayCarb":888,"HighDayFat":888,"HighDayCalories":888,"NormalDayProtein":141,"NormalDayCarb":151,"NormalDayFat":161,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":241,"HIITChangeCarioIntervals":251,"Week":3,"ModifiedDate":"2020-09-00"},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}]}`)
	json.Unmarshal(EXPECTED, &expectedAuth)
	if !cmp.Equal(reqAuth, expectedAuth) {
		t.Errorf("Auth returned unexpected body: \ngot  %+v \nwant %+v", reqAuth, expectedAuth)
		t.Fail()
	}
}
