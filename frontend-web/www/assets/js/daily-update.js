var myToken = "test2"; //FIXME remove

//get user info from server
var xmlHttp = new XMLHttpRequest();
xmlHttp.open("GET", "http://localhost:8888/userInfo", false);
// xmlHttp.setRequestHeader("Session-Token",getCookie("Session-Token"));
xmlHttp.setRequestHeader("Session-Token", myToken); //FIXME use correct session-token
xmlHttp.send(null);
var userData = JSON.parse(xmlHttp.responseText);

//FIXME remove
// console.log(JSON.stringify(userData.Recommendation));
// console.log(JSON.stringify(getLatestRecommendation(userData)));
//FIXME remove

updateCharts(userData);

//get UserBaseline json object, call other functions on complete
// fetch("http://localhost:8080/getUserBaseline", {
//   headers: {
//     'Session-Token': getCookie("Session-Token"),
//   }
// }).then(function (response) {
//   response.json().then(function (data) {
//     updateCharts(data);
//   });
// });


// updates baseline & recommendation charts with user data
function updateCharts(userData) {
  // redirect to profile page on empty data
  if (userData.Recommendation[0].NormalDayCalories == 0) {
    console.log("Baseline data blank, redirecting to profile...")
    // window.location.replace('http://localhost:5500/profile');  //FIXME
  }
  // === COACH RECOMMENDATION CHART ===
  //get latest recommendation object for charts
  var latestRec = getLatestRecommendation(userData)
  // only show if a recommendation exists
  if (latestRec == null) {
    document.getElementById("CoachRecContainer").remove();
    document.getElementById("CoachRecHr").remove();
  } else {
    document.getElementById("updateDateText").innerHTML = "Last Updated: " + latestRec.ModifiedDate;
    //Normal Day
    document.getElementById("coach-NProteinAmount").innerHTML = latestRec.NormalDayProtein;
    // document.getElementById("coach-NProteinRatio").innerHTML = latestRec.NProteinRatio;
    document.getElementById("coach-NCarbAmount").innerHTML = latestRec.NormalDayCarb;
    // document.getElementById("coach-NCarbRatio").innerHTML = latestRec.NCarbRatio;
    document.getElementById("coach-NFatAmount").innerHTML = latestRec.NormalDayFat;
    // document.getElementById("coach-NFatRatio").innerHTML = latestRec.NFatRatio;
    document.getElementById("coach-NCalories").innerHTML = latestRec.NormalDayCalories;
    document.getElementById("coach-NCaloriesRatio").innerHTML = (latestRec.NormalDayCalories/latestRec.NormalDayFat).toPrecision(1) + "%";
    //High Day
    document.getElementById("coach-HProteinAmount").innerHTML = latestRec.HighDayProtein;
    // document.getElementById("coach-HProteinRatio").innerHTML = latestRec.HProteinRatio;
    document.getElementById("coach-HCarbAmount").innerHTML = latestRec.HighDayCarb;
    // document.getElementById("coach-HCarbRatio").innerHTML = latestRec.HCarbRatio;
    document.getElementById("coach-HFatAmount").innerHTML = latestRec.HighDayFat;
    // document.getElementById("coach-HFatRatio").innerHTML = latestRec.HFatRatio;
    document.getElementById("coach-HCalories").innerHTML = latestRec.HighDayCalories;
    document.getElementById("coach-HCaloriesRatio").innerHTML = (latestRec.HighDayCalories/latestRec.HighDayFat).toPrecision(1) + "%";
    //Low Day
    document.getElementById("coach-LProteinAmount").innerHTML = latestRec.LowDayProtein;
    // document.getElementById("coach-LProteinRatio").innerHTML = latestRec.LProteinRatio;
    document.getElementById("coach-LCarbAmount").innerHTML = latestRec.LowDayCarb;
    // document.getElementById("coach-LCarbRatio").innerHTML = latestRec.LCarbRatio;
    document.getElementById("coach-LFatAmount").innerHTML = latestRec.LowDayFat;
    // document.getElementById("coach-LFatRatio").innerHTML = latestRec.LFatRatio;
    document.getElementById("coach-LCalories").innerHTML = latestRec.LowDayCalories;
    document.getElementById("coach-LCaloriesRatio").innerHTML = (latestRec.LowDayCalories/latestRec.LowDayFat).toPrecision(1) + "%";
  }

  // === BASELINE CHART === //TODO fix ratio references
  var baselineRec = userData.Recommendation[0]
  //Normal Day
  document.getElementById("NProteinAmount").innerHTML = baselineRec.NormalDayProtein;
  // document.getElementById("NProteinRatio").innerHTML = baselineRec.NProteinRatio;
  document.getElementById("NCarbAmount").innerHTML = baselineRec.NormalDayCarb;
  // document.getElementById("NCarbRatio").innerHTML = baselineRec.NCarbRatio;
  document.getElementById("NFatAmount").innerHTML = baselineRec.NormalDayFat;
  // document.getElementById("NFatRatio").innerHTML = baselineRec.NFatRatio;
  document.getElementById("NCalories").innerHTML = baselineRec.NormalDayCalories;
  document.getElementById("NCaloriesRatio").innerHTML = (baselineRec.NormalDayCalories/baselineRec.NormalDayFat).toPrecision(1) + "%";
  //High Day
  document.getElementById("HProteinAmount").innerHTML = baselineRec.HighDayProtein;
  // document.getElementById("HProteinRatio").innerHTML = baselineRec.HProteinRatio;
  document.getElementById("HCarbAmount").innerHTML = baselineRec.HighDayCarb;
  // document.getElementById("HCarbRatio").innerHTML = baselineRec.HCarbRatio;
  document.getElementById("HFatAmount").innerHTML = baselineRec.HighDayFat;
  // document.getElementById("HFatRatio").innerHTML = baselineRec.HFatRatio;
  document.getElementById("HCalories").innerHTML = baselineRec.HighDayCalories;
  document.getElementById("HCaloriesRatio").innerHTML = (baselineRec.HighDayCalories/baselineRec.HighDayFat).toPrecision(1) + "%";
  //Low Day
  document.getElementById("LProteinAmount").innerHTML = baselineRec.LowDayProtein;
  // document.getElementById("LProteinRatio").innerHTML = baselineRec.LProteinRatio;
  document.getElementById("LCarbAmount").innerHTML = baselineRec.LowDayCarb;
  // document.getElementById("LCarbRatio").innerHTML = baselineRec.LCarbRatio;
  document.getElementById("LFatAmount").innerHTML = baselineRec.LowDayFat;
  // document.getElementById("LFatRatio").innerHTML = baselineRec.LFatRatio;
  document.getElementById("LCalories").innerHTML = baselineRec.LowDayCalories;
  document.getElementById("LCaloriesRatio").innerHTML = (baselineRec.LowDayCalories/baselineRec.LowDayFat).toPrecision(1) + "%";
}

