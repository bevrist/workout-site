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
// console.log("weeks:: " + getCurrentWeek(userData.StartDate));
// console.log("days:: " + getCurrentDay(userData.StartDate));
userData.StartDate = "2020-10-01";
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

  //remove WaistCirc form field if data has been entered for the current week
  userData.Week[getCurrentWeek(userData.StartDate)].Day.some(item => {  //array.some so that return "true" breaks loop
    if (item.WaistCirc) {
      console.log("WaistCirc form field removed...");
      document.getElementById("waistCircColumn").remove();
      return true;
    }
  });

  // === COACH RECOMMENDATION CHART ===
  //get latest recommendation object for charts
  var latestRec = getLatestRecommendation(userData)
  // only show if a recommendation exists
  if (!latestRec) {
    document.getElementById("CoachRecContainer").remove();
    document.getElementById("CoachRecHr").remove();
  } else {
    document.getElementById("updateDateText").innerHTML = "Last Updated: " + latestRec.ModifiedDate;
    //Normal Day
    document.getElementById("coach-NProteinAmount").innerHTML = latestRec.NormalDayProtein;
    document.getElementById("coach-NProteinRatio").innerHTML = (latestRec.NormalDayProtein/latestRec.NormalDayCalories).toPrecision(1) + "%";
    document.getElementById("coach-NCarbAmount").innerHTML = latestRec.NormalDayCarb;
    document.getElementById("coach-NCarbRatio").innerHTML = (latestRec.NormalDayCarb/latestRec.NormalDayCalories).toPrecision(1) + "%";
    document.getElementById("coach-NFatAmount").innerHTML = latestRec.NormalDayFat;
    document.getElementById("coach-NFatRatio").innerHTML = (latestRec.NormalDayFat/latestRec.NormalDayCalories).toPrecision(1) + "%";
    document.getElementById("coach-NCalories").innerHTML = latestRec.NormalDayCalories;
    document.getElementById("coach-NCaloriesRatio").innerHTML = ((latestRec.NormalDayProtein+latestRec.NormalDayCarb+latestRec.NormalDayFat)/latestRec.NormalDayCalories).toPrecision(1) + "%";
    //High Day
    document.getElementById("coach-HProteinAmount").innerHTML = latestRec.HighDayProtein;
    document.getElementById("coach-HProteinRatio").innerHTML = (latestRec.HighDayProtein/latestRec.HighDayCalories).toPrecision(1) + "%";
    document.getElementById("coach-HCarbAmount").innerHTML = latestRec.HighDayCarb;
    document.getElementById("coach-HCarbRatio").innerHTML = (latestRec.HighDayCarb/latestRec.HighDayCalories).toPrecision(1) + "%";
    document.getElementById("coach-HFatAmount").innerHTML = latestRec.HighDayFat;
    document.getElementById("coach-HFatRatio").innerHTML = (latestRec.HighDayFat/latestRec.HighDayCalories).toPrecision(1) + "%";
    document.getElementById("coach-HCalories").innerHTML = latestRec.HighDayCalories;
    document.getElementById("coach-HCaloriesRatio").innerHTML = ((latestRec.HighDayProtein+latestRec.HighDayCarb+latestRec.HighDayFat)/latestRec.HighDayCalories).toPrecision(1) + "%";
    //Low Day
    document.getElementById("coach-LProteinAmount").innerHTML = latestRec.LowDayProtein;
    document.getElementById("coach-LProteinRatio").innerHTML = (latestRec.LowDayProtein/latestRec.LowDayCalories).toPrecision(1) + "%";
    document.getElementById("coach-LCarbAmount").innerHTML = latestRec.LowDayCarb;
    document.getElementById("coach-LCarbRatio").innerHTML = (latestRec.LowDayCarb/latestRec.LowDayCalories).toPrecision(1) + "%";
    document.getElementById("coach-LFatAmount").innerHTML = latestRec.LowDayFat;
    document.getElementById("coach-LFatRatio").innerHTML = (latestRec.LowDayFat/latestRec.LowDayCalories).toPrecision(1) + "%";
    document.getElementById("coach-LCalories").innerHTML = latestRec.LowDayCalories;
    document.getElementById("coach-LCaloriesRatio").innerHTML = ((latestRec.LowDayProtein + latestRec.LowDayCarb + latestRec.LowDayFat) / latestRec.LowDayCalories).toPrecision(1) + "%";
    //HIIT
    document.getElementById("coach-HIITCurrentCardioSession").innerHTML = latestRec.HIITCurrentCardioSession;
    document.getElementById("coach-HIITCurrentCardioIntervals").innerHTML = latestRec.HIITCurrentCardioIntervals;
  }

  // === BASELINE CHART ===
  var baselineRec = userData.Recommendation[0]
  //Normal Day
  document.getElementById("NProteinAmount").innerHTML = baselineRec.NormalDayProtein;
  document.getElementById("NProteinRatio").innerHTML = (baselineRec.NormalDayProtein/baselineRec.NormalDayCalories).toPrecision(1) + "%";
  document.getElementById("NCarbAmount").innerHTML = baselineRec.NormalDayCarb;
  document.getElementById("NCarbRatio").innerHTML = (baselineRec.NormalDayCarb/baselineRec.NormalDayCalories).toPrecision(1) + "%";
  document.getElementById("NFatAmount").innerHTML = baselineRec.NormalDayFat;
  document.getElementById("NFatRatio").innerHTML = (baselineRec.NormalDayFat/baselineRec.NormalDayCalories).toPrecision(1) + "%";
  document.getElementById("NCalories").innerHTML = baselineRec.NormalDayCalories;
  document.getElementById("NCaloriesRatio").innerHTML = ((baselineRec.NormalDayProtein+baselineRec.NormalDayCarb+baselineRec.NormalDayFat)/baselineRec.NormalDayCalories).toPrecision(1) + "%";
  //High Day
  document.getElementById("HProteinAmount").innerHTML = baselineRec.HighDayProtein;
  document.getElementById("HProteinRatio").innerHTML = (baselineRec.HighDayProtein/baselineRec.HighDayCalories).toPrecision(1) + "%";
  document.getElementById("HCarbAmount").innerHTML = baselineRec.HighDayCarb;
  document.getElementById("HCarbRatio").innerHTML = (baselineRec.HighDayCarb/baselineRec.HighDayCalories).toPrecision(1) + "%";
  document.getElementById("HFatAmount").innerHTML = baselineRec.HighDayFat;
  document.getElementById("HFatRatio").innerHTML = (baselineRec.HighDayFat/baselineRec.HighDayCalories).toPrecision(1) + "%";
  document.getElementById("HCalories").innerHTML = baselineRec.HighDayCalories;
  document.getElementById("HCaloriesRatio").innerHTML = ((baselineRec.HighDayProtein+baselineRec.HighDayCarb+baselineRec.HighDayFat)/baselineRec.HighDayCalories).toPrecision(1) + "%";
  //Low Day
  document.getElementById("LProteinAmount").innerHTML = baselineRec.LowDayProtein;
  document.getElementById("LProteinRatio").innerHTML = (baselineRec.LowDayProtein/baselineRec.LowDayCalories).toPrecision(1) + "%";
  document.getElementById("LCarbAmount").innerHTML = baselineRec.LowDayCarb;
  document.getElementById("LCarbRatio").innerHTML = (baselineRec.LowDayCarb/baselineRec.LowDayCalories).toPrecision(1) + "%";
  document.getElementById("LFatAmount").innerHTML = baselineRec.LowDayFat;
  document.getElementById("LFatRatio").innerHTML = (baselineRec.LowDayFat/baselineRec.LowDayCalories).toPrecision(1) + "%";
  document.getElementById("LCalories").innerHTML = baselineRec.LowDayCalories;
  document.getElementById("LCaloriesRatio").innerHTML = ((baselineRec.LowDayProtein+baselineRec.LowDayCarb+baselineRec.LowDayFat)/baselineRec.LowDayCalories).toPrecision(1) + "%";
}

