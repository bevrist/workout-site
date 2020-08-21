//FIXME remove
var userData = { "first_name": "Anthony", "last_name": "Hanna", "weight": 215, "waistcirc": 35.5, "heightinches": 75, "leanbodymass": "15%", "age": "20", "gender": "male", "startDate": "2020-08-15", "week": [{ "Day": [{ "fat": 100, "carbs": 100, "protein": 100, "total_calories": 300, "day_calorie": "high", "weight": 123, "cardio": "hit", "weight_training": "no" }, { "fat": 22, "carbs": 22, "protein": 22, "total_calories": 22, "day_calorie": "normal", "weight": 123, "cardio": "missed", "weight_training": "yes" }] }, { "Day": [{ "fat": 10, "carbs": 10, "protein": 10, "total_calories": 30, "day_calorie": "high", "weight": 123, "cardio": "none", "weight_training": "no" }, { "fat": 11, "carbs": 11, "protein": 11, "total_calories": 31, "day_calorie": "low", "weight": 123, "cardio": "hit", "weight_training": "no" }] }, { "Day": [{ "fat": 12, "carbs": 12, "protein": 12, "total_calories": 32, "day_calorie": "normal", "weight": 123, "cardio": "missed", "weight_training": "yes" }, { "fat": 14, "carbs": 14, "protein": 14, "total_calories": 34, "day_calorie": "low", "weight": 123, "cardio": "hit", "weight_training": "no" }] }] }
//FIXME REMOVE
document.cookie = "Session-Token=testUID; SameSite=Strict;";  //FIXME REMOVE

//link to current week
var currentWeekLink = "#";

// populate starting date
var startDate = new Date(userData.startDate);
document.getElementById("StartDate").value = startDate.toISOString().split('T')[0];
document.getElementById("Form-Session-Token-StartDate").value = getCookie("Session-Token");

//helper function for adding days to date
Date.prototype.addDays = function (days) {
  var date = new Date(this.valueOf());
  date.setDate(date.getDate() + days);
  return date;
}

// this function clones a unique form and renames items in form
function cloneForm(weekNum) {
  var myDiv = document.getElementById("myContainer");
  var divClone = myDiv.cloneNode(true);

  // document.getElementById("formWeek-0").action = "http://localhost:8080/userWeeklyTracking/" + weekNum;
  document.getElementById("formWeek-0").id = "formWeek-" + weekNum;
  document.getElementById("formLink-0").name = "formLink-" + weekNum;
  document.getElementById("formLink-0").id = "formLink-" + weekNum;
  document.getElementById("weekHeading-0").innerHTML = "Week " + weekNum;
  document.getElementById("weekHeading-0").id = "weekHeading-" + weekNum;
  document.getElementById("formSaveButton-0").innerHTML = "Save Week " + weekNum;
  document.getElementById("formSaveButton-0").id = "formSaveButton-" + weekNum;
  document.getElementById("Form-Session-Token-0").value = getCookie("Session-Token");
  document.getElementById("Form-Session-Token-0").id = "Form-Session-Token-" + weekNum;
  document.getElementById("myContainer").id = "week" + weekNum;
  document.body.appendChild(divClone);


  makeWeekChart(weekNum, 0);
  makeWeekChart(weekNum, 1);
  // makeWeekChart(weekNum,2);  //FIXME uncomment
  // makeWeekChart(weekNum,3);
  // makeWeekChart(weekNum,4);
  // makeWeekChart(weekNum,5);
  // makeWeekChart(weekNum,6);
  makeWeekChart(null, null);
}

var dayOffset = 1;
function makeWeekChart(weekNum, dayNum) {
  //if weekNum is null: remove last tableRow, return
  if (weekNum == null) {
    document.getElementById("tableRow").remove();
    // jump to particular location on page
    setTimeout(function () { location.href = currentWeekLink; }, 1000);
    return;
  }
  weekNum = weekNum - 1;
  var myDiv2 = document.getElementById("tableRow");
  var divClone2 = myDiv2.cloneNode(true);
  //populate existing week row with data and remove id's
  // calculate date for row
  var currDate = startDate.addDays(dayOffset++);
  document.getElementById("date").childNodes[0].innerHTML = currDate.getMonth() + 1 + "/" + currDate.getDate();
  document.getElementById("date").removeAttribute("id");
  //check if date is current date, set link to jump down to week on page completion
  var todayDate = new Date();
  if (currDate.toISOString().split('T')[0] == todayDate.toISOString().split('T')[0]) {
    currentWeekLink = "#formLink-" + (weekNum + 1);
  }
  //TODO: training //check for and store week of current date
  //day calories
  switch (userData.week[weekNum].Day[dayNum].day_calorie) {
    case "normal":
      document.getElementById("dayNorm").selected = "selected"
      break;
    case "low":
      document.getElementById("dayLow").selected = "selected"
      break;
    case "high":
      document.getElementById("dayHigh").selected = "selected"
      break;
  }
  document.getElementById("dayNorm").removeAttribute("id");
  document.getElementById("dayLow").removeAttribute("id");
  document.getElementById("dayHigh").removeAttribute("id");

  //cardio
  switch (userData.week[weekNum].Day[dayNum].cardio) {
    case "none":
      document.getElementById("cardioNone").selected = "selected"
      break;
    case "missed":
      document.getElementById("cardioMiss").selected = "selected"
      break;
    case "hit":
      document.getElementById("cardioHit").selected = "selected"
      break;
  }
  document.getElementById("cardioNone").removeAttribute("id");
  document.getElementById("cardioMiss").removeAttribute("id");
  document.getElementById("cardioHit").removeAttribute("id");

  //weight training
  switch (userData.week[weekNum].Day[dayNum].weight_training) {
    case "yes":
      document.getElementById("trainingYes").selected = "selected"
      break;
    case "no":
      document.getElementById("trainingNo").selected = "selected"
      break;
  }
  document.getElementById("trainingNo").removeAttribute("id");
  document.getElementById("trainingYes").removeAttribute("id");

  document.getElementById("fat").value = userData.week[weekNum].Day[dayNum].fat;
  document.getElementById("fat").removeAttribute("id");
  document.getElementById("carbs").value = userData.week[weekNum].Day[dayNum].carbs;
  document.getElementById("carbs").removeAttribute("id");
  document.getElementById("protein").value = userData.week[weekNum].Day[dayNum].protein;
  document.getElementById("protein").removeAttribute("id");
  document.getElementById("calories").value = userData.week[weekNum].Day[dayNum].total_calories;
  document.getElementById("calories").removeAttribute("id");
  document.getElementById("weight").value = userData.week[weekNum].Day[dayNum].weight;
  document.getElementById("weight").removeAttribute("id");
  //clone table row
  document.getElementById("tableRow").id = "tableRowNext";
  document.getElementById("tableRowNext").parentElement.appendChild(divClone2);
  document.getElementById("tableRowNext").removeAttribute("id");
}

//clone through week 24
for (var i = 1; i <= userData.week.length; i++) {
  cloneForm(i);
}
document.getElementById("myContainer").remove(); // delete empty clone form at end

// //get UserProfile json object, call other functions on complete
// fetch("http://localhost:8080/userWeeklyTracking/" + num, {
//   headers: {
//     'Session-Token': getCookie("Session-Token"),
//   }
// }).then(function (response) {
//   response.json().then(function (data) {
//     updateForm(data);
//   });
// });
