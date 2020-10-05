package main

import (
	// "bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	// "strings"
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
	os.Exit(m.Run())
}

func TestAPIVersion(t *testing.T) {
	req, err := http.Get("http://" + frontendApiAddress + "/apiVersion")
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.Fail()
	}
	// Check the response body is what we expect.
	expected := `{"apiVersion":1.0}`
	respBody, _ := ioutil.ReadAll(req.Body)
	if string(respBody) != expected {
		t.Errorf("Frontend-api returned unexpected body: got %v \nwant %v", string(respBody), expected)
		t.Fail()
	}
}

//TestGetUserInfo get user profile and verify UID is stripped
func TestGetUserInfo(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://" + frontendApiAddress + "/userInfo", nil)
	// set session token header for request
	req.Header.Set("Session-Token", "testUID")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Connection failed: %v", err)
		t.Fail()
	}
	//unmarshal response into struct
	reqBody, _ := ioutil.ReadAll(resp.Body)
	var reqBackend, expectedBackend structs.Client
	json.Unmarshal(reqBody, &reqBackend)
	//check that UID is missing from response
	if reqBackend.UID != "" {
		t.Errorf("UID Should not be present: UID=%+v", reqBackend.UID)
		t.Fail()
	}
	//compare received struct with expected struct
	EXPECTED := []byte(`{"FirstName":"Anthony","LastName":"Hanna","Weight":215,"WaistCirc":35.5,"HeightInches":75,"LeanBodyMass":15,"Age":20,"Gender":"male","Week":[{"Day":[{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":20,"Carbs":20,"Protein":20,"TotalCalories":32,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{"Fat":30,"Carbs":30,"Protein":30,"TotalCalories":33,"DayCalories":"high","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":40,"Carbs":40,"Protein":40,"TotalCalories":34,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":100,"Carbs":100,"Protein":100,"TotalCalories":300,"DayCalories":"normal","Weight":321,"Cardio":"missed","WeightTraining":"no"}]},{"Day":[{"Fat":11,"Carbs":11,"Protein":11,"TotalCalories":31,"DayCalories":"normal","Weight":222,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"no"},{},{},{},{},{}]},{"Day":[{"Fat":110,"Carbs":110,"Protein":110,"TotalCalories":310,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{"Fat":10,"Carbs":10,"Protein":10,"TotalCalories":30,"DayCalories":"normal","Weight":123,"Cardio":"missed","WeightTraining":"yes"},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]},{"Day":[{},{},{},{},{},{},{}]}],"Recommendation":[{"HighDayProtein":10,"HighDayCarb":11,"HighDayFat":12,"HighDayCalories":13,"NormalDayProtein":14,"NormalDayCarb":15,"NormalDayFat":16,"NormalDayCalories":17,"LowDayProtein":18,"LowDayCarb":19,"LowDayFat":20,"LowDayCalories":21,"HIITCurrentCardioSession":22,"HIITChangeCardioSession":23,"HIITCurrentCardioIntervals":24,"HIITChangeCardioIntervals":25,"Week":1,"ModifiedDate":"2020-09-13"},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}]}`)
	json.Unmarshal(EXPECTED, &expectedBackend)
	if !cmp.Equal(reqBackend, expectedBackend) {
		t.Errorf("Database returned unexpected body: \ngot -: %+v \nwant -: %+v", string(reqBody), string(EXPECTED))
		t.Fail()
	}
}

// func TestGetUserProfile(t *testing.T) {
// 	req, err := http.NewRequest("GET", "http://localhost/userProfile", nil)
// 	// set session token header for request
// 	req.Header.Set("Session-Token", "test")
// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		t.Errorf("Connection failed: %v", err)
// 		t.Fail()
// 	}
// 	// Check the response body is what we expect.
// 	expected := `{"FirstName":"Anthony","LastName":"Hannah","Weight":215,"WaistCirc":11,"HeightInches":72,"LeanBodyMass":15,"Age":27,"Gender":"female"}`
// 	respBody, _ := ioutil.ReadAll(resp.Body)
// 	if string(respBody) != expected {
// 		t.Errorf("Frontend-api returned unexpected body: got %v \nwant %v", string(respBody), expected)
// 		t.Fail()
// 	}
// }

// func TestGetUserBaseline(t *testing.T) {
// 	req, err := http.NewRequest("GET", "http://localhost/userBaseline", nil)
// 	// set session token header for request
// 	req.Header.Set("Session-Token", "test")
// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		t.Errorf("Connection failed: %v", err)
// 		t.Fail()
// 	}
// 	// Check the response body is what we expect.
// 	expected := `{"LowDay":2599,"NormalDay":2978,"HighDay":3357,"NFatRatio":0.25,"NCarbRatio":0.37,"NProteinRatio":0.38,"HFatRatio":0.3,"HCarbRatio":0.5,"HProteinRatio":0.2,"LFatRatio":0.41,"LCarbRatio":0.32,"LProteinRatio":0.27,"NFatAmount":83,"NCarbAmount":275,"NProteinAmount":283,"HFatAmount":112,"HCarbAmount":420,"HProteinAmount":168,"LFatAmount":118,"LCarbAmount":208,"LProteinAmount":175}`
// 	respBody, _ := ioutil.ReadAll(resp.Body)
// 	if string(respBody) != expected {
// 		t.Errorf("Frontend-api returned unexpected body: got %v \nwant %v", string(respBody), expected)
// 		t.Fail()
// 	}
// }
