package main
// this test requires mock mongoDB container to be running

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"os"
	"log"

	structs "../common"
	"github.com/google/go-cmp/cmp"
)

var databaseAddress string
// get service address from env
func TestMain(m *testing.M) {
	databaseAddress = os.Getenv("DATABASE_SERVICE_ADDRESS")
	if databaseAddress == "" {
		databaseAddress = "localhost:8050"
	}
	log.Println("Testing Database at address: " + databaseAddress)
    os.Exit(m.Run())
}

func TestApiVersion(t *testing.T) {
	req, err := http.Get("http://"+databaseAddress+"/apiVersion")
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
	req, err := http.Get("http://"+databaseAddress+"/userInfo/testUID")
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.Fail()
	}
	//unmarshal response into struct
	reqBody, _ := ioutil.ReadAll(req.Body)
	var reqAuth, expectedAuth structs.Client
	json.Unmarshal(reqBody, &reqAuth)
	//compare received struct with expected struct
	EXPECTED := []byte(`{"UID":"testUID","FirstName":"Anthony","LastName":"Hanna","Weight":215,"WaistCirc":35.5,"HeightInches":75,"LeanBodyMass":15,"Age":20,"StartDate":"2020-08-15","Gender":"male","Week":[{"Day":[{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalories":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":40,"Carbs":40,"Protein":40,"TotalCalories":34,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":100,"Carbs":100,"Protein":100,"TotalCalories":300,"DayCalories":"normal","Weight":321,"Cardio":"missed","WeightTraining":"no"}]},{"Day":[{"Fat":11,"Carbs":11,"Protein":11,"TotalCalories":31,"DayCalories":"normal","Weight":222,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{},{},{},{},{}]},{"Day":[{"Fat":110,"Carbs":110,"Protein":110,"TotalCalories":310,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]}],"Recommendation":[{"HighDayProtein":10,"HighDayCarb":11,"HighDayFat":12,"HighDayCalories":13,"NormalDayProtein":14,"NormalDayCarb":15,"NormalDayFat":16,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":24,"HIITChangeCardioIntervals":25,"ModifiedDate":"2020-09-13"},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}]}`)
	json.Unmarshal(EXPECTED, &expectedAuth)
	if !cmp.Equal(reqAuth, expectedAuth) {
		t.Errorf("Database returned unexpected body: \ngot -: %+v \nwant -: %+v", string(reqBody), string(EXPECTED))
		t.Fail()
	}
}

//get "null" from a non-existant user in DB
func TestGetEmptyUserInfo(t *testing.T) {
	req, err := http.Get("http://"+databaseAddress+"/userInfo/testUIDNoExists")
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
	postBody := []byte(`{"UID":"db-testUID2","FirstName":"Test","LastName":"User","Weight":111,"WaistCirc":111.1,"HeightInches":111,"LeanBodyMass":111,"Age":111,"Gender":"female","Week":[{"Day":[{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalories":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":40,"Carbs":40,"Protein":40,"TotalCalories":34,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":100,"Carbs":100,"Protein":100,"TotalCalories":300,"DayCalories":"normal","Weight":321,"Cardio":"missed","WeightTraining":"no"}]}]}`)
	// post some data
	resp, err := http.Post("http://"+databaseAddress+"/userInfo/db-testUID2", "application/json; charset=UTF-8", bytes.NewBuffer(postBody))
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
	req, err := http.Get("http://"+databaseAddress+"/userInfo/db-testUID2")
	if err != nil {
		t.Errorf("Request failed: %v", err)
		t.Fail()
	}
	//unmarshal response into struct
	reqBody, _ := ioutil.ReadAll(req.Body)
	var reqAuth, expectedAuth structs.Client
	json.Unmarshal(reqBody, &reqAuth)
	//compare received struct with expected struct
	EXPECTED := []byte(`{"UID":"db-testUID2","FirstName":"Test","LastName":"User","Weight":111,"WaistCirc":111.1,"HeightInches":111,"LeanBodyMass":111,"Age":111,"Gender":"female","Week":[{"Day":[{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalories":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":40,"Carbs":40,"Protein":40,"TotalCalories":34,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":100,"Carbs":100,"Protein":100,"TotalCalories":300,"DayCalories":"normal","Weight":321,"Cardio":"missed","WeightTraining":"no"}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]}],"Recommendation":[{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}]}`)
	json.Unmarshal(EXPECTED, &expectedAuth)
	if !cmp.Equal(reqAuth, expectedAuth) {
		t.Errorf("Database returned unexpected body: \ngot -: %+v \nwant -: %+v", string(reqBody), string(EXPECTED))
		t.Fail()
	}
}

