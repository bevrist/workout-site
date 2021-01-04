
//FIXME update this
const queryString = window.location.search;
const urlParams = new URLSearchParams(queryString);
const userUid = urlParams.get('uid')
//if uid url parameter is empty redirect to /admin/users
if (userUid == "") {
  console.log("no uid URL parameter, redirecting to admin/users...")
      // window.location.replace('http://localhost:5500/admin/users');
}
//get user data from api and store JSON in "userData"
var xmlHttp = new XMLHttpRequest();
xmlHttp.open("GET", "http://localhost:8888/admin/userInfo", false);
xmlHttp.setRequestHeader("User-UID",userUid);
// xmlHttp.setRequestHeader("Session-Token",getCookie("Session-Token"));
xmlHttp.setRequestHeader("Session-Token", "testUID"); //FIXME use correct session-token
xmlHttp.send(null);
var userData = JSON.parse(xmlHttp.responseText);
//abort if data is invalid
if (userData.FirstName == "") {
  document.getElementById("UserNameText").innerHTML = "Invalid Request"
  document.getElementById("myContainer").remove();
  exit();
}


document.title = "ADMIN-" + userData.FirstName + " " + userData.LastName;

//link to current week
var currentWeekLink = "#";

// populate starting date
var startDate = new Date(userData.StartDate);
document.getElementById("UserNameText").innerHTML = "Viewing User: " + userData.FirstName + " " + userData.LastName;

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
  // document.getElementById("formSaveButton-0").onclick = "submitForm(" + weekNum + ")";
  document.getElementById("formSaveButton-0").setAttribute("onclick", "submitForm(" + weekNum + ")");
  document.getElementById("formSaveButton-0").id = "formSaveButton-" + weekNum;
  // delete save buttons
  document.getElementById("formSaveButton-"+ weekNum).remove();
  document.getElementById("SaveConfirmationText-0").innerHTML = "";
  document.getElementById("SaveConfirmationText-0").id = "SaveConfirmationText-" + weekNum;
  document.getElementById("CoachRecContainer-0").id = "CoachRecContainer-" + weekNum;
  document.getElementById("myContainer").id = "week" + weekNum;
  populateCoachRecChart(weekNum);
  document.body.appendChild(divClone);


  makeWeekChart(weekNum, 0);
  makeWeekChart(weekNum, 1);
  makeWeekChart(weekNum, 2);
  makeWeekChart(weekNum, 3);
  makeWeekChart(weekNum, 4);
  makeWeekChart(weekNum, 5);
  makeWeekChart(weekNum, 6);
  makeWeekChart(null, null);
}