//==================================================
// Helper Functions

// returns latest recommendation object that has an "ModifiedDate"
function getLatestRecommendation(userData) {
  var latestRec = userData.Recommendation.filter(value => Object.keys(value).length !== 0).slice(-1)[0];
  if (latestRec.ModifiedDate) {
    return latestRec;
  }
  else {
    return null
  }
}

// returns the int for the current week as shown in history page
function getCurrentWeek(startingDate) {
  return Math.floor((new Date() - new Date(startingDate))/604800000);
}

// returns the int for the current day as shown in history page
function getCurrentDay(startingDate) {
  return Math.floor(((new Date() - new Date(startingDate))/86400000)%7);
}

//serialize form fields into json object
function serializeDailyUpdate(form) {
  var formJSON = {
    Fat: Number(document.getElementById("fat").value),
    Carbs: Number(document.getElementById("carbs").value),
    Protein: Number(document.getElementById("protein").value),
    Weight: Number(document.getElementById("weight").value),
    TotalCalories: Number(document.getElementById("totalCalories").value),
    DayCalories: document.getElementById("dayCalorie").value,
    Cardio: document.getElementById("cardio").value,
    WeightTraining: document.getElementById("weightTraining").value,
  };
  //add waistCirc field if present in form
  if (document.getElementById("waistCirc").value) {
    formJSON.WaistCirc = Number(document.getElementById("waistCirc").value);
  }
  return formJSON;
}

//submit form as JSON on "save" button click
function submitForm() {
  //check that form is valid before sending
  if (document.getElementById("DailyUpdateForm").checkValidity() == false) {
    console.log("Form Invalid.")
    return
  }
  //serialize form to JSON
  var dataObject = serializeDailyUpdate(document.getElementById("DailyUpdateForm"));
  var jsonData = JSON.stringify(dataObject);
  console.log(jsonData);  //FIXME remove
  //POST JSON to api
  var xmlHttp = new XMLHttpRequest();
  xmlHttp.open( "POST", "http://localhost:8888/userDaily/" + getCurrentWeek(userData.StartDate) + "/" + getCurrentDay(userData.StartDate), false );
  // xmlHttp.setRequestHeader("Session-Token",getCookie("Session-Token"));
  xmlHttp.setRequestHeader("Session-Token",myToken); //FIXME use correct session-token
  xmlHttp.send(jsonData);
  console.log("Server response: " + xmlHttp.responseText);
  //show note that save was successful
  document.getElementById("SaveConfirmationText").innerHTML = "&nbsp; &nbsp; &nbsp; Saved!";
}