//update existing User info
func TestUpdateUserInfo(t *testing.T) {
	//do post request and verify success
	postBody := []byte(`{"UID":"db-testUID2","FirstName":"TestUpdate","LastName":"UpdateUser","Weight":222,"WaistCirc":222.2,"HeightInches":222,"LeanBodyMass":222,"Age":222,"Gender":"male","Week":[{"Day":[{"Fat":222,"Carbs":222,"Protein":222,"TotalCalories":222,"DayCalories":"low","Weight":222,"Cardio":"missed","WeightTraining":"no"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalories":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":40,"Carbs":40,"Protein":40,"TotalCalories":34,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":100,"Carbs":100,"Protein":100,"TotalCalories":300,"DayCalories":"normal","Weight":321,"Cardio":"missed","WeightTraining":"no"}]}]}`)
	// post some data
	resp, err := http.Post("http://"+databaseAddress+"/userInfo/db-testUID2", "application/json; charset=UTF-8", bytes.NewBuffer(postBody))
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
	req, err := http.Get("http://"+databaseAddress+"/userInfo/db-testUID2")
	if err != nil {
		t.Errorf("Request failed: %v", err)
		t.Fail()
	}
	//unmarshal response into struct
	reqBody, _ := ioutil.ReadAll(req.Body)
	var reqAuth, expectedAuth structs.Client
	json.Unmarshal(reqBody, &reqAuth)
	//compare received struct with expected struct
	EXPECTED := []byte(`{"UID":"db-testUID2","FirstName":"TestUpdate","LastName":"UpdateUser","Weight":222,"WaistCirc":222.2,"HeightInches":222,"LeanBodyMass":222,"Age":222,"Gender":"male","Week":[{"Day":[{"Fat":222,"Carbs":222,"Protein":222,"TotalCalories":222,"DayCalories":"low","Weight":222,"Cardio":"missed","WeightTraining":"no"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalories":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":40,"Carbs":40,"Protein":40,"TotalCalories":34,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":100,"Carbs":100,"Protein":100,"TotalCalories":300,"DayCalories":"normal","Weight":321,"Cardio":"missed","WeightTraining":"no"}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]}],"Recommendation":[{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}]}`)
	json.Unmarshal(EXPECTED, &expectedAuth)
	if !cmp.Equal(reqAuth, expectedAuth) {
		t.Errorf("Database returned unexpected body: \ngot -: %+v \nwant -: %+v", string(reqBody), string(EXPECTED))
		t.Fail()
	}
}

//update only user profile data and verify week data untouched
func TestUpdateUserProfile(t *testing.T) {
	//create control user in DB
	postBody := []byte(`{"UID":"db-testUID3","FirstName":"TestUpdate","LastName":"UpdateUser","Weight":222,"WaistCirc":222.2,"HeightInches":222,"LeanBodyMass":222,"Age":222,"Gender":"male","Week":[{"Day":[{"Fat":222,"Carbs":222,"Protein":222,"TotalCalories":222,"DayCalories":"low","Weight":222,"Cardio":"missed","WeightTraining":"no"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalories":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":40,"Carbs":40,"Protein":40,"TotalCalories":34,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":100,"Carbs":100,"Protein":100,"TotalCalories":300,"DayCalories":"normal","Weight":321,"Cardio":"missed","WeightTraining":"no"}]}]}`)
	resp, err := http.Post("http://"+databaseAddress+"/userInfo/db-testUID3", "application/json; charset=UTF-8", bytes.NewBuffer(postBody))
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
	postBody = []byte(`{"UID":"db-testUID3","FirstName":"ProfileOnlyFirst","LastName":"ProfileOnlyLast","Weight":333,"WaistCirc":333,"HeightInches":333,"LeanBodyMass":333,"Age":333,"Gender":"female"}`)
	// post some data
	resp, err = http.Post("http://"+databaseAddress+"/userInfo/db-testUID3", "application/json; charset=UTF-8", bytes.NewBuffer(postBody))
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
	req, err := http.Get("http://"+databaseAddress+"/userInfo/db-testUID3")
	if err != nil {
		t.Errorf("Request failed: %v", err)
		t.Fail()
	}
	//unmarshal response into struct
	reqBody, _ := ioutil.ReadAll(req.Body)
	var reqAuth, expectedAuth structs.Client
	json.Unmarshal(reqBody, &reqAuth)
	//compare received struct with expected struct
	EXPECTED := []byte(`{"UID":"db-testUID3","FirstName":"ProfileOnlyFirst","LastName":"ProfileOnlyLast","Weight":333,"WaistCirc":333,"HeightInches":333,"LeanBodyMass":333,"Age":333,"Gender":"female","Week":[{"Day":[{"Fat":222,"Carbs":222,"Protein":222,"TotalCalories":222,"DayCalories":"low","Weight":222,"Cardio":"missed","WeightTraining":"no"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalories":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":40,"Carbs":40,"Protein":40,"TotalCalories":34,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":100,"Carbs":100,"Protein":100,"TotalCalories":300,"DayCalories":"normal","Weight":321,"Cardio":"missed","WeightTraining":"no"}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]}],"Recommendation":[{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}]}
	`)
	json.Unmarshal(EXPECTED, &expectedAuth)
	if !cmp.Equal(reqAuth, expectedAuth) {
		t.Errorf("Database returned unexpected body: \ngot -: %+v \nwant -: %+v", string(reqBody), string(EXPECTED))
		t.Fail()
	}
}