//delete unused rec charts and populate ones with data
function populateCoachRecChart(weekNum) {
  weekNum--;
  var coachRec = userData.Recommendation[weekNum];
  weekNum++;
  if (Object.keys(coachRec).length === 0 || !coachRec.ModifiedDate) {
    document.getElementById("CoachRecContainer-" + weekNum).remove();
  } else {
    document.getElementById("updateDateText").innerHTML = "Last Updated: " + coachRec.ModifiedDate;
    document.getElementById("updateDateText").removeAttribute("id");
    document.getElementById("CoachRecHeading").innerHTML = "Coach Rec Week " + weekNum;
    document.getElementById("CoachRecHeading").removeAttribute("id");
    //Normal Day
    document.getElementById("coach-NProteinAmount").innerHTML = coachRec.NormalDayProtein;
    document.getElementById("coach-NProteinAmount").removeAttribute("id");
    document.getElementById("coach-NProteinRatio").innerHTML = (coachRec.NormalDayProtein / coachRec.NormalDayCalories).toPrecision(1) + "%";
    document.getElementById("coach-NProteinRatio").removeAttribute("id");
    document.getElementById("coach-NCarbAmount").innerHTML = coachRec.NormalDayCarb;
    document.getElementById("coach-NCarbAmount").removeAttribute("id");
    document.getElementById("coach-NCarbRatio").innerHTML = (coachRec.NormalDayCarb / coachRec.NormalDayCalories).toPrecision(1) + "%";
    document.getElementById("coach-NCarbRatio").removeAttribute("id");
    document.getElementById("coach-NFatAmount").innerHTML = coachRec.NormalDayFat;
    document.getElementById("coach-NFatAmount").removeAttribute("id");
    document.getElementById("coach-NFatRatio").innerHTML = (coachRec.NormalDayFat / coachRec.NormalDayCalories).toPrecision(1) + "%";
    document.getElementById("coach-NFatRatio").removeAttribute("id");
    document.getElementById("coach-NCalories").innerHTML = coachRec.NormalDayCalories;
    document.getElementById("coach-NCalories").removeAttribute("id");
    document.getElementById("coach-NCaloriesRatio").innerHTML = ((coachRec.NormalDayProtein + coachRec.NormalDayCarb + coachRec.NormalDayFat) / coachRec.NormalDayCalories).toPrecision(1) + "%";
    document.getElementById("coach-NCaloriesRatio").removeAttribute("id");
    //High Day
    document.getElementById("coach-HProteinAmount").innerHTML = coachRec.HighDayProtein;
    document.getElementById("coach-HProteinAmount").removeAttribute("id");
    document.getElementById("coach-HProteinRatio").innerHTML = (coachRec.HighDayProtein / coachRec.HighDayCalories).toPrecision(1) + "%";
    document.getElementById("coach-HProteinRatio").removeAttribute("id");
    document.getElementById("coach-HCarbAmount").innerHTML = coachRec.HighDayCarb;
    document.getElementById("coach-HCarbAmount").removeAttribute("id");
    document.getElementById("coach-HCarbRatio").innerHTML = (coachRec.HighDayCarb / coachRec.HighDayCalories).toPrecision(1) + "%";
    document.getElementById("coach-HCarbRatio").removeAttribute("id");
    document.getElementById("coach-HFatAmount").innerHTML = coachRec.HighDayFat;
    document.getElementById("coach-HFatAmount").removeAttribute("id");
    document.getElementById("coach-HFatRatio").innerHTML = (coachRec.HighDayFat / coachRec.HighDayCalories).toPrecision(1) + "%";
    document.getElementById("coach-HFatRatio").removeAttribute("id");
    document.getElementById("coach-HCalories").innerHTML = coachRec.HighDayCalories;
    document.getElementById("coach-HCalories").removeAttribute("id");
    document.getElementById("coach-HCaloriesRatio").innerHTML = ((coachRec.HighDayProtein + coachRec.HighDayCarb + coachRec.HighDayFat) / coachRec.HighDayCalories).toPrecision(1) + "%";
    document.getElementById("coach-HCaloriesRatio").removeAttribute("id");
    //Low Day
    document.getElementById("coach-LProteinAmount").innerHTML = coachRec.LowDayProtein;
    document.getElementById("coach-LProteinAmount").removeAttribute("id");
    document.getElementById("coach-LProteinRatio").innerHTML = (coachRec.LowDayProtein / coachRec.LowDayCalories).toPrecision(1) + "%";
    document.getElementById("coach-LProteinRatio").removeAttribute("id");
    document.getElementById("coach-LCarbAmount").innerHTML = coachRec.LowDayCarb;
    document.getElementById("coach-LCarbAmount").removeAttribute("id");
    document.getElementById("coach-LCarbRatio").innerHTML = (coachRec.LowDayCarb / coachRec.LowDayCalories).toPrecision(1) + "%";
    document.getElementById("coach-LCarbRatio").removeAttribute("id");
    document.getElementById("coach-LFatAmount").innerHTML = coachRec.LowDayFat;
    document.getElementById("coach-LFatAmount").removeAttribute("id");
    document.getElementById("coach-LFatRatio").innerHTML = (coachRec.LowDayFat / coachRec.LowDayCalories).toPrecision(1) + "%";
    document.getElementById("coach-LFatRatio").removeAttribute("id");
    document.getElementById("coach-LCalories").innerHTML = coachRec.LowDayCalories;
    document.getElementById("coach-LCalories").removeAttribute("id");
    document.getElementById("coach-LCaloriesRatio").innerHTML = ((coachRec.LowDayProtein + coachRec.LowDayCarb + coachRec.LowDayFat) / coachRec.LowDayCalories).toPrecision(1) + "%";
    document.getElementById("coach-LCaloriesRatio").removeAttribute("id");
    //HIIT
    if (coachRec.HIITCurrentCardioSession == null) {
      document.getElementById("hiit-row").remove();
    } else {
      document.getElementById("hiit-row").removeAttribute("id");
      document.getElementById("coach-HIITCurrentCardioSession").innerHTML = coachRec.HIITCurrentCardioSession;
      document.getElementById("coach-HIITCurrentCardioSession").removeAttribute("id");
      document.getElementById("coach-HIITCurrentCardioIntervals").innerHTML = coachRec.HIITCurrentCardioIntervals;
      document.getElementById("coach-HIITCurrentCardioIntervals").removeAttribute("id");
    }
  }
}

