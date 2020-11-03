var myToken = "test"; //FIXME remove

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


// updates baseline charts with user data
function updateCharts(userData) {
  // redirect to profile page on empty data
  if (userData.Recommendation[0].NormalDayCalories == 0) {
    console.log("Baseline data blank, redirecting to profile...")
    // window.location.replace('http://localhost:5500/profile');  //FIXME
  }
  // === COACH RECOMMENDATION CHART ===
  // //get latest recommendation object for charts
  var latestRec = getLatestRecommendation(userData)
  if (latestRec == null) {
    document.getElementById("CoachRecContainer").remove();
    document.getElementById("CoachRecHr").remove();
  } else {
    document.getElementById("updateDateText").innerHTML = "Updated: " + latestRec.ModifiedDate;
    //Normal Day
    document.getElementById("coach-NProteinAmount").innerHTML = latestRec.NormalDayProtein;
    // document.getElementById("coach-NProteinRatio").innerHTML = latestRec.NProteinRatio;
    document.getElementById("coach-NCarbAmount").innerHTML = latestRec.NormalDayCarb;
    // document.getElementById("coach-NCarbRatio").innerHTML = latestRec.NCarbRatio;
    document.getElementById("coach-NFatAmount").innerHTML = latestRec.NormalDayFat;
    // document.getElementById("coach-NFatRatio").innerHTML = latestRec.NFatRatio;
    document.getElementById("coach-NCalories").innerHTML = latestRec.NormalDayCalories;
    //High Day
    document.getElementById("coach-HProteinAmount").innerHTML = latestRec.HighDayProtein;
    // document.getElementById("coach-HProteinRatio").innerHTML = latestRec.HProteinRatio;
    document.getElementById("coach-HCarbAmount").innerHTML = latestRec.HighDayCarb;
    // document.getElementById("coach-HCarbRatio").innerHTML = latestRec.HCarbRatio;
    document.getElementById("coach-HFatAmount").innerHTML = latestRec.HighDayFat;
    // document.getElementById("coach-HFatRatio").innerHTML = latestRec.HFatRatio;
    document.getElementById("coach-HCalories").innerHTML = latestRec.HighDayCalories;
    //Low Day
    document.getElementById("coach-LProteinAmount").innerHTML = latestRec.LowDayProtein;
    // document.getElementById("coach-LProteinRatio").innerHTML = latestRec.LProteinRatio;
    document.getElementById("coach-LCarbAmount").innerHTML = latestRec.LowDayCarb;
    // document.getElementById("coach-LCarbRatio").innerHTML = latestRec.LCarbRatio;
    document.getElementById("coach-LFatAmount").innerHTML = latestRec.LowDayFat;
    // document.getElementById("coach-LFatRatio").innerHTML = latestRec.LFatRatio;
    document.getElementById("coach-LCalories").innerHTML = latestRec.LowDayCalories;
  }

  // === BASELINE CHART === //TODO fix ratio references
  //Normal Day
  document.getElementById("NProteinAmount").innerHTML = userData.Recommendation[0].NormalDayProtein;
  // document.getElementById("NProteinRatio").innerHTML = userData.Recommendation[0].NProteinRatio;
  document.getElementById("NCarbAmount").innerHTML = userData.Recommendation[0].NormalDayCarb;
  // document.getElementById("NCarbRatio").innerHTML = userData.Recommendation[0].NCarbRatio;
  document.getElementById("NFatAmount").innerHTML = userData.Recommendation[0].NormalDayFat;
  // document.getElementById("NFatRatio").innerHTML = userData.Recommendation[0].NFatRatio;
  document.getElementById("NCalories").innerHTML = userData.Recommendation[0].NormalDayCalories;
  //High Day
  document.getElementById("HProteinAmount").innerHTML = userData.Recommendation[0].HighDayProtein;
  // document.getElementById("HProteinRatio").innerHTML = userData.Recommendation[0].HProteinRatio;
  document.getElementById("HCarbAmount").innerHTML = userData.Recommendation[0].HighDayCarb;
  // document.getElementById("HCarbRatio").innerHTML = userData.Recommendation[0].HCarbRatio;
  document.getElementById("HFatAmount").innerHTML = userData.Recommendation[0].HighDayFat;
  // document.getElementById("HFatRatio").innerHTML = userData.Recommendation[0].HFatRatio;
  document.getElementById("HCalories").innerHTML = userData.Recommendation[0].HighDayCalories;
  //Low Day
  document.getElementById("LProteinAmount").innerHTML = userData.Recommendation[0].LowDayProtein;
  // document.getElementById("LProteinRatio").innerHTML = userData.Recommendation[0].LProteinRatio;
  document.getElementById("LCarbAmount").innerHTML = userData.Recommendation[0].LowDayCarb;
  // document.getElementById("LCarbRatio").innerHTML = userData.Recommendation[0].LCarbRatio;
  document.getElementById("LFatAmount").innerHTML = userData.Recommendation[0].LowDayFat;
  // document.getElementById("LFatRatio").innerHTML = userData.Recommendation[0].LFatRatio;
  document.getElementById("LCalories").innerHTML = userData.Recommendation[0].LowDayCalories;
}


// returns latest recommendation object that is not the baseline (first obj in array)
function getLatestRecommendation(userData) {
  //make copy of user data to manipulate
  var tmpUserData = JSON.parse(JSON.stringify(userData.Recommendation));
  tmpUserData.shift();
  return tmpUserData.filter(value => Object.keys(value).length !== 0).slice(-1)[0];
}