// update user weekly data and verify profile data untouched
func TestUpdateUserWeekly(t *testing.T) {
	//create control user in DB
	postBody := []byte(`{"UID":"db-testUID4","FirstName":"Test4","LastName":"Update4","Weight":222,"WaistCirc":222.2,"HeightInches":222,"LeanBodyMass":222,"Age":222,"Gender":"male","Week":[{"Day":[{"Fat":222,"Carbs":222,"Protein":222,"TotalCalories":222,"DayCalories":"low","Weight":222,"Cardio":"missed","WeightTraining":"no"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalories":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":40,"Carbs":40,"Protein":40,"TotalCalories":34,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":100,"Carbs":100,"Protein":100,"TotalCalories":300,"DayCalories":"normal","Weight":321,"Cardio":"missed","WeightTraining":"no"}]}]}`)
	resp, err := http.Post("http://"+databaseAddress+"/userInfo/db-testUID4", "application/json; charset=UTF-8", bytes.NewBuffer(postBody))
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
	postBody = []byte(`{"Day":[{"Fat":777,"Carbs":777,"Protein":777,"TotalCalories":999,"DayCalories":"normal","Weight":999,"Cardio":"missed","WeightTraining":"yes"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalories":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":40,"Carbs":40,"Protein":40,"TotalCalories":34,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":123,"Carbs":123,"Protein":123,"TotalCalories":123,"DayCalories":"low","Weight":123,"Cardio":"missed","WeightTraining":"yes"}]}`)
	// post some data
	resp, err = http.Post("http://"+databaseAddress+"/userWeekly/2/db-testUID4", "application/json; charset=UTF-8", bytes.NewBuffer(postBody))
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
	req, err := http.Get("http://"+databaseAddress+"/userInfo/db-testUID4")
	if err != nil {
		t.Errorf("Request failed: %v", err)
		t.Fail()
	}
	//unmarshal response into struct
	reqBody, _ := ioutil.ReadAll(req.Body)
	var reqAuth, expectedAuth structs.Client
	json.Unmarshal(reqBody, &reqAuth)
	//compare received struct with expected struct
	EXPECTED := []byte(`{"UID":"db-testUID4","FirstName":"Test4","LastName":"Update4","Weight":222,"WaistCirc":222.2,"HeightInches":222,"LeanBodyMass":222,"Age":222,"Gender":"male","Week":[{"Day":[{"Fat":222,"Carbs":222,"Protein":222,"TotalCalories":222,"DayCalories":"low","Weight":222,"Cardio":"missed","WeightTraining":"no"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalories":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":40,"Carbs":40,"Protein":40,"TotalCalories":34,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":100,"Carbs":100,"Protein":100,"TotalCalories":300,"DayCalories":"normal","Weight":321,"Cardio":"missed","WeightTraining":"no"}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{"Fat":777,"Carbs":777,"Protein":777,"TotalCalories":999,"DayCalories":"normal","Weight":999,"Cardio":"missed","WeightTraining":"yes"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalories":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":40,"Carbs":40,"Protein":40,"TotalCalories":34,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":123,"Carbs":123,"Protein":123,"TotalCalories":123,"DayCalories":"low","Weight":123,"Cardio":"missed","WeightTraining":"yes"}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]}],"Recommendation":[{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}]}`)
	json.Unmarshal(EXPECTED, &expectedAuth)
	if !cmp.Equal(reqAuth, expectedAuth) {
		t.Errorf("Database returned unexpected body: \ngot -: %+v \nwant -: %+v", string(reqBody), string(EXPECTED))
		t.Fail()
	}
}

// update user weekly data for user with broken week data
func TestUpdateInvalidUserWeekly(t *testing.T) {
	//create control user in DB
	postBody := []byte(`{"UID":"db-testUID5","FirstName":"Test5","LastName":"Update5","Weight":555,"WaistCirc":555.2,"HeightInches":555,"LeanBodyMass":555,"Age":555,"Gender":"female","Week":[{"Day":[{"Fat":555,"Carbs":555,"Protein":555,"TotalCalories":555,"DayCalories":"high","Weight":555,"Cardio":"missed","WeightTraining":"yes"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalories":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":40,"Carbs":40,"Protein":40,"TotalCalories":34,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":5555,"Carbs":5555,"Protein":5555,"TotalCalories":5555,"DayCalories":"normal","Weight":5555,"Cardio":"missed","WeightTraining":"no"}]}]}`)
	resp, err := http.Post("http://"+databaseAddress+"/userInfo/db-testUID5", "application/json; charset=UTF-8", bytes.NewBuffer(postBody))
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
	postBody = []byte(`{"Day":[{"Fat":321,"Carbs":321,"Protein":321,"TotalCalories":321,"DayCalories":"high","Weight":312,"Cardio":"missed","WeightTraining":"no"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalories":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":40,"Carbs":40,"Protein":40,"TotalCalories":34,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":888,"Carbs":888,"Protein":888,"TotalCalories":888,"DayCalories":"high","Weight":888,"Cardio":"missed","WeightTraining":"testSuccess"}]}`)
	// post some data
	resp, err = http.Post("http://"+databaseAddress+"/userWeekly/23/db-testUID5", "application/json; charset=UTF-8", bytes.NewBuffer(postBody))
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
	req, err := http.Get("http://"+databaseAddress+"/userInfo/db-testUID5")
	if err != nil {
		t.Errorf("Request failed: %v", err)
		t.Fail()
	}
	//unmarshal response into struct
	reqBody, _ := ioutil.ReadAll(req.Body)
	var reqAuth, expectedAuth structs.Client
	json.Unmarshal(reqBody, &reqAuth)
	//compare received struct with expected struct
	EXPECTED := []byte(`{"UID":"db-testUID5","FirstName":"Test5","LastName":"Update5","Weight":555,"WaistCirc":555.2,"HeightInches":555,"LeanBodyMass":555,"Age":555,"Gender":"female","Week":[{"Day":[{"Fat":555,"Carbs":555,"Protein":555,"TotalCalories":555,"DayCalories":"high","Weight":555,"Cardio":"missed","WeightTraining":"yes"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalories":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":40,"Carbs":40,"Protein":40,"TotalCalories":34,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":5555,"Carbs":5555,"Protein":5555,"TotalCalories":5555,"DayCalories":"normal","Weight":5555,"Cardio":"missed","WeightTraining":"no"}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{"Fat":321,"Carbs":321,"Protein":321,"TotalCalories":321,"DayCalories":"high","Weight":312,"Cardio":"missed","WeightTraining":"no"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalories":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":40,"Carbs":40,"Protein":40,"TotalCalories":34,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":888,"Carbs":888,"Protein":888,"TotalCalories":888,"DayCalories":"high","Weight":888,"Cardio":"missed","WeightTraining":"testSuccess"}]}],"Recommendation":[{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}]}`)
	json.Unmarshal(EXPECTED, &expectedAuth)
	if !cmp.Equal(reqAuth, expectedAuth) {
		t.Errorf("Database returned unexpected body: \ngot -: %+v \nwant -: %+v", string(reqBody), string(EXPECTED))
		t.Fail()
	}
}

