package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"

	structs "../common"

	"github.com/google/go-cmp/cmp"
)

var frontendApiAddress string

// get service address from env
func TestMain(m *testing.M) {
	frontendApiAddress = os.Getenv("FRONTEND_API_SERVICE_ADDRESS")
	if frontendApiAddress == "" {
		frontendApiAddress = "localhost:8050"
	}
	log.Println("Testing Frontend-Api at address: " + frontendApiAddress)
	os.Exit(m.Run())
}

func TestAPIVersion(t *testing.T) {
	req, err := http.Get("http://" + frontendApiAddress + "/apiVersion")
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.FailNow()
	}
	// Check the response body is what we expect.
	expected := `{"apiVersion":1.0}`
	respBody, _ := ioutil.ReadAll(req.Body)
	if string(respBody) != expected {
		t.Errorf("Frontend-api returned unexpected body: got %v \nwant %v", string(respBody), expected)
		t.FailNow()
	}
}

//TestGetUserInfo get user profile and verify UID is stripped
func TestGetUserInfo(t *testing.T) {
	var UID = "testUID"
	req, _ := http.NewRequest("GET", "http://"+frontendApiAddress+"/userInfo", nil)
	// set session token header for request
	req.AddCookie(&http.Cookie{Name: "Authorization", Value: UID})
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.FailNow()
	}
	//unmarshal response into struct
	reqBody, _ := ioutil.ReadAll(resp.Body)
	var reqBackend, expectedBackend structs.Client
	json.Unmarshal(reqBody, &reqBackend)
	//check that UID is missing from response
	if reqBackend.UID != "" {
		t.Errorf("UID Should not be present: UID=%+v", reqBackend.UID)
		t.FailNow()
	}
	//compare received struct with expected struct
	EXPECTED := []byte(`{"FirstName":"Anthony","LastName":"Hanna","Weight":215,"WaistCirc":35,"HeightInches":75,"LeanBodyMass":15,"Age":20,"StartDate":"2020-08-15","Gender":"male","Week":[{"Day":[{"WaistCirc":20,"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalories":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":40,"Carbs":40,"Protein":40,"TotalCalories":34,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":100,"Carbs":100,"Protein":100,"TotalCalories":300,"DayCalories":"normal","Weight":321,"Cardio":"missed","WeightTraining":"no"}]},{"Day":[{"Fat":11,"Carbs":11,"Protein":11,"TotalCalories":31,"DayCalories":"normal","Weight":222,"Cardio":"missed","WeightTraining":"yes"},{"WaistCirc":20,"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{},{},{},{},{}]},{"Day":[{"Fat":110,"Carbs":110,"Protein":110,"TotalCalories":310,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]}],"Recommendation":[{"HighDayProtein":10,"HighDayCarb":11,"HighDayFat":12,"HighDayCalories":13,"NormalDayProtein":14,"NormalDayCarb":15,"NormalDayFat":16,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITCurrentCardioIntervals":24},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}]}`)
	json.Unmarshal(EXPECTED, &expectedBackend)
	if !cmp.Equal(reqBackend, expectedBackend) {
		t.Errorf("Database returned unexpected body: \ngot -: %+v \nwant -: %+v", string(reqBody), string(EXPECTED))
		t.FailNow()
	}
}

//TestGetInvalidUserInfo get empty user
func TestGetInvalidUserInfo(t *testing.T) {
	var UID = "testfail"
	req, _ := http.NewRequest("GET", "http://"+frontendApiAddress+"/userInfo", nil)
	// set session token header for request
	req.AddCookie(&http.Cookie{Name: "Authorization", Value: UID})
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.FailNow()
	}
	//compare received struct with expected struct
	if resp.StatusCode != 401 {
		t.Errorf("Status Incorrect, should be 401 -: http.StatusCode=%v", resp.StatusCode)
		t.FailNow()
	}
}

//TestCreateNewSimpleUser creates & validates new user
func TestCreateNewSimpleUser(t *testing.T) {
	var UID = "front-api-testUID1"
	//post data and validate that request succeeded
	var jsonStr = []byte(`{"FirstName":"Anthony1","LastName":"Hanna1","Weight":2151,"WaistCirc":35.51,"HeightInches":751,"LeanBodyMass":151,"Age":201,"Gender":"female"}`)
	req, _ := http.NewRequest("POST", "http://"+frontendApiAddress+"/userInfo", bytes.NewBuffer(jsonStr))
	// set session token header for request
	req.AddCookie(&http.Cookie{Name: "Authorization", Value: UID})
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.FailNow()
	}
	reqBody, _ := ioutil.ReadAll(resp.Body)
	//check that post response is "ok"
	if string(reqBody) != "ok" {
		t.Errorf("Post Request Response Incorrect (should be 'ok'): %v", string(reqBody))
		t.FailNow()
	}
	//========================================
	//Validate UserInfo is correct
	req2, _ := http.NewRequest("GET", "http://"+frontendApiAddress+"/userInfo", nil)
	// set session token header for request
	req2.AddCookie(&http.Cookie{Name: "Authorization", Value: UID})
	client2 := &http.Client{}
	resp2, err := client2.Do(req2)
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.FailNow()
	}
	//unmarshal response into struct
	reqBody2, _ := ioutil.ReadAll(resp2.Body)
	var reqBackend2, expectedBackend2 structs.Client
	json.Unmarshal(reqBody2, &reqBackend2)
	//check that UID is missing from response
	if reqBackend2.UID != "" {
		t.Errorf("UID Should not be present: UID=%+v", reqBackend2.UID)
		t.FailNow()
	}
	//compare received struct with expected struct
	EXPECTED := []byte(`{"FirstName":"Anthony1","LastName":"Hanna1","Weight":2151,"WaistCirc":35.51,"HeightInches":751,"LeanBodyMass":151,"Age":201,"Gender":"female","Week":[{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]}],"Recommendation":[{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}]}`)
	json.Unmarshal(EXPECTED, &expectedBackend2)
	if !cmp.Equal(reqBackend2, expectedBackend2) {
		t.Errorf("Database returned unexpected body: \ngot -: %+v \nwant -: %+v", string(reqBody2), string(EXPECTED))
		t.FailNow()
	}
}

//TestCreateNewSimpleUser2 creates & validates new user with start_date
func TestCreateNewSimpleUser2(t *testing.T) {
	var UID = "front-api-testUID1-2"
	//post data and validate that request succeeded
	var jsonStr = []byte(`{"FirstName":"Anthony1-2","LastName":"Hanna1-2","StartDate":"2222-08-15","Weight":2151,"WaistCirc":35.51,"HeightInches":751,"LeanBodyMass":151,"Age":201,"Gender":"female"}`)
	req, _ := http.NewRequest("POST", "http://"+frontendApiAddress+"/userInfo", bytes.NewBuffer(jsonStr))
	// set session token header for request
	req.AddCookie(&http.Cookie{Name: "Authorization", Value: UID})
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.FailNow()
	}
	reqBody, _ := ioutil.ReadAll(resp.Body)
	//check that post response is "ok"
	if string(reqBody) != "ok" {
		t.Errorf("Post Request Response Incorrect (should be 'ok'): %v", string(reqBody))
		t.FailNow()
	}
	//========================================
	//Validate UserInfo is correct
	req2, _ := http.NewRequest("GET", "http://"+frontendApiAddress+"/userInfo", nil)
	// set session token header for request
	req2.AddCookie(&http.Cookie{Name: "Authorization", Value: UID})
	client2 := &http.Client{}
	resp2, err := client2.Do(req2)
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.FailNow()
	}
	//unmarshal response into struct
	reqBody2, _ := ioutil.ReadAll(resp2.Body)
	var reqBackend2, expectedBackend2 structs.Client
	json.Unmarshal(reqBody2, &reqBackend2)
	//check that UID is missing from response
	if reqBackend2.UID != "" {
		t.Errorf("UID Should not be present: UID=%+v", reqBackend2.UID)
		t.FailNow()
	}
	//compare received struct with expected struct
	EXPECTED := []byte(`{"FirstName":"Anthony1-2","LastName":"Hanna1-2","Weight":2151,"WaistCirc":35.51,"HeightInches":751,"LeanBodyMass":151,"Age":201,"StartDate":"2222-08-15","Gender":"female","Week":[{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]}],"Recommendation":[{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}]}`)
	json.Unmarshal(EXPECTED, &expectedBackend2)
	if !cmp.Equal(reqBackend2, expectedBackend2) {
		t.Errorf("Database returned unexpected body: \ngot -: %+v \nwant -: %+v", string(reqBody2), string(EXPECTED))
		t.FailNow()
	}
}

