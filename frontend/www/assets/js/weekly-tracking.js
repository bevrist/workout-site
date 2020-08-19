// add Session-Token to hidden form field
// document.getElementById("Form-Session-Token").value = getCookie("Session-Token");  //FIXME

//FIXME remove
var userData = { "first_name": "Anthony", "last_name": "Hanna", "weight": 215, "waistcirc": 35.5, "heightinches": 75, "leanbodymass": "15%", "age": "20", "gender": "male", "week": [{ "Day": [{ "fat": 100, "carbs": 100, "protein": 100, "total_calories": 300, "day_calorie": "normal", "weight": 123, "cardio": "missed", "weight_training": "yes" }, { "fat": 22, "carbs": 22, "protein": 22, "total_calories": 22, "day_calorie": "normal", "weight": 123, "cardio": "missed", "weight_training": "yes" }] }, { "Day": [{ "fat": 10, "carbs": 10, "protein": 10, "total_calories": 30, "day_calorie": "normal", "weight": 123, "cardio": "missed", "weight_training": "yes" }, { "fat": 11, "carbs": 11, "protein": 11, "total_calories": 31, "day_calorie": "normal", "weight": 123, "cardio": "missed", "weight_training": "yes" }] }, { "Day": [{ "fat": 12, "carbs": 12, "protein": 12, "total_calories": 32, "day_calorie": "normal", "weight": 123, "cardio": "missed", "weight_training": "yes" }, { "fat": 14, "carbs": 14, "protein": 14, "total_calories": 34, "day_calorie": "normal", "weight": 123, "cardio": "missed", "weight_training": "yes" }] }] }

// this function clones a unique form and renames items in form
function cloneForm(weekNum) {
  var myDiv = document.getElementById("myContainer");
  var divClone = myDiv.cloneNode(true);

  // document.getElementById("formWeek-0").action = "http://localhost:8080/userWeeklyTracking/" + weekNum;
  document.getElementById("formWeek-0").id = "formWeek-" + weekNum;
  document.getElementById("formLink-0").name = "asdf" + weekNum;
  document.getElementById("formLink-0").id = "formLink-" + weekNum;
  document.getElementById("weekHeading-0").innerHTML = "Week " + weekNum;
  document.getElementById("weekHeading-0").id = "weekHeading-" + weekNum;
  document.getElementById("Form-Session-Token-0").value = getCookie("Session-Token");
  document.getElementById("Form-Session-Token-0").id = "Form-Session-Token-" + weekNum;
  document.getElementById("myContainer").id = "week" + weekNum;
  document.body.appendChild(divClone);

  makeWeekChart(weekNum, 0);
  makeWeekChart(weekNum, 1);
  // makeWeekChart(weekNum,2);  //FIXME
  // makeWeekChart(weekNum,3);
  // makeWeekChart(weekNum,4);
  // makeWeekChart(weekNum,5);
  // makeWeekChart(weekNum,6);
  makeWeekChart(null, null);
}

function makeWeekChart(weekNum, dayNum) {
  //if weeknum is null: dont appendChild, return
  if (weekNum == null) {
    document.getElementById("tableRow").remove();
    return;
  }
  weekNum = weekNum - 1;
  var myDiv2 = document.getElementById("tableRow");
  var divClone2 = myDiv2.cloneNode(true);
  //populate existing week row with data and remove id
  //TODO: date, day, cardio, training //check for and store week of current date
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

cloneForm(1);
cloneForm(2);
cloneForm(3);
//clone through week 24
document.getElementById("myContainer").remove(); // delete empty clone form at end

// jump to particular location on page //TODO: calculate current week and redirect here
// location.href = "#asdf9";


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