// add user Recommendation object
func TestAddUserRecommendation(t *testing.T) {
	//create control user in DB
	postBody := []byte(`{"UID":"db-testUID6","FirstName":"Test6","LastName":"Update6","Weight":666,"WaistCirc":666.2,"HeightInches":666,"LeanBodyMass":666,"Age":666,"Gender":"female","Week":[{"Day":[{"Fat":666,"Carbs":666,"Protein":666,"TotalCalories":666,"DayCalories":"high","Weight":666,"Cardio":"missed","WeightTraining":"yes"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalories":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":40,"Carbs":40,"Protein":40,"TotalCalories":34,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":6665,"Carbs":6665,"Protein":6665,"TotalCalories":6665,"DayCalories":"normal","Weight":6665,"Cardio":"missed","WeightTraining":"no"}]}]}`)
	resp, err := http.Post("http://"+databaseAddress+"/userInfo/db-testUID6", "application/json; charset=UTF-8", bytes.NewBuffer(postBody))
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
	postBody = []byte(`{"Recommendation":[{"HighDayProtein":10,"HighDayCarb":11,"HighDayFat":12,"HighDayCalories":13,"NormalDayProtein":14,"NormalDayCarb":15,"NormalDayFat":16,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":24,"HIITChangeCardioIntervals":25,"ModifiedDate":"2020-09-13"}]}`)
	// post some data
	resp, err = http.Post("http://"+databaseAddress+"/userInfo/db-testUID6", "application/json; charset=UTF-8", bytes.NewBuffer(postBody))
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
	req, err := http.Get("http://"+databaseAddress+"/userInfo/db-testUID6")
	if err != nil {
		t.Errorf("Request failed: %v", err)
		t.Fail()
	}
	//unmarshal response into struct
	reqBody, _ := ioutil.ReadAll(req.Body)
	var reqAuth, expectedAuth structs.Client
	json.Unmarshal(reqBody, &reqAuth)
	//compare received struct with expected struct
	EXPECTED := []byte(`{"UID":"db-testUID6","FirstName":"Test6","LastName":"Update6","Weight":666,"WaistCirc":666.2,"HeightInches":666,"LeanBodyMass":666,"Age":666,"Gender":"female","Week":[{"Day":[{"Fat":666,"Carbs":666,"Protein":666,"TotalCalories":666,"DayCalories":"high","Weight":666,"Cardio":"missed","WeightTraining":"yes"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalories":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":40,"Carbs":40,"Protein":40,"TotalCalories":34,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":6665,"Carbs":6665,"Protein":6665,"TotalCalories":6665,"DayCalories":"normal","Weight":6665,"Cardio":"missed","WeightTraining":"no"}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]}],"Recommendation":[{"HighDayProtein":10,"HighDayCarb":11,"HighDayFat":12,"HighDayCalories":13,"NormalDayProtein":14,"NormalDayCarb":15,"NormalDayFat":16,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":24,"HIITChangeCardioIntervals":25,"ModifiedDate":"2020-09-13"},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}]}`)
	json.Unmarshal(EXPECTED, &expectedAuth)
	if !cmp.Equal(reqAuth, expectedAuth) {
		t.Errorf("Database returned unexpected body: \ngot -: %+v \nwant -: %+v", string(reqBody), string(EXPECTED))
		t.Fail()
	}
}