//TestUpdateUserInfo verifies updating user data works correctly
func TestUpdateUserInfo(t *testing.T) {
	var UID = "front-api-testUID1-data"
	//post prop data for testing against
	var jsonStr = []byte(`{"FirstName":"Anthony1-data","LastName":"Hanna","Weight":100,"WaistCirc":55.5,"HeightInches":55,"LeanBodyMass":55,"Age":55,"Gender":"male","StartDate":"2020-10-30","Week":[{"Day":[{},{},{},{},{},{"Fat":1000,"Carbs":1000,"Protein":100,"TotalCalories":300,"DayCalories":"normal","Weight":321,"Cardio":"missed","WeightTraining":"no"},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{"Fat":1010,"Carbs":1010,"Protein":1010,"TotalCalories":1010,"DayCalories":"normal","Weight":10,"Cardio":"missed","WeightTraining":"yes"},{}]}],"Recommendation":[{"HighDayProtein":50,"HighDayCarb":11,"HighDayFat":12,"HighDayCalories":13,"NormalDayProtein":14,"NormalDayCarb":15,"NormalDayFat":16,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":24,"HIITChangeCardioIntervals":25,"ModifiedDate":"2020-09-15"},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}]}`)
	req, _ := http.NewRequest("POST", "http://"+frontendApiAddress+"/userInfo", bytes.NewBuffer(jsonStr))
	// set session token header for request
	req.AddCookie(&http.Cookie{Name: "Authorization", Value: UID})
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.FailNow()
	}
	reqBody, _ := ioutil.ReadAll(resp.Body)
	//check that post response is "ok"
	if string(reqBody) != "ok" {
		t.Errorf("Post Request Response Incorrect (should be 'ok'): %v", string(reqBody))
		t.FailNow()
	}
	//========================================
	//post update data
	var jsonStrNew = []byte(`{"FirstName":"Anthony1-333","LastName":"Hanna","Weight":100,"WaistCirc":55.5,"HeightInches":55,"LeanBodyMass":55,"Age":55,"Gender":"male","StartDate":"2033-10-30"}`)
	reqNew, _ := http.NewRequest("POST", "http://"+frontendApiAddress+"/userInfo", bytes.NewBuffer(jsonStrNew))
	// set session token header for request
	reqNew.AddCookie(&http.Cookie{Name: "Authorization", Value: UID})
	reqNew.Header.Set("Content-Type", "application/json")
	clientNew := &http.Client{}
	respNew, err := clientNew.Do(reqNew)
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.FailNow()
	}
	reqBodyNew, _ := ioutil.ReadAll(respNew.Body)
	//check that post response is "ok"
	if string(reqBodyNew) != "ok" {
		t.Errorf("Post Request Response Incorrect (should be 'ok'): %v", string(reqBodyNew))
		t.FailNow()
	}
	//========================================
	//Validate UserInfo is correct
	req2, _ := http.NewRequest("GET", "http://"+frontendApiAddress+"/userInfo", nil)
	// set session token header for request
	req2.AddCookie(&http.Cookie{Name: "Authorization", Value: UID})
	client2 := &http.Client{}
	resp2, err := client2.Do(req2)
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.FailNow()
	}
	//unmarshal response into struct
	reqBody2, _ := ioutil.ReadAll(resp2.Body)
	var reqBackend2, expectedBackend2 structs.Client
	json.Unmarshal(reqBody2, &reqBackend2)
	//check that UID is missing from response
	if reqBackend2.UID != "" {
		t.Errorf("UID Should not be present: UID=%+v", reqBackend2.UID)
		t.FailNow()
	}
	//compare received struct with expected struct
	EXPECTED := []byte(`{"FirstName":"Anthony1-333","LastName":"Hanna","Weight":100,"WaistCirc":55.5,"HeightInches":55,"LeanBodyMass":55,"Age":55,"StartDate":"2033-10-30","Gender":"male","Week":[{"Day":[{},{},{},{},{},{"Fat":1000,"Carbs":1000,"Protein":100,"TotalCalories":300,"DayCalories":"normal","Weight":321,"Cardio":"missed","WeightTraining":"no"},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{"Fat":1010,"Carbs":1010,"Protein":1010,"TotalCalories":1010,"DayCalories":"normal","Weight":10,"Cardio":"missed","WeightTraining":"yes"},{}]}],"Recommendation":[{"HighDayProtein":50,"HighDayCarb":11,"HighDayFat":12,"HighDayCalories":13,"NormalDayProtein":14,"NormalDayCarb":15,"NormalDayFat":16,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":24,"HIITChangeCardioIntervals":25,"ModifiedDate":"2020-09-15"},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}]} `)
	json.Unmarshal(EXPECTED, &expectedBackend2)
	if !cmp.Equal(reqBackend2, expectedBackend2) {
		t.Errorf("Database returned unexpected body: \ngot -: %+v \nwant -: %+v", string(reqBody2), string(EXPECTED))
		t.FailNow()
	}
}

//TestUpdateUserInfoDate verifies updating only user date works correctly
func TestUpdateUserInfoDate(t *testing.T) {
	var UID = "front-api-testUID1-date"
	//post prop data for testing against
	var jsonStr = []byte(`{"FirstName":"Anthony1-date","LastName":"Hanna","Weight":100,"WaistCirc":55.5,"HeightInches":55,"LeanBodyMass":55,"Age":55,"Gender":"male","StartDate":"2020-10-30","Week":[{"Day":[{},{},{},{},{},{"Fat":1000,"Carbs":1000,"Protein":100,"TotalCalories":300,"DayCalories":"normal","Weight":321,"Cardio":"missed","WeightTraining":"no"},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{"Fat":1010,"Carbs":1010,"Protein":1010,"TotalCalories":1010,"DayCalories":"normal","Weight":10,"Cardio":"missed","WeightTraining":"yes"},{}]}],"Recommendation":[{"HighDayProtein":50,"HighDayCarb":11,"HighDayFat":12,"HighDayCalories":13,"NormalDayProtein":14,"NormalDayCarb":15,"NormalDayFat":16,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":24,"HIITChangeCardioIntervals":25,"ModifiedDate":"2020-09-15"},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}]}`)
	req, _ := http.NewRequest("POST", "http://"+frontendApiAddress+"/userInfo", bytes.NewBuffer(jsonStr))
	// set session token header for request
	req.AddCookie(&http.Cookie{Name: "Authorization", Value: UID})
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.FailNow()
	}
	reqBody, _ := ioutil.ReadAll(resp.Body)
	//check that post response is "ok"
	if string(reqBody) != "ok" {
		t.Errorf("Post Request Response Incorrect (should be 'ok'): %v", string(reqBody))
		t.FailNow()
	}
	//========================================
	//post update data
	var jsonStrNew = []byte(`{"StartDate":"3333-10-30"}`)
	reqNew, _ := http.NewRequest("POST", "http://"+frontendApiAddress+"/userInfo", bytes.NewBuffer(jsonStrNew))
	// set session token header for request
	reqNew.AddCookie(&http.Cookie{Name: "Authorization", Value: UID})
	reqNew.Header.Set("Content-Type", "application/json")
	clientNew := &http.Client{}
	respNew, err := clientNew.Do(reqNew)
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.FailNow()
	}
	reqBodyNew, _ := ioutil.ReadAll(respNew.Body)
	//check that post response is "ok"
	if string(reqBodyNew) != "ok" {
		t.Errorf("Post Request Response Incorrect (should be 'ok'): %v", string(reqBodyNew))
		t.FailNow()
	}
	//========================================
	//Validate UserInfo is correct
	req2, _ := http.NewRequest("GET", "http://"+frontendApiAddress+"/userInfo", nil)
	// set session token header for request
	req2.AddCookie(&http.Cookie{Name: "Authorization", Value: UID})
	client2 := &http.Client{}
	resp2, err := client2.Do(req2)
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.FailNow()
	}
	//unmarshal response into struct
	reqBody2, _ := ioutil.ReadAll(resp2.Body)
	var reqBackend2, expectedBackend2 structs.Client
	json.Unmarshal(reqBody2, &reqBackend2)
	//check that UID is missing from response
	if reqBackend2.UID != "" {
		t.Errorf("UID Should not be present: UID=%+v", reqBackend2.UID)
		t.FailNow()
	}
	//compare received struct with expected struct
	EXPECTED := []byte(`{"FirstName":"Anthony1-date","LastName":"Hanna","Weight":100,"WaistCirc":55.5,"HeightInches":55,"LeanBodyMass":55,"Age":55,"StartDate":"3333-10-30","Gender":"male","Week":[{"Day":[{},{},{},{},{},{"Fat":1000,"Carbs":1000,"Protein":100,"TotalCalories":300,"DayCalories":"normal","Weight":321,"Cardio":"missed","WeightTraining":"no"},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{"Fat":1010,"Carbs":1010,"Protein":1010,"TotalCalories":1010,"DayCalories":"normal","Weight":10,"Cardio":"missed","WeightTraining":"yes"},{}]}],"Recommendation":[{"HighDayProtein":50,"HighDayCarb":11,"HighDayFat":12,"HighDayCalories":13,"NormalDayProtein":14,"NormalDayCarb":15,"NormalDayFat":16,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":24,"HIITChangeCardioIntervals":25,"ModifiedDate":"2020-09-15"},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}]} `)
	json.Unmarshal(EXPECTED, &expectedBackend2)
	if !cmp.Equal(reqBackend2, expectedBackend2) {
		t.Errorf("Database returned unexpected body: \ngot -: %+v \nwant -: %+v", string(reqBody2), string(EXPECTED))
		t.FailNow()
	}
}