// returns latest recommendation object that is not the baseline (first obj in array)
function getLatestRecommendation(userData) {
  //make copy of user data to manipulate
  var tmpUserData = JSON.parse(JSON.stringify(userData.Recommendation));
  tmpUserData.shift();
  return tmpUserData.filter(value => Object.keys(value).length !== 0).slice(-1)[0];
}





// ============= YOINKED FROM PROFILE.js ==================
//FIXME

//serialize form fields into json object
function serializeProfile(form) {
  return {
    Fat: Number(document.getElementById("fat").value),
    Carbs: Number(document.getElementById("carbs").value),
    Protein: Number(document.getElementById("protein").value),
    Weight: Number(document.getElementById("weight").value),
    TotalCalories: Number(document.getElementById("totalCalories").value),
    DayCalories: document.getElementById("dayCalorie").value,
    Cardio: document.getElementById("cardio").value,
    WeightTraining: document.getElementById("weightTraining").value,
  };
}

//submit form as JSON on "save" button click
function submitForm() {
  //check that form is valid before sending
  if (document.getElementById("DailyUpdateForm").checkValidity() == false) {
    console.log("Form Invalid.")
    return
  }
  //serialize form to JSON
  var dataObject = serializeProfile(document.getElementById("DailyUpdateForm"));
  var jsonData = JSON.stringify(dataObject);
  console.log(jsonData);  //FIXME remove
  //TODO POST json to api
  //POST JSON to api
  var xmlHttp = new XMLHttpRequest();
  xmlHttp.open( "POST", "http://localhost:8888/userInfo", false );
  // xmlHttp.setRequestHeader("Session-Token",getCookie("Session-Token"));
  xmlHttp.setRequestHeader("Session-Token",myToken); //FIXME use correct session-token
  xmlHttp.send(jsonData);
  console.log("Server response: " + xmlHttp.responseText);
  //show note that save was successful
  document.getElementById("SaveConfirmationText").innerHTML = "&nbsp; &nbsp; &nbsp; Saved!";
}
// ============= YOINKED FROM PROFILE.js ==================
//FIXME

