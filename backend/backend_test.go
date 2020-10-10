package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	structs "../common"

	"github.com/google/go-cmp/cmp"
)

var backendAPIAddress string

// get service address from env
func TestMain(m *testing.M) {
	backendAPIAddress = os.Getenv("BACKEND_API_SERVICE_ADDRESS")
	if backendAPIAddress == "" {
		backendAPIAddress = "localhost:8090"
	}
	log.Println("Testing Backend-Api at address: " + backendAPIAddress)
	os.Exit(m.Run())
}

func TestAPIVersion(t *testing.T) {
	req, err := http.Get("http://" + backendAPIAddress + "/apiVersion")
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.FailNow()
	}
	// Check the response body is what we expect.
	expected := `{"apiVersion":1.0}`
	respBody, _ := ioutil.ReadAll(req.Body)
	if string(respBody) != expected {
		t.Errorf("Backend-api returned unexpected body: got %v \nwant %v", string(respBody), expected)
		t.FailNow()
	}
}

//TestGenerateUserBaseline test that first user recommendation is created
func TestGenerateUserBaseline(t *testing.T) {
	var UID = "backend-testUID1"
	//post data and validate that request succeeded
	var jsonStr = []byte(`{"FirstName":"Anthony1","LastName":"Hanna1","Weight":2151,"WaistCirc":35.51,"HeightInches":751,"LeanBodyMass":151,"Age":201,"Gender":"female"}`)
	req, _ := http.NewRequest("POST", "http://"+backendAPIAddress+"/generateUserBaseline/"+UID, bytes.NewBuffer(jsonStr))
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
	req2, _ := http.NewRequest("GET", "http://"+backendAPIAddress+"/userInfo/"+UID, nil)
	// set session token header for request
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
	//compare received struct with expected struct
	EXPECTED := []byte(`{"UID":"backend-testUID1","FirstName":"Anthony1","LastName":"Hanna1","Weight":2151,"WaistCirc":35.51,"HeightInches":751,"LeanBodyMass":151,"Age":201,"Gender":"female","Week":[{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]}],"Recommendation":[{"HighDayProtein":968,"HighDayCarb":2420,"HighDayFat":645,"NormalDayProtein":1631,"NormalDayCarb":1589,"NormalDayFat":477,"LowDayProtein":1012,"LowDayCarb":1199,"LowDayFat":683,"ModifiedDate":"2020-10-09"},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}]}`)
	json.Unmarshal(EXPECTED, &expectedBackend2)
	loc, _ := time.LoadLocation("America/Los_Angeles")
	expectedBackend2.Recommendation[0].ModifiedDate = time.Now().In(loc).Format("2006-01-02")
	if !cmp.Equal(reqBackend2, expectedBackend2) {
		t.Errorf("Database returned unexpected body: \ngot -: %+v \nwant -: %+v", string(reqBody2), string(EXPECTED))
		t.FailNow()
	}
}

//TestGenerateUserBaseline2 test that first user recommendation is created
func TestGenerateUserBaseline2(t *testing.T) {
	var UID = "backend-testUID2"
	//post data and validate that request succeeded
	var jsonStr = []byte(`{"FirstName":"Anthony2","LastName":"Hanna1","Weight":210,"WaistCirc":34,"HeightInches":75,"LeanBodyMass":20,"Age":22,"Gender":"male"}`)
	req, _ := http.NewRequest("POST", "http://"+backendAPIAddress+"/generateUserBaseline/"+UID, bytes.NewBuffer(jsonStr))
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
	req2, _ := http.NewRequest("GET", "http://"+backendAPIAddress+"/userInfo/"+UID, nil)
	// set session token header for request
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
	//compare received struct with expected struct
	EXPECTED := []byte(`{"UID":"backend-testUID2","FirstName":"Anthony2","LastName":"Hanna1","Weight":210,"WaistCirc":34,"HeightInches":75,"LeanBodyMass":20,"Age":22,"Gender":"male","Week":[{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]}],"Recommendation":[{"HighDayProtein":171,"HighDayCarb":428,"HighDayFat":114,"NormalDayProtein":288,"NormalDayCarb":281,"NormalDayFat":84,"LowDayProtein":179,"LowDayCarb":212,"LowDayFat":121,"ModifiedDate":"2020-10-09"},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}]}`)
	json.Unmarshal(EXPECTED, &expectedBackend2)
	loc, _ := time.LoadLocation("America/Los_Angeles")
	expectedBackend2.Recommendation[0].ModifiedDate = time.Now().In(loc).Format("2006-01-02")
	if !cmp.Equal(reqBackend2, expectedBackend2) {
		t.Errorf("Database returned unexpected body: \ngot -: %+v \nwant -: %+v", string(reqBody2), string(EXPECTED))
		t.FailNow()
	}
}