//TestCreateNewComplexUser creates & validates new user
func TestCreateNewComplexUser(t *testing.T) {
	var UID = "front-api-testUID2"
	//post data and validate that request succeeded
	var jsonStr = []byte(`{"FirstName":"Anthony2","LastName":"Hanna2","Weight":2152,"WaistCirc":35.52,"HeightInches":75,"LeanBodyMass":15,"Age":20,"Gender":"male","Week":[{"Day":[{},{},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalories":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{},{},{},{}]},{"Day":[{"Fat":112,"Carbs":112,"Protein":112,"TotalCalories":312,"DayCalories":"normal","Weight":2222,"Cardio":"missed","WeightTraining":"yes"},{},{},{},{},{},{}]},{"Day":[{},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]}],"Recommendation":[{},{"HighDayProtein":102,"HighDayCarb":112,"HighDayFat":122,"HighDayCalories":132,"NormalDayProtein":142,"NormalDayCarb":152,"NormalDayFat":16,"NormalDayCalories":17,"LowDayProtein":182,"LowDayCarb":192,"LowDayFat":202,"LowDayCalories":212,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":242,"HIITChangeCardioIntervals":252,"ModifiedDate":"2020-09-12"},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}]}`)
	req, _ := http.NewRequest("POST", "http://"+frontendApiAddress+"/userInfo", bytes.NewBuffer(jsonStr))
	// set session token header for request
	req.AddCookie(&http.Cookie{Name: "Authorization", Value: UID})
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.FailNow()
	}
	reqBody, _ := ioutil.ReadAll(resp.Body)
	//check that post response is "ok"
	if string(reqBody) != "ok" {
		t.Errorf("Post Request Response Incorrect (should be 'ok'): %v", string(reqBody))
		t.FailNow()
	}
	//========================================
	//Validate UserInfo is correct
	req2, _ := http.NewRequest("GET", "http://"+frontendApiAddress+"/userInfo", nil)
	// set session token header for request
	req2.AddCookie(&http.Cookie{Name: "Authorization", Value: UID})
	client2 := &http.Client{}
	resp2, err := client2.Do(req2)
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.FailNow()
	}
	//unmarshal response into struct
	reqBody2, _ := ioutil.ReadAll(resp2.Body)
	var reqBackend2, expectedBackend2 structs.Client
	json.Unmarshal(reqBody2, &reqBackend2)
	//check that UID is missing from response
	if reqBackend2.UID != "" {
		t.Errorf("UID Should not be present: UID=%+v", reqBackend2.UID)
		t.FailNow()
	}
	//compare received struct with expected struct
	EXPECTED := []byte(`{"FirstName":"Anthony2","LastName":"Hanna2","Weight":2152,"WaistCirc":35.52,"HeightInches":75,"LeanBodyMass":15,"Age":20,"Gender":"male","Week":[{"Day":[{},{},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalories":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{},{},{},{}]},{"Day":[{"Fat":112,"Carbs":112,"Protein":112,"TotalCalories":312,"DayCalories":"normal","Weight":2222,"Cardio":"missed","WeightTraining":"yes"},{},{},{},{},{},{}]},{"Day":[{},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]}],"Recommendation":[{},{"HighDayProtein":102,"HighDayCarb":112,"HighDayFat":122,"HighDayCalories":132,"NormalDayProtein":142,"NormalDayCarb":152,"NormalDayFat":16,"NormalDayCalories":17,"LowDayProtein":182,"LowDayCarb":192,"LowDayFat":202,"LowDayCalories":212,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":242,"HIITChangeCardioIntervals":252,"ModifiedDate":"2020-09-12"},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}]}`)
	json.Unmarshal(EXPECTED, &expectedBackend2)
	if !cmp.Equal(reqBackend2, expectedBackend2) {
		t.Errorf("Database returned unexpected body: \ngot -: %+v \nwant -: %+v", string(reqBody2), string(EXPECTED))
		t.FailNow()
	}
}

//TestUpdateUserWeekly validates updating weekly data
func TestUpdateUserWeekly(t *testing.T) {
	var UID = "front-api-testUID3"
	//post prop data for testing against
	var jsonStr = []byte(`{"FirstName":"Anthony3","LastName":"Hanna","Weight":215,"WaistCirc":35.5,"HeightInches":75,"LeanBodyMass":15,"Age":20,"Gender":"male","Week":[{"Day":[{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalories":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":40,"Carbs":40,"Protein":40,"TotalCalories":34,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":100,"Carbs":100,"Protein":100,"TotalCalories":300,"DayCalories":"normal","Weight":321,"Cardio":"missed","WeightTraining":"no"}]},{"Day":[{"Fat":11,"Carbs":11,"Protein":11,"TotalCalories":31,"DayCalories":"normal","Weight":222,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{},{},{},{},{}]},{"Day":[{"Fat":110,"Carbs":110,"Protein":110,"TotalCalories":310,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]}],"Recommendation":[{"HighDayProtein":10,"HighDayCarb":11,"HighDayFat":12,"HighDayCalories":13,"NormalDayProtein":14,"NormalDayCarb":15,"NormalDayFat":16,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":24,"HIITChangeCardioIntervals":25,"ModifiedDate":"2020-09-13"},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}]}`)
	req, _ := http.NewRequest("POST", "http://"+frontendApiAddress+"/userInfo", bytes.NewBuffer(jsonStr))
	// set session token header for request
	req.AddCookie(&http.Cookie{Name: "Authorization", Value: UID})
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.FailNow()
	}
	reqBody, _ := ioutil.ReadAll(resp.Body)
	//check that post response is "ok"
	if string(reqBody) != "ok" {
		t.Errorf("Post Request Response Incorrect (should be 'ok'): %v", string(reqBody))
		t.FailNow()
	}
	//========================================
	//post update data
	var jsonStrNew = []byte(`{"Day":[{},{},{},{"Fat":333,"Carbs":333,"Protein":333,"TotalCalories":333,"DayCalories":"normal","Weight":333,"Cardio":"missed","WeightTraining":"no"},{},{"Fat":3333,"Carbs":3333,"Protein":3333,"TotalCalories":3333,"DayCalories":"normal","Weight":3333,"Cardio":"missed","WeightTraining":"no"},{}]}`)
	reqNew, _ := http.NewRequest("POST", "http://"+frontendApiAddress+"/userWeekly/2", bytes.NewBuffer(jsonStrNew))
	// set session token header for request
	reqNew.AddCookie(&http.Cookie{Name: "Authorization", Value: UID})
	reqNew.Header.Set("Content-Type", "application/json")
	clientNew := &http.Client{}
	respNew, err := clientNew.Do(reqNew)
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.FailNow()
	}
	reqBodyNew, _ := ioutil.ReadAll(respNew.Body)
	//check that post response is "ok"
	if string(reqBodyNew) != "ok" {
		t.Errorf("Post Request Response Incorrect (should be 'ok'): %v", string(reqBodyNew))
		t.FailNow()
	}
	//========================================
	//Validate UserInfo is correct
	req2, _ := http.NewRequest("GET", "http://"+frontendApiAddress+"/userInfo", nil)
	// set session token header for request
	req2.AddCookie(&http.Cookie{Name: "Authorization", Value: UID})
	client2 := &http.Client{}
	resp2, err := client2.Do(req2)
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.FailNow()
	}
	//unmarshal response into struct
	reqBody2, _ := ioutil.ReadAll(resp2.Body)
	var reqBackend2, expectedBackend2 structs.Client
	json.Unmarshal(reqBody2, &reqBackend2)
	//check that UID is missing from response
	if reqBackend2.UID != "" {
		t.Errorf("UID Should not be present: UID=%+v", reqBackend2.UID)
		t.FailNow()
	}
	//compare received struct with expected struct
	EXPECTED := []byte(`{"FirstName":"Anthony3","LastName":"Hanna","Weight":215,"WaistCirc":35.5,"HeightInches":75,"LeanBodyMass":15,"Age":20,"Gender":"male","Week":[{"Day":[{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalories":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":40,"Carbs":40,"Protein":40,"TotalCalories":34,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":100,"Carbs":100,"Protein":100,"TotalCalories":300,"DayCalories":"normal","Weight":321,"Cardio":"missed","WeightTraining":"no"}]},{"Day":[{"Fat":11,"Carbs":11,"Protein":11,"TotalCalories":31,"DayCalories":"normal","Weight":222,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{},{},{},{},{}]},{"Day":[{},{},{},{"Fat":333,"Carbs":333,"Protein":333,"TotalCalories":333,"DayCalories":"normal","Weight":333,"Cardio":"missed","WeightTraining":"no"},{},{"Fat":3333,"Carbs":3333,"Protein":3333,"TotalCalories":3333,"DayCalories":"normal","Weight":3333,"Cardio":"missed","WeightTraining":"no"},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]}],"Recommendation":[{"HighDayProtein":10,"HighDayCarb":11,"HighDayFat":12,"HighDayCalories":13,"NormalDayProtein":14,"NormalDayCarb":15,"NormalDayFat":16,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":24,"HIITChangeCardioIntervals":25,"ModifiedDate":"2020-09-13"},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}]}`)
	json.Unmarshal(EXPECTED, &expectedBackend2)
	if !cmp.Equal(reqBackend2, expectedBackend2) {
		t.Errorf("Database returned unexpected body: \ngot -: %+v \nwant -: %+v", string(reqBody2), string(EXPECTED))
		t.FailNow()
	}
}