var dayOffset = 1;
// populates form boxes for each week from user data
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
  // highlight current date on chart when generating, else just display date
  var todayDate = new Date();
  todaysShortDate = todayDate.getMonth() + 1 + "/" + todayDate.getDate();
  formShortDate = currDate.getMonth() + 1 + "/" + currDate.getDate();
  if (todaysShortDate == formShortDate) {
    document.getElementById("date").childNodes[0].innerHTML = "<mark>"+formShortDate+"</mark>";
  } else {
    document.getElementById("date").childNodes[0].innerHTML = formShortDate;
  }
  document.getElementById("date").removeAttribute("id");
  //check if date is current date, set page anchor link to jump down to week on page completion
  if (currDate.toISOString().split('T')[0] == todayDate.toISOString().split('T')[0]) {
    currentWeekLink = "#formLink-" + (weekNum + 1);
  }
  //day calories
  switch (userData.Week[weekNum].Day[dayNum].DayCalories) {
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
  document.getElementById("day").id = "day-" + weekNum+dayNum;
  document.getElementById("dayNorm").removeAttribute("id");
  document.getElementById("dayLow").removeAttribute("id");
  document.getElementById("dayHigh").removeAttribute("id");

  //cardio
  switch (userData.Week[weekNum].Day[dayNum].Cardio) {
    case "missed":
      document.getElementById("cardioMiss").selected = "selected"
      break;
    case "hit":
      document.getElementById("cardioHit").selected = "selected"
      break;
  }
  document.getElementById("cardio").id = "cardio-" + weekNum+dayNum;
  document.getElementById("cardioMiss").removeAttribute("id");
  document.getElementById("cardioHit").removeAttribute("id");

  //weight training
  switch (userData.Week[weekNum].Day[dayNum].WeightTraining) {
    case "yes":
      document.getElementById("trainingYes").selected = "selected"
      break;
    case "no":
      document.getElementById("trainingNo").selected = "selected"
      break;
  }
  document.getElementById("training").id = "training-" + weekNum + dayNum;
  document.getElementById("trainingNo").removeAttribute("id");
  document.getElementById("trainingYes").removeAttribute("id");

  document.getElementById("fat").value = userData.Week[weekNum].Day[dayNum].Fat;
  document.getElementById("fat").id = "fat-" + weekNum + dayNum;
  document.getElementById("carbs").value = userData.Week[weekNum].Day[dayNum].Carbs;
  document.getElementById("carbs").id = "carbs-" + weekNum + dayNum;
  document.getElementById("protein").value = userData.Week[weekNum].Day[dayNum].Protein;
  document.getElementById("protein").id = "protein-" + weekNum + dayNum;
  document.getElementById("calories").value = userData.Week[weekNum].Day[dayNum].TotalCalories;
  document.getElementById("calories").id = "calories-" + weekNum + dayNum;
  document.getElementById("weight").value = userData.Week[weekNum].Day[dayNum].Weight;
  document.getElementById("weight").id = "weight-" + weekNum + dayNum;
  if ("WaistCirc" in userData.Week[weekNum ].Day[dayNum]) {
    document.getElementById("waist-circ").value = userData.Week[weekNum].Day[dayNum].WaistCirc;
  }
  document.getElementById("waist-circ").id = "waist-circ-" + weekNum + dayNum;
  //clone table row
  document.getElementById("tableRow").id = "tableRowNext";
  document.getElementById("tableRowNext").parentElement.appendChild(divClone2);
  document.getElementById("tableRowNext").removeAttribute("id");
}

//clone through week 24
for (var i = 1; i <= userData.Week.length; i++) {
  cloneForm(i);
}
document.getElementById("myContainer").remove(); // delete empty clone form at end

// //serialize form fields into json object
// function serializeWeekForm(weekNum) {
//   let DayArr = [
//     serializeDayForm(weekNum, 0),
//     serializeDayForm(weekNum, 1),
//     serializeDayForm(weekNum, 2),
//     serializeDayForm(weekNum, 3),
//     serializeDayForm(weekNum, 4),
//     serializeDayForm(weekNum, 5),
//     serializeDayForm(weekNum, 6),
//   ];
//   var formJSON = {
//     Day: DayArr,
//   }
//   // console.log(formJSON);
//   return formJSON;
// }

// //serialize a day object in the form
// function serializeDayForm(weekNum, dayNum) {
//   var dayJSON = {
//     Fat: Number(document.getElementById("fat-" + weekNum+dayNum).value),
//     Carbs: Number(document.getElementById("carbs-" + weekNum+dayNum).value),
//     Protein: Number(document.getElementById("protein-" + weekNum+dayNum).value),
//     Weight: Number(document.getElementById("weight-" + weekNum+dayNum).value),
//     TotalCalories: Number(document.getElementById("calories-" + weekNum+dayNum).value),
//     DayCalories: document.getElementById("day-" + weekNum+dayNum).value,
//     Cardio: document.getElementById("cardio-" + weekNum+dayNum).value,
//     WeightTraining: document.getElementById("training-" + weekNum + dayNum).value,
//   }
//   if (document.getElementById("waist-circ-" + weekNum + dayNum).value != 0) {
//     dayJSON.WaistCirc = Number(document.getElementById("waist-circ-" + weekNum + dayNum).value)
//   }
//   return dayJSON;
// }

// //submit form as JSON on "save" button click
// function submitForm(weekNum) {
//   //stop from from refreshing page
//   event.preventDefault(); //FIXME see if can remove
//   //serialize form to JSON
//   var dataObject = serializeWeekForm(weekNum-1);
//   var jsonData = JSON.stringify(dataObject);
//   //POST JSON to api
//   var xmlHttp = new XMLHttpRequest();
//   xmlHttp.open( "POST", "http://localhost:8888/userWeekly/" + (weekNum-1), false );
//   xmlHttp.setRequestHeader("Session-Token",getCookie("Session-Token"));
//   // xmlHttp.setRequestHeader("Session-Token",myToken); //FIXME use correct session-token
//   xmlHttp.send(jsonData);
//   console.log(jsonData);
//   console.log("Server response: " + xmlHttp.responseText);
//   //show note that save was successful
//   document.getElementById("SaveConfirmationText-" + weekNum).innerHTML = "&nbsp; &nbsp; &nbsp; Week " + weekNum + " Saved!";
// }