// update user Recommendation object
func TestUpdateUserRecommendation(t *testing.T) {
	//create control user in DB
	postBody := []byte(`{"UID":"db-testUID7","FirstName":"Test7","LastName":"Update7","Weight":777,"WaistCirc":777.2,"HeightInches":777,"LeanBodyMass":777,"Age":777,"Gender":"female","Recommendation":[{"HighDayProtein":10,"HighDayCarb":11,"HighDayFat":12,"HighDayCalories":13,"NormalDayProtein":14,"NormalDayCarb":15,"NormalDayFat":16,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":24,"HIITChangeCardioIntervals":25,"ModifiedDate":"2020-09-13"}],"Week":[{"Day":[{"Fat":777,"Carbs":777,"Protein":777,"TotalCalories":777,"DayCalories":"high","Weight":777,"Cardio":"missed","WeightTraining":"yes"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalories":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":40,"Carbs":40,"Protein":40,"TotalCalories":34,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":7775,"Carbs":7775,"Protein":7775,"TotalCalories":7775,"DayCalories":"normal","Weight":7775,"Cardio":"missed","WeightTraining":"no"}]}]}`)
	resp, err := http.Post("http://"+databaseAddress+"/userInfo/db-testUID7", "application/json; charset=UTF-8", bytes.NewBuffer(postBody))
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
	postBody = []byte(`{"HighDayProtein":777,"HighDayCarb":777,"HighDayFat":777,"HighDayCalories":777,"NormalDayProtein":141,"NormalDayCarb":151,"NormalDayFat":161,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":241,"HIITChangeCardioIntervals":251,"ModifiedDate":"2020-09-77"}`)
	// post some data
	resp, err = http.Post("http://"+databaseAddress+"/userRecommendation/2/db-testUID7", "application/json; charset=UTF-8", bytes.NewBuffer(postBody))
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
	req, err := http.Get("http://"+databaseAddress+"/userInfo/db-testUID7")
	if err != nil {
		t.Errorf("Request failed: %v", err)
		t.Fail()
	}
	//unmarshal response into struct
	reqBody, _ := ioutil.ReadAll(req.Body)
	var reqAuth, expectedAuth structs.Client
	json.Unmarshal(reqBody, &reqAuth)
	//compare received struct with expected struct
	EXPECTED := []byte(`{"UID":"db-testUID7","FirstName":"Test7","LastName":"Update7","Weight":777,"WaistCirc":777.2,"HeightInches":777,"LeanBodyMass":777,"Age":777,"Gender":"female","Week":[{"Day":[{"Fat":777,"Carbs":777,"Protein":777,"TotalCalories":777,"DayCalories":"high","Weight":777,"Cardio":"missed","WeightTraining":"yes"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalories":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":40,"Carbs":40,"Protein":40,"TotalCalories":34,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":7775,"Carbs":7775,"Protein":7775,"TotalCalories":7775,"DayCalories":"normal","Weight":7775,"Cardio":"missed","WeightTraining":"no"}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]}],"Recommendation":[{"HighDayProtein":10,"HighDayCarb":11,"HighDayFat":12,"HighDayCalories":13,"NormalDayProtein":14,"NormalDayCarb":15,"NormalDayFat":16,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":24,"HIITChangeCardioIntervals":25,"ModifiedDate":"2020-09-13"},{},{"HighDayProtein":777,"HighDayCarb":777,"HighDayFat":777,"HighDayCalories":777,"NormalDayProtein":141,"NormalDayCarb":151,"NormalDayFat":161,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":241,"HIITChangeCardioIntervals":251,"ModifiedDate":"2020-09-77"},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}]}`)
	json.Unmarshal(EXPECTED, &expectedAuth)
	if !cmp.Equal(reqAuth, expectedAuth) {
		t.Errorf("Database returned unexpected body: \ngot -: %+v \nwant -: %+v", string(reqBody), string(EXPECTED))
		t.Fail()
	}
}

// Another Update user Recommendation object
func TestUpdateUserRecommendation2(t *testing.T) {
	//create control user in DB
	postBody := []byte(`{"UID":"db-testUID8","FirstName":"Test8","LastName":"Update8","Weight":888,"WaistCirc":888.2,"HeightInches":888,"LeanBodyMass":888,"Age":888,"Gender":"female","Week":[{"Day":[{"Fat":888,"Carbs":888,"Protein":888,"TotalCalories":888,"DayCalories":"high","Weight":888,"Cardio":"missed","WeightTraining":"yes"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"}]}],"Recommendation":[{"HighDayProtein":10,"HighDayCarb":11,"HighDayFat":12,"HighDayCalories":13,"NormalDayProtein":14,"NormalDayCarb":15,"NormalDayFat":16,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":24,"HIITChangeCardioIntervals":25,"ModifiedDate":"2020-09-13"},{},{"HighDayProtein":888,"HighDayCarb":888,"HighDayFat":888,"HighDayCalories":888,"NormalDayProtein":141,"NormalDayCarb":151,"NormalDayFat":161,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":241,"HIITChangeCardioIntervals":251,"ModifiedDate":"2020-09-00"},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}]}`)
	resp, err := http.Post("http://"+databaseAddress+"/userInfo/db-testUID8", "application/json; charset=UTF-8", bytes.NewBuffer(postBody))
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
	postBody = []byte(`{"HighDayProtein":888,"HighDayCarb":888,"HighDayFat":888,"HighDayCalories":888,"NormalDayProtein":141,"NormalDayCarb":151,"NormalDayFat":161,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":241,"HIITChangeCardioIntervals":251,"ModifiedDate":"2020-09-88"}`)
	// post some data
	resp, err = http.Post("http://"+databaseAddress+"/userRecommendation/0/db-testUID8", "application/json; charset=UTF-8", bytes.NewBuffer(postBody))
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
	req, err := http.Get("http://"+databaseAddress+"/userInfo/db-testUID8")
	if err != nil {
		t.Errorf("Request failed: %v", err)
		t.Fail()
	}
	//unmarshal response into struct
	reqBody, _ := ioutil.ReadAll(req.Body)
	var reqAuth, expectedAuth structs.Client
	json.Unmarshal(reqBody, &reqAuth)
	//compare received struct with expected struct
	EXPECTED := []byte(`{"UID":"db-testUID8","FirstName":"Test8","LastName":"Update8","Weight":888,"WaistCirc":888.2,"HeightInches":888,"LeanBodyMass":888,"Age":888,"Gender":"female","Week":[{"Day":[{"Fat":888,"Carbs":888,"Protein":888,"TotalCalories":888,"DayCalories":"high","Weight":888,"Cardio":"missed","WeightTraining":"yes"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]}],"Recommendation":[{"HighDayProtein":888,"HighDayCarb":888,"HighDayFat":888,"HighDayCalories":888,"NormalDayProtein":141,"NormalDayCarb":151,"NormalDayFat":161,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":241,"HIITChangeCardioIntervals":251,"ModifiedDate":"2020-09-88"},{},{"HighDayProtein":888,"HighDayCarb":888,"HighDayFat":888,"HighDayCalories":888,"NormalDayProtein":141,"NormalDayCarb":151,"NormalDayFat":161,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":241,"HIITChangeCardioIntervals":251,"ModifiedDate":"2020-09-00"},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}]}
`)
	json.Unmarshal(EXPECTED, &expectedAuth)
	if !cmp.Equal(reqAuth, expectedAuth) {
		t.Errorf("Database returned unexpected body: \ngot -: %+v \nwant -: %+v", string(reqBody), string(EXPECTED))
		t.Fail()
	}
}