//TestUpdateUserWeekly2 validates updating weekly data again
func TestUpdateUserWeekly2(t *testing.T) {
	var UID = "front-api-testUID4"
	//post prop data for testing against
	var jsonStr = []byte(`{"FirstName":"Anthony4","LastName":"Hanna","Weight":215,"WaistCirc":35.5,"HeightInches":75,"LeanBodyMass":15,"Age":20,"Gender":"male","Week":[{"Day":[{},{},{},{},{},{"Fat":100,"Carbs":100,"Protein":100,"TotalCalories":300,"DayCalories":"normal","Weight":321,"Cardio":"missed","WeightTraining":"no"},{}]},{"Day":[{"Fat":11,"Carbs":11,"Protein":11,"TotalCalories":31,"DayCalories":"normal","Weight":222,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{},{},{},{},{}]},{"Day":[{},{},{},{"Fat":123,"Carbs":333,"Protein":333,"TotalCalories":333,"DayCalories":"normal","Weight":333,"Cardio":"missed","WeightTraining":"no"},{},{"Fat":3333,"Carbs":3333,"Protein":3333,"TotalCalories":3333,"DayCalories":"normal","Weight":3333,"Cardio":"missed","WeightTraining":"no"},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{"Fat":1010,"Carbs":1010,"Protein":1010,"TotalCalories":1010,"DayCalories":"normal","Weight":10,"Cardio":"missed","WeightTraining":"yes"},{}]}],"Recommendation":[{"HighDayProtein":10,"HighDayCarb":11,"HighDayFat":12,"HighDayCalories":13,"NormalDayProtein":14,"NormalDayCarb":15,"NormalDayFat":16,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":24,"HIITChangeCardioIntervals":25,"ModifiedDate":"2020-09-13"},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}]}`)
	req, _ := http.NewRequest("POST", "http://"+frontendApiAddress+"/userInfo", bytes.NewBuffer(jsonStr))
	// set session token header for request
	req.AddCookie(&http.Cookie{Name: "Authorization", Value: UID})
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.FailNow()
	}
	reqBody, _ := ioutil.ReadAll(resp.Body)
	//check that post response is "ok"
	if string(reqBody) != "ok" {
		t.Errorf("Post Request Response Incorrect (should be 'ok'): %v", string(reqBody))
		t.FailNow()
	}
	//========================================
	//post update data
	var jsonStrNew = []byte(`{"Day":[{},{"Fat":444,"Carbs":444,"Protein":444,"TotalCalories":444,"DayCalories":"normal","Weight":444,"Cardio":"missed","WeightTraining":"yes"},{},{"Fat":333,"Carbs":333,"Protein":333,"TotalCalories":333,"DayCalories":"normal","Weight":333,"Cardio":"missed","WeightTraining":"no"},{},{},{"Fat":3333,"Carbs":3333,"Protein":3333,"TotalCalories":3333,"DayCalories":"normal","Weight":3333,"Cardio":"missed","WeightTraining":"no"}]}`)
	reqNew, _ := http.NewRequest("POST", "http://"+frontendApiAddress+"/userWeekly/23", bytes.NewBuffer(jsonStrNew))
	// set session token header for request
	reqNew.AddCookie(&http.Cookie{Name: "Authorization", Value: UID})
	reqNew.Header.Set("Content-Type", "application/json")
	clientNew := &http.Client{}
	respNew, err := clientNew.Do(reqNew)
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.FailNow()
	}
	reqBodyNew, _ := ioutil.ReadAll(respNew.Body)
	//check that post response is "ok"
	if string(reqBodyNew) != "ok" {
		t.Errorf("Post Request Response Incorrect (should be 'ok'): %v", string(reqBodyNew))
		t.FailNow()
	}
	//========================================
	//Validate UserInfo is correct
	req2, _ := http.NewRequest("GET", "http://"+frontendApiAddress+"/userInfo", nil)
	// set session token header for request
	req2.AddCookie(&http.Cookie{Name: "Authorization", Value: UID})
	client2 := &http.Client{}
	resp2, err := client2.Do(req2)
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.FailNow()
	}
	//unmarshal response into struct
	reqBody2, _ := ioutil.ReadAll(resp2.Body)
	var reqBackend2, expectedBackend2 structs.Client
	json.Unmarshal(reqBody2, &reqBackend2)
	//check that UID is missing from response
	if reqBackend2.UID != "" {
		t.Errorf("UID Should not be present: UID=%+v", reqBackend2.UID)
		t.FailNow()
	}
	//compare received struct with expected struct
	EXPECTED := []byte(`{"FirstName":"Anthony4","LastName":"Hanna","Weight":215,"WaistCirc":35.5,"HeightInches":75,"LeanBodyMass":15,"Age":20,"Gender":"male","Week":[{"Day":[{},{},{},{},{},{"Fat":100,"Carbs":100,"Protein":100,"TotalCalories":300,"DayCalories":"normal","Weight":321,"Cardio":"missed","WeightTraining":"no"},{}]},{"Day":[{"Fat":11,"Carbs":11,"Protein":11,"TotalCalories":31,"DayCalories":"normal","Weight":222,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{},{},{},{},{}]},{"Day":[{},{},{},{"Fat":123,"Carbs":333,"Protein":333,"TotalCalories":333,"DayCalories":"normal","Weight":333,"Cardio":"missed","WeightTraining":"no"},{},{"Fat":3333,"Carbs":3333,"Protein":3333,"TotalCalories":3333,"DayCalories":"normal","Weight":3333,"Cardio":"missed","WeightTraining":"no"},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{"Fat":444,"Carbs":444,"Protein":444,"TotalCalories":444,"DayCalories":"normal","Weight":444,"Cardio":"missed","WeightTraining":"yes"},{},{"Fat":333,"Carbs":333,"Protein":333,"TotalCalories":333,"DayCalories":"normal","Weight":333,"Cardio":"missed","WeightTraining":"no"},{},{},{"Fat":3333,"Carbs":3333,"Protein":3333,"TotalCalories":3333,"DayCalories":"normal","Weight":3333,"Cardio":"missed","WeightTraining":"no"}]}],"Recommendation":[{"HighDayProtein":10,"HighDayCarb":11,"HighDayFat":12,"HighDayCalories":13,"NormalDayProtein":14,"NormalDayCarb":15,"NormalDayFat":16,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":24,"HIITChangeCardioIntervals":25,"ModifiedDate":"2020-09-13"},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}]}`)
	json.Unmarshal(EXPECTED, &expectedBackend2)
	if !cmp.Equal(reqBackend2, expectedBackend2) {
		t.Errorf("Database returned unexpected body: \ngot -: %+v \nwant -: %+v", string(reqBody2), string(EXPECTED))
		t.FailNow()
	}
}

//TestUpdateUserDaily validates updating daily data
func TestUpdateUserDaily(t *testing.T) {
	var UID = "front-api-testUID5"
	//post prop data for testing against
	var jsonStr = []byte(`{"FirstName":"Anthony5","LastName":"Hanna","Weight":555,"WaistCirc":55.5,"HeightInches":55,"LeanBodyMass":55,"Age":55,"Gender":"male","Week":[{"Day":[{},{},{},{},{},{"Fat":500,"Carbs":500,"Protein":500,"TotalCalories":300,"DayCalories":"normal","Weight":321,"Cardio":"missed","WeightTraining":"no"},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{"Fat":5050,"Carbs":5050,"Protein":5050,"TotalCalories":5050,"DayCalories":"normal","Weight":10,"Cardio":"missed","WeightTraining":"yes"},{}]}],"Recommendation":[{"HighDayProtein":50,"HighDayCarb":11,"HighDayFat":12,"HighDayCalories":13,"NormalDayProtein":14,"NormalDayCarb":15,"NormalDayFat":16,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":24,"HIITChangeCardioIntervals":25,"ModifiedDate":"2020-09-15"},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}]}`)
	req, _ := http.NewRequest("POST", "http://"+frontendApiAddress+"/userInfo", bytes.NewBuffer(jsonStr))
	// set session token header for request
	req.AddCookie(&http.Cookie{Name: "Authorization", Value: UID})
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.FailNow()
	}
	reqBody, _ := ioutil.ReadAll(resp.Body)
	//check that post response is "ok"
	if string(reqBody) != "ok" {
		t.Errorf("Post Request Response Incorrect (should be 'ok'): %v", string(reqBody))
		t.FailNow()
	}
	//========================================
	//post update data
	var jsonStrNew = []byte(`{"Fat":555,"Carbs":555,"Protein":555,"TotalCalories":555,"DayCalories":"normal","Weight":555,"Cardio":"missed","WeightTraining":"no"}`)
	reqNew, _ := http.NewRequest("POST", "http://"+frontendApiAddress+"/userDaily/1/1", bytes.NewBuffer(jsonStrNew))
	// set session token header for request
	reqNew.AddCookie(&http.Cookie{Name: "Authorization", Value: UID})
	reqNew.Header.Set("Content-Type", "application/json")
	clientNew := &http.Client{}
	respNew, err := clientNew.Do(reqNew)
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.FailNow()
	}
	reqBodyNew, _ := ioutil.ReadAll(respNew.Body)
	//check that post response is "ok"
	if string(reqBodyNew) != "ok" {
		t.Errorf("Post Request Response Incorrect (should be 'ok'): %v", string(reqBodyNew))
		t.FailNow()
	}
	//========================================
	//Validate UserInfo is correct
	req2, _ := http.NewRequest("GET", "http://"+frontendApiAddress+"/userInfo", nil)
	// set session token header for request
	req2.AddCookie(&http.Cookie{Name: "Authorization", Value: UID})
	client2 := &http.Client{}
	resp2, err := client2.Do(req2)
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.FailNow()
	}
	//unmarshal response into struct
	reqBody2, _ := ioutil.ReadAll(resp2.Body)
	var reqBackend2, expectedBackend2 structs.Client
	json.Unmarshal(reqBody2, &reqBackend2)
	//check that UID is missing from response
	if reqBackend2.UID != "" {
		t.Errorf("UID Should not be present: UID=%+v", reqBackend2.UID)
		t.FailNow()
	}
	//compare received struct with expected struct
	EXPECTED := []byte(`{"FirstName":"Anthony5","LastName":"Hanna","Weight":555,"WaistCirc":55.5,"HeightInches":55,"LeanBodyMass":55,"Age":55,"Gender":"male","Week":[{"Day":[{},{},{},{},{},{"Fat":500,"Carbs":500,"Protein":500,"TotalCalories":300,"DayCalories":"normal","Weight":321,"Cardio":"missed","WeightTraining":"no"},{}]},{"Day":[{},{"Fat":555,"Carbs":555,"Protein":555,"TotalCalories":555,"DayCalories":"normal","Weight":555,"Cardio":"missed","WeightTraining":"no"},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{"Fat":5050,"Carbs":5050,"Protein":5050,"TotalCalories":5050,"DayCalories":"normal","Weight":10,"Cardio":"missed","WeightTraining":"yes"},{}]}],"Recommendation":[{"HighDayProtein":50,"HighDayCarb":11,"HighDayFat":12,"HighDayCalories":13,"NormalDayProtein":14,"NormalDayCarb":15,"NormalDayFat":16,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":24,"HIITChangeCardioIntervals":25,"ModifiedDate":"2020-09-15"},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}]}`)
	json.Unmarshal(EXPECTED, &expectedBackend2)
	if !cmp.Equal(reqBackend2, expectedBackend2) {
		t.Errorf("Database returned unexpected body: \ngot -: %+v \nwant -: %+v", string(reqBody2), string(EXPECTED))
		t.FailNow()
	}
}