//TestGenerateUserBaseline3 test that first user recommendation is created odd data is passed
func TestGenerateUserBaseline3(t *testing.T) {
	var UID = "backend-testUID3"
	//post data and validate that request succeeded
	var jsonStr = []byte(`{"FirstName":"Anthony3","LastName":"Hanna1","Weight":123,"WaistCirc":20,"HeightInches":60,"LeanBodyMass":20,"Age":22,"Gender":"female"}`)
	req, _ := http.NewRequest("POST", "http://"+backendAPIAddress+"/generateUserBaseline/"+UID, bytes.NewBuffer(jsonStr))
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
	req2, _ := http.NewRequest("GET", "http://"+backendAPIAddress+"/userInfo/"+UID, nil)
	// set session token header for request
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
	//compare received struct with expected struct
	EXPECTED := []byte(`{"UID":"backend-testUID3","FirstName":"Anthony3","LastName":"Hanna1","Weight":123,"WaistCirc":20,"HeightInches":60,"LeanBodyMass":20,"Age":22,"Gender":"female","Week":[{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]}],"Recommendation":[{"HighDayProtein":106,"HighDayCarb":264,"HighDayFat":70,"NormalDayProtein":178,"NormalDayCarb":173,"NormalDayFat":52,"LowDayProtein":110,"LowDayCarb":131,"LowDayFat":74,"ModifiedDate":"2020-10-09"},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}]}`)
	json.Unmarshal(EXPECTED, &expectedBackend2)
	loc, _ := time.LoadLocation("America/Los_Angeles")
	expectedBackend2.Recommendation[0].ModifiedDate = time.Now().In(loc).Format("2006-01-02")
	if !cmp.Equal(reqBackend2, expectedBackend2) {
		t.Errorf("Database returned unexpected body: \ngot -: %+v \nwant -: %+v", string(reqBody2), string(EXPECTED))
		t.FailNow()
	}
}

//TestUpdateUserRecommendation add new user recommendation
func TestUpdateUserRecommendation(t *testing.T) {
	var UID = "back-api-testUID3"
	//post prop data for testing against
	var jsonStr = []byte(`{"UID":"back-api-testUID3","FirstName":"Test3","LastName":"Update7","Weight":777,"WaistCirc":777.2,"HeightInches":777,"LeanBodyMass":777,"Age":777,"Gender":"female","Week":[{"Day":[{"Fat":777,"Carbs":777,"Protein":777,"TotalCalories":777,"DayCalories":"high","Weight":777,"Cardio":"missed","WeightTraining":"yes"},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]}],"Recommendation":[{"HighDayProtein":10,"HighDayCarb":11,"HighDayFat":12,"HighDayCalories":13,"NormalDayProtein":14,"NormalDayCarb":15,"NormalDayFat":16,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":24,"HIITChangeCardioIntervals":25,"ModifiedDate":"2020-09-13"},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}]}`)
	req, _ := http.NewRequest("POST", "http://"+backendAPIAddress+"/userInfo/"+UID, bytes.NewBuffer(jsonStr))
	// set session token header for request
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
	//	post update data
	var jsonStrNew = []byte(`{"HighDayProtein":10,"HighDayCarb":11,"HighDayFat":12,"HighDayCalories":13,"NormalDayProtein":14,"NormalDayCarb":15,"NormalDayFat":16,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":24,"HIITChangeCardioIntervals":25}`)
	reqNew, _ := http.NewRequest("POST", "http://"+backendAPIAddress+"/userRecommendation/2/"+UID, bytes.NewBuffer(jsonStrNew))
	// set session token header for request

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
	//Validate UserInfo is correct
	req2, _ := http.NewRequest("GET", "http://"+backendAPIAddress+"/userInfo/"+UID, nil)
	// set session token header for request
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
	//compare received struct with expected struct
	EXPECTED := []byte(`{"UID":"back-api-testUID3","FirstName":"Test3","LastName":"Update7","Weight":777,"WaistCirc":777.2,"HeightInches":777,"LeanBodyMass":777,"Age":777,"Gender":"female","Week":[{"Day":[{"Fat":777,"Carbs":777,"Protein":777,"TotalCalories":777,"DayCalories":"high","Weight":777,"Cardio":"missed","WeightTraining":"yes"},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]}],"Recommendation":[{"HighDayProtein":10,"HighDayCarb":11,"HighDayFat":12,"HighDayCalories":13,"NormalDayProtein":14,"NormalDayCarb":15,"NormalDayFat":16,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":24,"HIITChangeCardioIntervals":25,"ModifiedDate":"2020-09-13"},{},{"HighDayProtein":10,"HighDayCarb":11,"HighDayFat":12,"HighDayCalories":13,"NormalDayProtein":14,"NormalDayCarb":15,"NormalDayFat":16,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":24,"HIITChangeCardioIntervals":25,"ModifiedDate":"2020-10-01"},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}]}`)
	json.Unmarshal(EXPECTED, &expectedBackend2)
	loc, _ := time.LoadLocation("America/Los_Angeles")
	expectedBackend2.Recommendation[2].ModifiedDate = time.Now().In(loc).Format("2006-01-02")
	if !cmp.Equal(reqBackend2, expectedBackend2) {
		t.Errorf("Database returned unexpected body: \ngot -: %+v \nwant -: %+v", string(reqBody2), string(EXPECTED))
		t.FailNow()
	}
}