// test inserting new day on existing Day
func TestUpdateDay(t *testing.T) {
	//create control user in DB
	postBody := []byte(`{"UID":"db-testUID9","FirstName":"Test9","LastName":"Update9","Weight":999,"WaistCirc":999.2,"HeightInches":999,"LeanBodyMass":999,"Age":999,"Gender":"female","Week":[{"Day":[{"Fat":999,"Carbs":999,"Protein":999,"TotalCalories":999,"DayCalories":"high","Weight":999,"Cardio":"missed","WeightTraining":"yes"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"}]}],"Recommendation":[{"HighDayProtein":10,"HighDayCarb":11,"HighDayFat":12,"HighDayCalories":13,"NormalDayProtein":14,"NormalDayCarb":15,"NormalDayFat":16,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":24,"HIITChangeCardioIntervals":25,"ModifiedDate":"2020-09-13"},{},{"HighDayProtein":999,"HighDayCarb":999,"HighDayFat":999,"HighDayCalories":999,"NormalDayProtein":141,"NormalDayCarb":151,"NormalDayFat":161,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":241,"HIITChangeCardioIntervals":251,"ModifiedDate":"2020-09-99"},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}]}`)
	resp, err := http.Post("http://"+databaseAddress+"/userInfo/db-testUID9", "application/json; charset=UTF-8", bytes.NewBuffer(postBody))
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
	postBody = []byte(`{"Fat":1,"Carbs":2,"Protein":3,"TotalCalories":4,"DayCalories":"low","Weight":5,"Cardio":"no","WeightTraining":"TestWorked"}`)
	// post some data
	resp, err = http.Post("http://"+databaseAddress+"/userDaily/0/0/db-testUID9", "application/json; charset=UTF-8", bytes.NewBuffer(postBody))
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
	req, err := http.Get("http://"+databaseAddress+"/userInfo/db-testUID9")
	if err != nil {
		t.Errorf("Request failed: %v", err)
		t.Fail()
	}
	//unmarshal response into struct
	reqBody, _ := ioutil.ReadAll(req.Body)
	var reqAuth, expectedAuth structs.Client
	json.Unmarshal(reqBody, &reqAuth)
	//compare received struct with expected struct
	EXPECTED := []byte(`{"UID":"db-testUID9","FirstName":"Test9","LastName":"Update9","Weight":999,"WaistCirc":999.2,"HeightInches":999,"LeanBodyMass":999,"Age":999,"Gender":"female","Week":[{"Day":[{"Fat":1,"Carbs":2,"Protein":3,"TotalCalories":4,"DayCalories":"low","Weight":5,"Cardio":"no","WeightTraining":"TestWorked"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]}],"Recommendation":[{"HighDayProtein":10,"HighDayCarb":11,"HighDayFat":12,"HighDayCalories":13,"NormalDayProtein":14,"NormalDayCarb":15,"NormalDayFat":16,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":24,"HIITChangeCardioIntervals":25,"ModifiedDate":"2020-09-13"},{},{"HighDayProtein":999,"HighDayCarb":999,"HighDayFat":999,"HighDayCalories":999,"NormalDayProtein":141,"NormalDayCarb":151,"NormalDayFat":161,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":241,"HIITChangeCardioIntervals":251,"ModifiedDate":"2020-09-99"},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}]}`)
	json.Unmarshal(EXPECTED, &expectedAuth)
	if !cmp.Equal(reqAuth, expectedAuth) {
		t.Errorf("Database returned unexpected body: \ngot -: %+v \nwant -: %+v", string(reqBody), string(EXPECTED))
		t.Fail()
	}
}