//TestUpdateUserDailyNext validates updating daily data immediately following previous day
func TestUpdateUserDailyNext(t *testing.T) {
	var UID = "front-api-testUID6"
	//post prop data for testing against
	var jsonStr = []byte(`{"FirstName":"Anthony6","LastName":"Hanna","Weight":555,"WaistCirc":55.5,"HeightInches":55,"LeanBodyMass":55,"Age":55,"Gender":"male","Week":[{"Day":[{},{},{},{},{},{"Fat":500,"Carbs":500,"Protein":500,"TotalCalories":300,"DayCalories":"normal","Weight":321,"Cardio":"missed","WeightTraining":"no"},{}]},{"Day":[{},{"Fat":555,"Carbs":555,"Protein":555,"TotalCalories":555,"DayCalories":"normal","Weight":555,"Cardio":"missed","WeightTraining":"no"},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{"Fat":5050,"Carbs":5050,"Protein":5050,"TotalCalories":5050,"DayCalories":"normal","Weight":10,"Cardio":"missed","WeightTraining":"yes"},{}]}],"Recommendation":[{"HighDayProtein":50,"HighDayCarb":11,"HighDayFat":12,"HighDayCalories":13,"NormalDayProtein":14,"NormalDayCarb":15,"NormalDayFat":16,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":24,"HIITChangeCardioIntervals":25,"ModifiedDate":"2020-09-15"},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}]}`)
	req, _ := http.NewRequest("POST", "http://"+frontendApiAddress+"/userInfo", bytes.NewBuffer(jsonStr))
	// set session token header for request
	req.AddCookie(&http.Cookie{Name: "Authorization", Value: UID})
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.FailNow()
	}
	reqBody, _ := ioutil.ReadAll(resp.Body)
	//check that post response is "ok"
	if string(reqBody) != "ok" {
		t.Errorf("Post Request Response Incorrect (should be 'ok'): %v", string(reqBody))
		t.FailNow()
	}
	//========================================
	//post update data
	var jsonStrNew = []byte(`{"Fat":666,"Carbs":666,"Protein":666,"TotalCalories":666,"DayCalories":"normal","Weight":666,"Cardio":"missed","WeightTraining":"yes"}`)
	reqNew, _ := http.NewRequest("POST", "http://"+frontendApiAddress+"/userDaily/1/2", bytes.NewBuffer(jsonStrNew))
	// set session token header for request
	reqNew.AddCookie(&http.Cookie{Name: "Authorization", Value: UID})
	reqNew.Header.Set("Content-Type", "application/json")
	clientNew := &http.Client{}
	respNew, err := clientNew.Do(reqNew)
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.FailNow()
	}
	reqBodyNew, _ := ioutil.ReadAll(respNew.Body)
	//check that post response is "ok"
	if string(reqBodyNew) != "ok" {
		t.Errorf("Post Request Response Incorrect (should be 'ok'): %v", string(reqBodyNew))
		t.FailNow()
	}
	//========================================
	//Validate UserInfo is correct
	req2, _ := http.NewRequest("GET", "http://"+frontendApiAddress+"/userInfo", nil)
	// set session token header for request
	req2.AddCookie(&http.Cookie{Name: "Authorization", Value: UID})
	client2 := &http.Client{}
	resp2, err := client2.Do(req2)
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.FailNow()
	}
	//unmarshal response into struct
	reqBody2, _ := ioutil.ReadAll(resp2.Body)
	var reqBackend2, expectedBackend2 structs.Client
	json.Unmarshal(reqBody2, &reqBackend2)
	//check that UID is missing from response
	if reqBackend2.UID != "" {
		t.Errorf("UID Should not be present: UID=%+v", reqBackend2.UID)
		t.FailNow()
	}
	//compare received struct with expected struct
	EXPECTED := []byte(`{"FirstName":"Anthony6","LastName":"Hanna","Weight":555,"WaistCirc":55.5,"HeightInches":55,"LeanBodyMass":55,"Age":55,"Gender":"male","Week":[{"Day":[{},{},{},{},{},{"Fat":500,"Carbs":500,"Protein":500,"TotalCalories":300,"DayCalories":"normal","Weight":321,"Cardio":"missed","WeightTraining":"no"},{}]},{"Day":[{},{"Fat":555,"Carbs":555,"Protein":555,"TotalCalories":555,"DayCalories":"normal","Weight":555,"Cardio":"missed","WeightTraining":"no"},{"Fat":666,"Carbs":666,"Protein":666,"TotalCalories":666,"DayCalories":"normal","Weight":666,"Cardio":"missed","WeightTraining":"yes"},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{"Fat":5050,"Carbs":5050,"Protein":5050,"TotalCalories":5050,"DayCalories":"normal","Weight":10,"Cardio":"missed","WeightTraining":"yes"},{}]}],"Recommendation":[{"HighDayProtein":50,"HighDayCarb":11,"HighDayFat":12,"HighDayCalories":13,"NormalDayProtein":14,"NormalDayCarb":15,"NormalDayFat":16,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":24,"HIITChangeCardioIntervals":25,"ModifiedDate":"2020-09-15"},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}]}`)
	json.Unmarshal(EXPECTED, &expectedBackend2)
	if !cmp.Equal(reqBackend2, expectedBackend2) {
		t.Errorf("Database returned unexpected body: \ngot -: %+v \nwant -: %+v", string(reqBody2), string(EXPECTED))
		t.FailNow()
	}
}