//TestUpdateUserRecommendation2 test overwriting recommendation
func TestUpdateUserRecommendation2(t *testing.T) {
	var UID = "back-api-testUID4"
	//post prop data for testing against
	var jsonStr = []byte(`{"UID":"back-api-testUID4","FirstName":"Test4","LastName":"Update7","Weight":777,"WaistCirc":777.2,"HeightInches":777,"LeanBodyMass":777,"Age":777,"Gender":"female","Week":[{"Day":[{"Fat":777,"Carbs":777,"Protein":777,"TotalCalories":777,"DayCalories":"high","Weight":777,"Cardio":"missed","WeightTraining":"yes"},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]}],"Recommendation":[{"HighDayProtein":10,"HighDayCarb":11,"HighDayFat":12,"HighDayCalories":13,"NormalDayProtein":14,"NormalDayCarb":15,"NormalDayFat":16,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":24,"HIITChangeCardioIntervals":25,"ModifiedDate":"2019-09-13"},{},{"HighDayProtein":10,"HighDayCarb":11,"HighDayFat":12,"HighDayCalories":13,"NormalDayProtein":14,"NormalDayCarb":15,"NormalDayFat":16,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":24,"HIITChangeCardioIntervals":25,"ModifiedDate":"2019-10-44"},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}]}`)
	req, _ := http.NewRequest("POST", "http://"+backendAPIAddress+"/userInfo/"+UID, bytes.NewBuffer(jsonStr))
	// set session token header for request
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
	//	post update data
	var jsonStrNew = []byte(`{"HighDayProtein":44,"HighDayCarb":44,"HighDayFat":44,"HighDayCalories":44,"NormalDayProtein":44,"NormalDayCarb":44,"NormalDayFat":44,"NormalDayCalories":44,"LowDayProtein":44,"LowDayCarb":44,"LowDayFat":44,"LowDayCalories":44,"HIITCurrentCardioSession":44,"HIITChangeCardioSession":44,"HIITCurrentCardioIntervals":44,"HIITChangeCardioIntervals":44}`)
	reqNew, _ := http.NewRequest("POST", "http://"+backendAPIAddress+"/userRecommendation/0/"+UID, bytes.NewBuffer(jsonStrNew))
	// set session token header for request

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
	//Validate UserInfo is correct
	req2, _ := http.NewRequest("GET", "http://"+backendAPIAddress+"/userInfo/"+UID, nil)
	// set session token header for request
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
	//compare received struct with expected struct
	EXPECTED := []byte(`{"UID":"back-api-testUID4","FirstName":"Test4","LastName":"Update7","Weight":777,"WaistCirc":777.2,"HeightInches":777,"LeanBodyMass":777,"Age":777,"Gender":"female","Week":[{"Day":[{"Fat":777,"Carbs":777,"Protein":777,"TotalCalories":777,"DayCalories":"high","Weight":777,"Cardio":"missed","WeightTraining":"yes"},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]}],"Recommendation":[{"HighDayProtein":44,"HighDayCarb":44,"HighDayFat":44,"HighDayCalories":44,"NormalDayProtein":44,"NormalDayCarb":44,"NormalDayFat":44,"NormalDayCalories":44,"LowDayProtein":44,"LowDayCarb":44,"LowDayFat":44,"LowDayCalories":44,"HIITCurrentCardioSession":44,"HIITChangeCardioSession":44,"HIITCurrentCardioIntervals":44,"HIITChangeCardioIntervals":44,"ModifiedDate":"2020-10-09"},{},{"HighDayProtein":10,"HighDayCarb":11,"HighDayFat":12,"HighDayCalories":13,"NormalDayProtein":14,"NormalDayCarb":15,"NormalDayFat":16,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":24,"HIITChangeCardioIntervals":25,"ModifiedDate":"2019-10-44"},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}]}`)
	json.Unmarshal(EXPECTED, &expectedBackend2)
	loc, _ := time.LoadLocation("America/Los_Angeles")
	expectedBackend2.Recommendation[0].ModifiedDate = time.Now().In(loc).Format("2006-01-02")
	if !cmp.Equal(reqBackend2, expectedBackend2) {
		t.Errorf("Database returned unexpected body: \ngot -: %+v \nwant -: %+v", string(reqBody2), string(EXPECTED))
		t.FailNow()
	}
}