// test inserting new day on existing Week
func TestUpdateDayWeek(t *testing.T) {
	//create control user in DB
	postBody := []byte(`{"UID":"db-testUID9w","FirstName":"Test9w","LastName":"Update9w","Weight":99910,"WaistCirc":99910.2,"HeightInches":99910,"LeanBodyMass":99910,"Age":99910,"Gender":"male","Week":[{"Day":[{"Fat":999,"Carbs":999,"Protein":999,"TotalCalories":999,"DayCalories":"high","Weight":999,"Cardio":"missed","WeightTraining":"yes"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"}]}],"Recommendation":[{"HighDayProtein":10,"HighDayCarb":11,"HighDayFat":12,"HighDayCalories":13,"NormalDayProtein":14,"NormalDayCarb":15,"NormalDayFat":16,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":24,"HIITChangeCardioIntervals":25,"ModifiedDate":"2020-09-13"},{},{"HighDayProtein":999,"HighDayCarb":999,"HighDayFat":999,"HighDayCalories":999,"NormalDayProtein":141,"NormalDayCarb":151,"NormalDayFat":161,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":241,"HIITChangeCardioIntervals":251,"ModifiedDate":"2020-09-ww"},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}]}`)
	resp, err := http.Post("http://"+databaseAddress+"/userInfo/db-testUID9w", "application/json; charset=UTF-8", bytes.NewBuffer(postBody))
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
	postBody = []byte(`{"Fat":1,"Carbs":2,"Protein":3,"TotalCalories":4,"DayCalories":"low","Weight":5,"Cardio":"no","WeightTraining":"TestWowwwwwrked"}`)
	// post some data
	resp, err = http.Post("http://"+databaseAddress+"/userDaily/0/5/db-testUID9w", "application/json; charset=UTF-8", bytes.NewBuffer(postBody))
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
	req, err := http.Get("http://"+databaseAddress+"/userInfo/db-testUID9w")
	if err != nil {
		t.Errorf("Request failed: %v", err)
		t.Fail()
	}
	//unmarshal response into struct
	reqBody, _ := ioutil.ReadAll(req.Body)
	var reqAuth, expectedAuth structs.Client
	json.Unmarshal(reqBody, &reqAuth)
	//compare received struct with expected struct
	EXPECTED := []byte(`{"UID":"db-testUID9w","FirstName":"Test9w","LastName":"Update9w","Weight":99910,"WaistCirc":99910.2,"HeightInches":99910,"LeanBodyMass":99910,"Age":99910,"Gender":"male","Week":[{"Day":[{"Fat":999,"Carbs":999,"Protein":999,"TotalCalories":999,"DayCalories":"high","Weight":999,"Cardio":"missed","WeightTraining":"yes"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{},{},{},{"Fat":1,"Carbs":2,"Protein":3,"TotalCalories":4,"DayCalories":"low","Weight":5,"Cardio":"no","WeightTraining":"TestWowwwwwrked"},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]}],"Recommendation":[{"HighDayProtein":10,"HighDayCarb":11,"HighDayFat":12,"HighDayCalories":13,"NormalDayProtein":14,"NormalDayCarb":15,"NormalDayFat":16,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":24,"HIITChangeCardioIntervals":25,"ModifiedDate":"2020-09-13"},{},{"HighDayProtein":999,"HighDayCarb":999,"HighDayFat":999,"HighDayCalories":999,"NormalDayProtein":141,"NormalDayCarb":151,"NormalDayFat":161,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":241,"HIITChangeCardioIntervals":251,"ModifiedDate":"2020-09-ww"},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}]}`)
	json.Unmarshal(EXPECTED, &expectedAuth)
	if !cmp.Equal(reqAuth, expectedAuth) {
		t.Errorf("Database returned unexpected body: \ngot -: %+v \nwant -: %+v", string(reqBody), string(EXPECTED))
		t.Fail()
	}
}

// test inserting new day on empty week
func TestUpdateNewWeekDay(t *testing.T) {
	//create control user in DB
	postBody := []byte(`{"UID":"db-testUID10","FirstName":"Test10","LastName":"Update10","Weight":101,"WaistCirc":101.2,"HeightInches":101,"LeanBodyMass":101,"Age":101,"Gender":"female","Week":[{"Day":[{"Fat":101,"Carbs":101,"Protein":101,"TotalCalories":101,"DayCalories":"high","Weight":101,"Cardio":"missed","WeightTraining":"yes"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"}]}],"Recommendation":[{"HighDayProtein":10,"HighDayCarb":11,"HighDayFat":12,"HighDayCalories":13,"NormalDayProtein":14,"NormalDayCarb":15,"NormalDayFat":16,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":24,"HIITChangeCardioIntervals":25,"ModifiedDate":"2020-09-13"},{},{"HighDayProtein":101,"HighDayCarb":101,"HighDayFat":101,"HighDayCalories":101,"NormalDayProtein":141,"NormalDayCarb":151,"NormalDayFat":161,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":241,"HIITChangeCardioIntervals":251,"ModifiedDate":"2020-09-99"},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}]}`)
	resp, err := http.Post("http://"+databaseAddress+"/userInfo/db-testUID10", "application/json; charset=UTF-8", bytes.NewBuffer(postBody))
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
	postBody = []byte(`{"Fat":1,"Carbs":2,"Protein":3,"TotalCalories":4,"DayCalories":"low","Weight":5,"Cardio":"no","WeightTraining":"weightTrained"}`)
	// post some data
	resp, err = http.Post("http://"+databaseAddress+"/userDaily/23/1/db-testUID10", "application/json; charset=UTF-8", bytes.NewBuffer(postBody))
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
	req, err := http.Get("http://"+databaseAddress+"/userInfo/db-testUID10")
	if err != nil {
		t.Errorf("Request failed: %v", err)
		t.Fail()
	}
	//unmarshal response into struct
	reqBody, _ := ioutil.ReadAll(req.Body)
	var reqAuth, expectedAuth structs.Client
	json.Unmarshal(reqBody, &reqAuth)
	//compare received struct with expected struct
	EXPECTED := []byte(`{"UID":"db-testUID10","FirstName":"Test10","LastName":"Update10","Weight":101,"WaistCirc":101.2,"HeightInches":101,"LeanBodyMass":101,"Age":101,"Gender":"female","Week":[{"Day":[{"Fat":101,"Carbs":101,"Protein":101,"TotalCalories":101,"DayCalories":"high","Weight":101,"Cardio":"missed","WeightTraining":"yes"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{"Fat":1,"Carbs":2,"Protein":3,"TotalCalories":4,"DayCalories":"low","Weight":5,"Cardio":"no","WeightTraining":"weightTrained"},{},{},{},{},{}]}],"Recommendation":[{"HighDayProtein":10,"HighDayCarb":11,"HighDayFat":12,"HighDayCalories":13,"NormalDayProtein":14,"NormalDayCarb":15,"NormalDayFat":16,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":24,"HIITChangeCardioIntervals":25,"ModifiedDate":"2020-09-13"},{},{"HighDayProtein":101,"HighDayCarb":101,"HighDayFat":101,"HighDayCalories":101,"NormalDayProtein":141,"NormalDayCarb":151,"NormalDayFat":161,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":241,"HIITChangeCardioIntervals":251,"ModifiedDate":"2020-09-99"},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}]}`)
	json.Unmarshal(EXPECTED, &expectedAuth)
	if !cmp.Equal(reqAuth, expectedAuth) {
		t.Errorf("Database returned unexpected body: \ngot -: %+v \nwant -: %+v", string(reqBody), string(EXPECTED))
		t.Fail()
	}
}