//TestUpdateUserDailyOverwrite validates updating daily data immediately following previous day
func TestUpdateUserDailyOverwrite(t *testing.T) {
	var UID = "front-api-testUID7"
	//post prop data for testing against
	var jsonStr = []byte(`{"FirstName":"Anthony7","LastName":"Hanna","Weight":555,"WaistCirc":55.5,"HeightInches":55,"LeanBodyMass":55,"Age":55,"Gender":"male","Week":[{"Day":[{},{},{},{},{},{"Fat":500,"Carbs":500,"Protein":500,"TotalCalories":300,"DayCalories":"normal","Weight":321,"Cardio":"missed","WeightTraining":"no"},{}]},{"Day":[{},{"Fat":555,"Carbs":555,"Protein":555,"TotalCalories":555,"DayCalories":"normal","Weight":555,"Cardio":"missed","WeightTraining":"no"},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{"Fat":5050,"Carbs":5050,"Protein":5050,"TotalCalories":5050,"DayCalories":"normal","Weight":10,"Cardio":"missed","WeightTraining":"yes"},{"Fat":9999,"Carbs":9999,"Protein":9999,"TotalCalories":9999,"DayCalories":"low","Weight":9999,"Cardio":"missed","WeightTraining":"no"}]}],"Recommendation":[{"HighDayProtein":50,"HighDayCarb":11,"HighDayFat":12,"HighDayCalories":13,"NormalDayProtein":14,"NormalDayCarb":15,"NormalDayFat":16,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":24,"HIITChangeCardioIntervals":25,"ModifiedDate":"2020-09-15"},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}]}`)
	req, _ := http.NewRequest("POST", "http://"+frontendApiAddress+"/userInfo", bytes.NewBuffer(jsonStr))
	// set session token header for request
	req.AddCookie(&http.Cookie{Name: "Authorization", Value: UID})
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.FailNow()
	}
	reqBody, _ := ioutil.ReadAll(resp.Body)
	//check that post response is "ok"
	if string(reqBody) != "ok" {
		t.Errorf("Post Request Response Incorrect (should be 'ok'): %v", string(reqBody))
		t.FailNow()
	}
	//========================================
	//post update data
	var jsonStrNew = []byte(`{"Fat":777,"Carbs":777,"Protein":777,"TotalCalories":777,"DayCalories":"normal","Weight":777,"Cardio":"missed","WeightTraining":"no"}`)
	reqNew, _ := http.NewRequest("POST", "http://"+frontendApiAddress+"/userDaily/23/6", bytes.NewBuffer(jsonStrNew))
	// set session token header for request
	reqNew.AddCookie(&http.Cookie{Name: "Authorization", Value: UID})
	reqNew.Header.Set("Content-Type", "application/json")
	clientNew := &http.Client{}
	respNew, err := clientNew.Do(reqNew)
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.FailNow()
	}
	reqBodyNew, _ := ioutil.ReadAll(respNew.Body)
	//check that post response is "ok"
	if string(reqBodyNew) != "ok" {
		t.Errorf("Post Request Response Incorrect (should be 'ok'): %v", string(reqBodyNew))
		t.FailNow()
	}
	//========================================
	//Validate UserInfo is correct
	req2, _ := http.NewRequest("GET", "http://"+frontendApiAddress+"/userInfo", nil)
	// set session token header for request
	req2.AddCookie(&http.Cookie{Name: "Authorization", Value: UID})
	client2 := &http.Client{}
	resp2, err := client2.Do(req2)
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.FailNow()
	}
	//unmarshal response into struct
	reqBody2, _ := ioutil.ReadAll(resp2.Body)
	var reqBackend2, expectedBackend2 structs.Client
	json.Unmarshal(reqBody2, &reqBackend2)
	//check that UID is missing from response
	if reqBackend2.UID != "" {
		t.Errorf("UID Should not be present: UID=%+v", reqBackend2.UID)
		t.FailNow()
	}
	//compare received struct with expected struct
	EXPECTED := []byte(`{"FirstName":"Anthony7","LastName":"Hanna","Weight":555,"WaistCirc":55.5,"HeightInches":55,"LeanBodyMass":55,"Age":55,"Gender":"male","Week":[{"Day":[{},{},{},{},{},{"Fat":500,"Carbs":500,"Protein":500,"TotalCalories":300,"DayCalories":"normal","Weight":321,"Cardio":"missed","WeightTraining":"no"},{}]},{"Day":[{},{"Fat":555,"Carbs":555,"Protein":555,"TotalCalories":555,"DayCalories":"normal","Weight":555,"Cardio":"missed","WeightTraining":"no"},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{"Fat":5050,"Carbs":5050,"Protein":5050,"TotalCalories":5050,"DayCalories":"normal","Weight":10,"Cardio":"missed","WeightTraining":"yes"},{"Fat":777,"Carbs":777,"Protein":777,"TotalCalories":777,"DayCalories":"normal","Weight":777,"Cardio":"missed","WeightTraining":"no"}]}],"Recommendation":[{"HighDayProtein":50,"HighDayCarb":11,"HighDayFat":12,"HighDayCalories":13,"NormalDayProtein":14,"NormalDayCarb":15,"NormalDayFat":16,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":24,"HIITChangeCardioIntervals":25,"ModifiedDate":"2020-09-15"},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}]}`)
	json.Unmarshal(EXPECTED, &expectedBackend2)
	if !cmp.Equal(reqBackend2, expectedBackend2) {
		t.Errorf("Database returned unexpected body: \ngot -: %+v \nwant -: %+v", string(reqBody2), string(EXPECTED))
		t.FailNow()
	}
}

//FIXME revise this when BACKEND/generateUserBaseline is returning correct results
// //TestGenerateUserBaseline validates updating daily data immediately following previous day
// func TestGenerateUserBaseline(t *testing.T) {
// 	var UID = "front-api-testUID8"
// 	//post update data
// 	var jsonStrNew = []byte(`{"FirstName":"Anthony8","LastName":"Hanna","Weight":250,"WaistCirc":50.5,"HeightInches":120,"LeanBodyMass":100,"Age":26,"Gender":"male"}`)
// 	reqNew, _ := http.NewRequest("POST", "http://" + frontendApiAddress + "/generateUserBaseline", bytes.NewBuffer(jsonStrNew))
// 	// set session token header for request
// 	reqNew.AddCookie(&http.Cookie{Name: "Authorization", Value: UID})
// 	reqNew.Header.Set("Content-Type", "application/json")
// 	clientNew := &http.Client{}
// 	respNew, err := clientNew.Do(reqNew)
// 	if err != nil {
// 		t.Errorf("Connection failed: %v", err)
// 		t.FailNow()
// 	}
// 	reqBodyNew, _ := ioutil.ReadAll(respNew.Body)
// 	//check that post response is "ok"
// 	if string(reqBodyNew) != "ok" {
// 		t.Errorf("Post Request Response Incorrect (should be 'ok'): %v", string(reqBodyNew))
// 		t.FailNow()
// 	}
// 	//========================================
// 	//Validate UserInfo is correct
// 	req2, _ := http.NewRequest("GET", "http://" + frontendApiAddress + "/userInfo", nil)
// 	// set session token header for request
// 	req2.AddCookie(&http.Cookie{Name: "Authorization", Value: UID})
// 	client2 := &http.Client{}
// 	resp2, err := client2.Do(req2)
// 	if err != nil {
// 		t.Errorf("Connection failed: %v", err)
// 		t.FailNow()
// 	}
// 	//unmarshal response into struct
// 	reqBody2, _ := ioutil.ReadAll(resp2.Body)
// 	var reqBackend2, expectedBackend2 structs.Client
// 	json.Unmarshal(reqBody2, &reqBackend2)
// 	//check that UID is missing from response
// 	if reqBackend2.UID != "" {
// 		t.Errorf("UID Should not be present: UID=%+v", reqBackend2.UID)
// 		t.FailNow()
// 	}
// 	//compare received struct with expected struct
// 	EXPECTED := []byte(`{"FirstName":"Anthony8","LastName":"Hanna","Weight":250,"WaistCirc":50.5,"HeightInches":120,"LeanBodyMass":100,"Age":26,"Gender":"male","Week":[{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]}],"Recommendation":[{"HighDayProtein":233,"HighDayCarb":584,"HighDayFat":156,"NormalDayProtein":393,"NormalDayCarb":383,"NormalDayFat":115,"LowDayProtein":244,"LowDayCarb":289,"LowDayFat":165},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}]}`)
// 	json.Unmarshal(EXPECTED, &expectedBackend2)
// 	if !cmp.Equal(reqBackend2, expectedBackend2) {
// 		t.Errorf("Database returned unexpected body: \ngot -: %+v \nwant -: %+v", string(reqBody2), string(EXPECTED))
// 		t.FailNow()
// 	}
// }

//FIXME revise this when BACKEND/generateUserBaseline is returning correct results
// //TestGenerateUserBaselineOnExistingData generates baselinedata over existing user data
// func TestGenerateUserBaselineOnExistingData(t *testing.T) {
// 	var UID = "front-api-testUID9"
// 	//post prop data for testing against
// 	var jsonStr = []byte(`{"FirstName":"Anthony9","LastName":"Hanna","Weight":555,"WaistCirc":55.5,"HeightInches":55,"LeanBodyMass":55,"Age":55,"Gender":"male","Week":[{"Day":[{},{},{},{},{},{"Fat":500,"Carbs":500,"Protein":500,"TotalCalories":300,"DayCalories":"normal","Weight":321,"Cardio":"missed","WeightTraining":"no"},{}]},{"Day":[{},{"Fat":555,"Carbs":555,"Protein":555,"TotalCalories":555,"DayCalories":"normal","Weight":555,"Cardio":"missed","WeightTraining":"no"},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{"Fat":5050,"Carbs":5050,"Protein":5050,"TotalCalories":5050,"DayCalories":"normal","Weight":10,"Cardio":"missed","WeightTraining":"yes"},{}]}],"Recommendation":[{"HighDayProtein":1111,"HighDayCarb":1111,"HighDayFat":1211,"HighDayCalories":1311,"NormalDayProtein":1411,"NormalDayCarb":1511,"NormalDayFat":1611,"NormalDayCalories":1711,"LowDayProtein":1811,"LowDayCarb":1911,"LowDayFat":2011,"LowDayCalories":2111,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":24,"HIITChangeCardioIntervals":25,"ModifiedDate":"2020-09-111"},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{"HighDayProtein":99,"HighDayCarb":99,"HighDayFat":99,"HighDayCalories":99,"NormalDayProtein":14,"NormalDayCarb":15,"NormalDayFat":16,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":99,"HIITChangeCardioIntervals":99,"ModifiedDate":"2020-09-99"},{},{},{}]}`)
// 	req, _ := http.NewRequest("POST", "http://" + frontendApiAddress + "/userInfo", bytes.NewBuffer(jsonStr))
// 	// set session token header for request
// 	req.AddCookie(&http.Cookie{Name: "Authorization", Value: UID})
// 	req.Header.Set("Content-Type", "application/json")
// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		t.Errorf("Connection failed: %v", err)
// 		t.FailNow()
// 	}
// 	reqBody, _ := ioutil.ReadAll(resp.Body)
// 	//check that post response is "ok"
// 	if string(reqBody) != "ok" {
// 		t.Errorf("Post Request Response Incorrect (should be 'ok'): %v", string(reqBody))
// 		t.FailNow()
// 	}
// 	//========================================
// 	//post update data
// 	var jsonStrNew = []byte(`{"FirstName":"Anthony99","LastName":"Hanna","Weight":251,"WaistCirc":50.5,"HeightInches":120,"LeanBodyMass":100,"Age":26,"Gender":"male"}`)
// 	reqNew, _ := http.NewRequest("POST", "http://" + frontendApiAddress + "/generateUserBaseline", bytes.NewBuffer(jsonStrNew))
// 	// set session token header for request
// 	reqNew.AddCookie(&http.Cookie{Name: "Authorization", Value: UID})
// 	reqNew.Header.Set("Content-Type", "application/json")
// 	clientNew := &http.Client{}
// 	respNew, err := clientNew.Do(reqNew)
// 	if err != nil {
// 		t.Errorf("Connection failed: %v", err)
// 		t.FailNow()
// 	}
// 	reqBodyNew, _ := ioutil.ReadAll(respNew.Body)
// 	//check that post response is "ok"
// 	if string(reqBodyNew) != "ok" {
// 		t.Errorf("Post Request Response Incorrect (should be 'ok'): %v", string(reqBodyNew))
// 		t.FailNow()
// 	}
// 	//========================================
// 	//Validate UserInfo is correct
// 	req2, _ := http.NewRequest("GET", "http://" + frontendApiAddress + "/userInfo", nil)
// 	// set session token header for request
// 	req2.AddCookie(&http.Cookie{Name: "Authorization", Value: UID})
// 	client2 := &http.Client{}
// 	resp2, err := client2.Do(req2)
// 	if err != nil {
// 		t.Errorf("Connection failed: %v", err)
// 		t.FailNow()
// 	}
// 	//unmarshal response into struct
// 	reqBody2, _ := ioutil.ReadAll(resp2.Body)
// 	var reqBackend2, expectedBackend2 structs.Client
// 	json.Unmarshal(reqBody2, &reqBackend2)
// 	//check that UID is missing from response
// 	if reqBackend2.UID != "" {
// 		t.Errorf("UID Should not be present: UID=%+v", reqBackend2.UID)
// 		t.FailNow()
// 	}
// 	//compare received struct with expected struct
// 	EXPECTED := []byte(`{"FirstName":"Anthony99","LastName":"Hanna","Weight":251,"WaistCirc":50.5,"HeightInches":120,"LeanBodyMass":100,"Age":26,"Gender":"male","Week":[{"Day":[{},{},{},{},{},{"Fat":500,"Carbs":500,"Protein":500,"TotalCalories":300,"DayCalories":"normal","Weight":321,"Cardio":"missed","WeightTraining":"no"},{}]},{"Day":[{},{"Fat":555,"Carbs":555,"Protein":555,"TotalCalories":555,"DayCalories":"normal","Weight":555,"Cardio":"missed","WeightTraining":"no"},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{"Fat":5050,"Carbs":5050,"Protein":5050,"TotalCalories":5050,"DayCalories":"normal","Weight":10,"Cardio":"missed","WeightTraining":"yes"},{}]}],"Recommendation":[{"HighDayProtein":234,"HighDayCarb":585,"HighDayFat":156,"NormalDayProtein":394,"NormalDayCarb":384,"NormalDayFat":115,"LowDayProtein":244,"LowDayCarb":290,"LowDayFat":165},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}]}`)
// 	json.Unmarshal(EXPECTED, &expectedBackend2)
// 	if !cmp.Equal(reqBackend2, expectedBackend2) {
// 		t.Errorf("Database returned unexpected body: \ngot -: %+v \nwant -: %+v", string(reqBody2), string(EXPECTED))
// 		t.FailNow()
// 	}
// }

//TODO perform data validation on Week and Modified date/ensure they are handled server side on frontend-api
//TestUpdateUserRecommendations add user recommendation to profile
// func TestUpdateUserRecommendations(t *testing.T) {
// 	var UID = "front-api-testUID10"
// 	//post prop data for testing against
// 	var jsonStr = []byte(`{"FirstName":"Anthony10","LastName":"Hanna","Weight":100,"WaistCirc":55.5,"HeightInches":55,"LeanBodyMass":55,"Age":55,"Gender":"male","Week":[{"Day":[{},{},{},{},{},{"Fat":1000,"Carbs":1000,"Protein":100,"TotalCalories":300,"DayCalories":"normal","Weight":321,"Cardio":"missed","WeightTraining":"no"},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{"Fat":1010,"Carbs":1010,"Protein":1010,"TotalCalories":1010,"DayCalories":"normal","Weight":10,"Cardio":"missed","WeightTraining":"yes"},{}]}],"Recommendation":[{"HighDayProtein":50,"HighDayCarb":11,"HighDayFat":12,"HighDayCalories":13,"NormalDayProtein":14,"NormalDayCarb":15,"NormalDayFat":16,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":24,"HIITChangeCardioIntervals":25,"ModifiedDate":"2020-09-15"},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}]}`)
// 	req, _ := http.NewRequest("POST", "http://" + frontendApiAddress + "/userInfo", bytes.NewBuffer(jsonStr))
// 	// set session token header for request
// 	req.AddCookie(&http.Cookie{Name: "Authorization", Value: UID})
// 	req.Header.Set("Content-Type", "application/json")
// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		t.Errorf("Connection failed: %v", err)
// 		t.FailNow()
// 	}
// 	reqBody, _ := ioutil.ReadAll(resp.Body)
// 	//check that post response is "ok"
// 	if string(reqBody) != "ok" {
// 		t.Errorf("Post Request Response Incorrect (should be 'ok'): %v", string(reqBody))
// 		t.FailNow()
// 	}
// 	//========================================
// 	//post update data
// 	var jsonStrNew = []byte(`{"HighDayProtein":1009,"HighDayCarb":100,"HighDayFat":100,"HighDayCalories":100,"NormalDayProtein":1004,"NormalDayCarb":1005,"NormalDayFat":1006,"NormalDayCalories":1007,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":1001,"HIITChangeCardioIntervals":25,"ModifiedDate":"2020-09-1009"}`)
// 	reqNew, _ := http.NewRequest("POST", "http://" + frontendApiAddress + "/userRecommendations/0", bytes.NewBuffer(jsonStrNew))
// 	// set session token header for request
// 	reqNew.AddCookie(&http.Cookie{Name: "Authorization", Value: UID})
// 	reqNew.Header.Set("Content-Type", "application/json")
// 	clientNew := &http.Client{}
// 	respNew, err := clientNew.Do(reqNew)
// 	if err != nil {
// 		t.Errorf("Connection failed: %v", err)
// 		t.FailNow()
// 	}
// 	reqBodyNew, _ := ioutil.ReadAll(respNew.Body)
// 	//check that post response is "ok"
// 	if string(reqBodyNew) != "ok" {
// 		t.Errorf("Post Request Response Incorrect (should be 'ok'): %v", string(reqBodyNew))
// 		t.FailNow()
// 	}
// 	//========================================
// 	//Validate UserInfo is correct
// 	req2, _ := http.NewRequest("GET", "http://" + frontendApiAddress + "/userInfo", nil)
// 	// set session token header for request
// 	req2.AddCookie(&http.Cookie{Name: "Authorization", Value: UID})
// 	client2 := &http.Client{}
// 	resp2, err := client2.Do(req2)
// 	if err != nil {
// 		t.Errorf("Connection failed: %v", err)
// 		t.FailNow()
// 	}
// 	//unmarshal response into struct
// 	reqBody2, _ := ioutil.ReadAll(resp2.Body)
// 	var reqBackend2, expectedBackend2 structs.Client
// 	json.Unmarshal(reqBody2, &reqBackend2)
// 	//check that UID is missing from response
// 	if reqBackend2.UID != "" {
// 		t.Errorf("UID Should not be present: UID=%+v", reqBackend2.UID)
// 		t.FailNow()
// 	}
// 	//compare received struct with expected struct
// 	EXPECTED := []byte(`{"FirstName":"Anthony10","LastName":"Hanna","Weight":100,"WaistCirc":55.5,"HeightInches":55,"LeanBodyMass":55,"Age":55,"Gender":"male","Week":[{"Day":[{},{},{},{},{},{"Fat":1000,"Carbs":1000,"Protein":100,"TotalCalories":300,"DayCalories":"normal","Weight":321,"Cardio":"missed","WeightTraining":"no"},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{"Fat":1010,"Carbs":1010,"Protein":1010,"TotalCalories":1010,"DayCalories":"normal","Weight":10,"Cardio":"missed","WeightTraining":"yes"},{}]}],"Recommendation":[{"HighDayProtein":1009,"HighDayCarb":100,"HighDayFat":100,"HighDayCalories":100,"NormalDayProtein":1004,"NormalDayCarb":1005,"NormalDayFat":1006,"NormalDayCalories":1007,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":1001,"HIITChangeCardioIntervals":25,"ModifiedDate":"2020-09-1009"},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}]}`)
// 	json.Unmarshal(EXPECTED, &expectedBackend2)
// 	if !cmp.Equal(reqBackend2, expectedBackend2) {
// 		t.Errorf("Database returned unexpected body: \ngot -: %+v \nwant -: %+v", string(reqBody2), string(EXPECTED))
// 		t.FailNow()
// 	}
// }