// another test inserting new day on empty week
func TestUpdateNewWeekDay2(t *testing.T) {
	//create control user in DB
	postBody := []byte(`{"UID":"db-testUID11","FirstName":"Test11","LastName":"Update11","Weight":101,"WaistCirc":101.2,"HeightInches":101,"LeanBodyMass":101,"Age":101,"Gender":"female","Week":[{"Day":[{"Fat":101,"Carbs":101,"Protein":101,"TotalCalories":101,"DayCalories":"high","Weight":101,"Cardio":"missed","WeightTraining":"yes"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"}]}],"Recommendation":[{"HighDayProtein":10,"HighDayCarb":11,"HighDayFat":12,"HighDayCalories":13,"NormalDayProtein":14,"NormalDayCarb":15,"NormalDayFat":16,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":24,"HIITChangeCardioIntervals":25,"ModifiedDate":"2020-09-13"},{},{"HighDayProtein":101,"HighDayCarb":101,"HighDayFat":101,"HighDayCalories":101,"NormalDayProtein":141,"NormalDayCarb":151,"NormalDayFat":161,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":241,"HIITChangeCardioIntervals":251,"ModifiedDate":"2020-09-111"}]}`)
	resp, err := http.Post("http://"+databaseAddress+"/userInfo/db-testUID11", "application/json; charset=UTF-8", bytes.NewBuffer(postBody))
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
	postBody = []byte(`{"Fat":1,"Carbs":2,"Protein":3,"TotalCalories":4,"DayCalories":"low","Weight":5,"Cardio":"no","WeightTraining":"weightTrained"}`)
	// post some data
	resp, err = http.Post("http://"+databaseAddress+"/userDaily/23/6/db-testUID11", "application/json; charset=UTF-8", bytes.NewBuffer(postBody))
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
	req, err := http.Get("http://"+databaseAddress+"/userInfo/db-testUID11")
	if err != nil {
		t.Errorf("Request failed: %v", err)
		t.Fail()
	}
	//unmarshal response into struct
	reqBody, _ := ioutil.ReadAll(req.Body)
	var reqAuth, expectedAuth structs.Client
	json.Unmarshal(reqBody, &reqAuth)
	//compare received struct with expected struct
	EXPECTED := []byte(`{"UID":"db-testUID11","FirstName":"Test11","LastName":"Update11","Weight":101,"WaistCirc":101.2,"HeightInches":101,"LeanBodyMass":101,"Age":101,"Gender":"female","Week":[{"Day":[{"Fat":101,"Carbs":101,"Protein":101,"TotalCalories":101,"DayCalories":"high","Weight":101,"Cardio":"missed","WeightTraining":"yes"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{"Fat":1,"Carbs":2,"Protein":3,"TotalCalories":4,"DayCalories":"low","Weight":5,"Cardio":"no","WeightTraining":"weightTrained"}]}],"Recommendation":[{"HighDayProtein":10,"HighDayCarb":11,"HighDayFat":12,"HighDayCalories":13,"NormalDayProtein":14,"NormalDayCarb":15,"NormalDayFat":16,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":24,"HIITChangeCardioIntervals":25,"ModifiedDate":"2020-09-13"},{},{"HighDayProtein":101,"HighDayCarb":101,"HighDayFat":101,"HighDayCalories":101,"NormalDayProtein":141,"NormalDayCarb":151,"NormalDayFat":161,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":241,"HIITChangeCardioIntervals":251,"ModifiedDate":"2020-09-111"},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}]}`)
	json.Unmarshal(EXPECTED, &expectedAuth)
	if !cmp.Equal(reqAuth, expectedAuth) {
		t.Errorf("Database returned unexpected body: \ngot -: %+v \nwant -: %+v", string(reqBody), string(EXPECTED))
		t.Fail()
	}
}