//TODO perform data validation on Week and Modified date/ensure they are handled server side on frontend-api
//TestAddUserRecommendations add second user recommendation to profile
// func TestAddUserRecommendations(t *testing.T) {
// 	var UID = "front-api-testUID1a"
// 	//post prop data for testing against
// 	var jsonStr = []byte(`{"FirstName":"Anthony11","LastName":"Hanna","Weight":100,"WaistCirc":55.5,"HeightInches":55,"LeanBodyMass":55,"Age":55,"Gender":"male","Week":[{"Day":[{},{},{},{},{},{"Fat":1000,"Carbs":1000,"Protein":100,"TotalCalories":300,"DayCalories":"normal","Weight":321,"Cardio":"missed","WeightTraining":"no"},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{"Fat":1010,"Carbs":1010,"Protein":1010,"TotalCalories":1010,"DayCalories":"normal","Weight":10,"Cardio":"missed","WeightTraining":"yes"},{}]}],"Recommendation":[{"HighDayProtein":1009,"HighDayCarb":100,"HighDayFat":100,"HighDayCalories":100,"NormalDayProtein":1004,"NormalDayCarb":1005,"NormalDayFat":1006,"NormalDayCalories":1007,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":1001,"HIITChangeCardioIntervals":25,"Week":1000,"ModifiedDate":"2020-09-1009"},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}]}`)
// 	req, _ := http.NewRequest("POST", "http://" + frontendApiAddress + "/userInfo", bytes.NewBuffer(jsonStr))
// 	// set session token header for request
// 	req.AddCookie(&http.Cookie{Name: "Authorization", Value: UID})
// 	req.Header.Set("Content-Type", "application/json")
// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		t.Errorf("Connection failed: %v", err)
// 		t.FailNow()
// 	}
// 	reqBody, _ := ioutil.ReadAll(resp.Body)
// 	//check that post response is "ok"
// 	if string(reqBody) != "ok" {
// 		t.Errorf("Post Request Response Incorrect (should be 'ok'): %v", string(reqBody))
// 		t.FailNow()
// 	}
// 	//========================================
// 	//post update data
// 	var jsonStrNew = []byte(`{"HighDayProtein":11,"HighDayCarb":11,"HighDayFat":11,"HighDayCalories":11,"NormalDayProtein":11,"NormalDayCarb":1005,"NormalDayFat":1006,"NormalDayCalories":1007,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":1001,"HIITChangeCardioIntervals":25}`)
// 	reqNew, _ := http.NewRequest("POST", "http://" + frontendApiAddress + "/userRecommendations/2", bytes.NewBuffer(jsonStrNew))
// 	// set session token header for request
// 	reqNew.AddCookie(&http.Cookie{Name: "Authorization", Value: UID})
// 	reqNew.Header.Set("Content-Type", "application/json")
// 	clientNew := &http.Client{}
// 	respNew, err := clientNew.Do(reqNew)
// 	if err != nil {
// 		t.Errorf("Connection failed: %v", err)
// 		t.FailNow()
// 	}
// 	reqBodyNew, _ := ioutil.ReadAll(respNew.Body)
// 	//check that post response is "ok"
// 	if string(reqBodyNew) != "ok" {
// 		t.Errorf("Post Request Response Incorrect (should be 'ok'): %v", string(reqBodyNew))
// 		t.FailNow()
// 	}
// 	//========================================
// 	//Validate UserInfo is correct
// 	req2, _ := http.NewRequest("GET", "http://" + frontendApiAddress + "/userInfo", nil)
// 	// set session token header for request
// 	req2.AddCookie(&http.Cookie{Name: "Authorization", Value: UID})
// 	client2 := &http.Client{}
// 	resp2, err := client2.Do(req2)
// 	if err != nil {
// 		t.Errorf("Connection failed: %v", err)
// 		t.FailNow()
// 	}
// 	//unmarshal response into struct
// 	reqBody2, _ := ioutil.ReadAll(resp2.Body)
// 	var reqBackend2, expectedBackend2 structs.Client
// 	json.Unmarshal(reqBody2, &reqBackend2)
// 	//check that UID is missing from response
// 	if reqBackend2.UID != "" {
// 		t.Errorf("UID Should not be present: UID=%+v", reqBackend2.UID)
// 		t.FailNow()
// 	}
// 	//compare received struct with expected struct
// 	EXPECTED := []byte(`{"FirstName":"Anthony11","LastName":"Hanna","Weight":100,"WaistCirc":55.5,"HeightInches":55,"LeanBodyMass":55,"Age":55,"Gender":"male","Week":[{"Day":[{},{},{},{},{},{"Fat":1000,"Carbs":1000,"Protein":100,"TotalCalories":300,"DayCalories":"normal","Weight":321,"Cardio":"missed","WeightTraining":"no"},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{"Fat":1010,"Carbs":1010,"Protein":1010,"TotalCalories":1010,"DayCalories":"normal","Weight":10,"Cardio":"missed","WeightTraining":"yes"},{}]}],"Recommendation":[{"HighDayProtein":1009,"HighDayCarb":100,"HighDayFat":100,"HighDayCalories":100,"NormalDayProtein":1004,"NormalDayCarb":1005,"NormalDayFat":1006,"NormalDayCalories":1007,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":1001,"HIITChangeCardioIntervals":25,"Week":1000,"ModifiedDate":"2020-09-1009"},{},{"HighDayProtein":11,"HighDayCarb":11,"HighDayFat":11,"HighDayCalories":11,"NormalDayProtein":11,"NormalDayCarb":1005,"NormalDayFat":1006,"NormalDayCalories":1007,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":1001,"HIITChangeCardioIntervals":25,"Week":1111,"ModifiedDate":"2020-11-11"},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}]}`)
// 	json.Unmarshal(EXPECTED, &expectedBackend2)
// 	loc, _ := time.LoadLocation("America/Los_Angeles")
// 	expectedBackend2.Recommendation[2].ModifiedDate = time.Now().In(loc).Format("2006-01-02")
// 	if !cmp.Equal(reqBackend2, expectedBackend2) {
// 		t.Errorf("Database returned unexpected body: \ngot -: %+v \nwant -: %+v", string(reqBody2), string(EXPECTED))
// 		t.FailNow()
// 	}
// }

func TestAdminGetUserInfo(t *testing.T) {
	var UID = "test3"
	var userUID = "testUID"
	req, _ := http.NewRequest("GET", "http://"+frontendApiAddress+"/admin/userInfo", nil)
	// set session token header for request
	req.AddCookie(&http.Cookie{Name: "Authorization", Value: UID})
	req.Header.Set("User-UID", userUID)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.FailNow()
	}
	//unmarshal response into struct
	reqBody, _ := ioutil.ReadAll(resp.Body)
	var reqBackend, expectedBackend structs.Client
	json.Unmarshal(reqBody, &reqBackend)
	//check that UID is missing from response
	if reqBackend.UID != "" {
		t.Errorf("UID Should not be present: UID=%+v", reqBackend.UID)
		t.FailNow()
	}
	//compare received struct with expected struct
	EXPECTED := []byte(`{"FirstName":"Anthony","LastName":"Hanna","Weight":215,"WaistCirc":35,"HeightInches":75,"LeanBodyMass":15,"Age":20,"StartDate":"2020-08-15","Gender":"male","Week":[{"Day":[{"WaistCirc":20,"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalories":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":40,"Carbs":40,"Protein":40,"TotalCalories":34,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":100,"Carbs":100,"Protein":100,"TotalCalories":300,"DayCalories":"normal","Weight":321,"Cardio":"missed","WeightTraining":"no"}]},{"Day":[{"Fat":11,"Carbs":11,"Protein":11,"TotalCalories":31,"DayCalories":"normal","Weight":222,"Cardio":"missed","WeightTraining":"yes"},{"WaistCirc":20,"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{},{},{},{},{}]},{"Day":[{"Fat":110,"Carbs":110,"Protein":110,"TotalCalories":310,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]}],"Recommendation":[{"HighDayProtein":10,"HighDayCarb":11,"HighDayFat":12,"HighDayCalories":13,"NormalDayProtein":14,"NormalDayCarb":15,"NormalDayFat":16,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITCurrentCardioIntervals":24},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}]}`)
	json.Unmarshal(EXPECTED, &expectedBackend)
	if !cmp.Equal(reqBackend, expectedBackend) {
		t.Errorf("frontend-api returned unexpected body: \ngot -: %+v \nwant -: %+v", string(reqBody), string(EXPECTED))
		t.FailNow()
	}
}

func TestInvalidAdminGetUserInfo(t *testing.T) {
	var UID = "test2"
	var userUID = "testUID"
	req, _ := http.NewRequest("GET", "http://"+frontendApiAddress+"/admin/userInfo", nil)
	// set session token header for request
	req.AddCookie(&http.Cookie{Name: "Authorization", Value: UID})
	req.Header.Set("User-UID", userUID)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.FailNow()
	}
	//compare received struct with expected struct
	if resp.StatusCode != 403 {
		t.Errorf("Status Incorrect, should be 403 -: http.StatusCode=%v", resp.StatusCode)
		t.FailNow()
	}
}

//FIXME write tests for /admin/userRecommendation/{week}
