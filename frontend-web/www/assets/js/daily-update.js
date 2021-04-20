document.title = "Daily Update";

var userData = "";

//get UserProfile json object from server, call other functions on complete
fetch("http://localhost:8888/userInfo").then(function (response) {
  if (response.status === 401) {
    //redirect to sign in on auth failure
    console.log("auth fail. redirecting...");
    window.location.href = "http://localhost:5500/auth";
  } else {
    response.json().then(function (userData) {
      //if new user, direct to profile
      if (isNewUser(userData) == true) {
        console.log("user baseline missing, redirecting to profile page...");
        // window.location.replace('http://localhost:5500/profile');
      }
      // only update form if data already exists
      if (userData.FirstName != "") {
        document.title =
          "Daily Update - " + userData.FirstName + " " + userData.LastName;
        updateCharts(userData);
      }
    });
  }
});

// update "Todays Date" text
var todayLocalDate = new Date().getTimezoneOffset() * 60000;
var localISOTime = new Date(Date.now() - todayLocalDate)
  .toISOString()
  .split("T")[0];
document.getElementById("TodayDateText").innerHTML =
  "Today's Date: " + localISOTime; //TODO add note for too early too late users
//TODO disable form when out of active time <fieldset disabled="disabled">

// updates baseline & recommendation charts with user data
function updateCharts(userData) {
  // redirect to profile page on empty data
  if (userData.Recommendation[0].NormalDayCalories == 0) {
    console.log("Baseline data blank, redirecting to profile...");
    // window.location.replace('http://localhost:5500/profile');
  }

  //remove WaistCirc form field if data has been entered for the current week
  userData.Week[getCurrentWeek(userData.StartDate)].Day.some((item) => {
    //array.some so that return "true" breaks loop
    if (item.WaistCirc) {
      console.log("WaistCirc form field removed...");
      document.getElementById("waistCircColumn").remove();
      return true;
    }
  });

  // === COACH RECOMMENDATION CHART ===
  //get latest recommendation object for charts
  var latestRec = getLatestRecommendation(userData);
  // only show if a recommendation exists
  if (!latestRec) {
    document.getElementById("CoachRecContainer").remove();
    document.getElementById("CoachRecHr").remove();
  } else {
    document.getElementById("updateDateText").innerHTML =
      "Last Updated: " + latestRec.ModifiedDate;
    //Normal Day
    document.getElementById("coach-NProteinAmount").innerHTML =
      latestRec.NormalDayProtein;
    document.getElementById("coach-NProteinRatio").innerHTML =
      (latestRec.NormalDayProtein / latestRec.NormalDayCalories).toPrecision(
        1
      ) + "%";
    document.getElementById("coach-NCarbAmount").innerHTML =
      latestRec.NormalDayCarb;
    document.getElementById("coach-NCarbRatio").innerHTML =
      (latestRec.NormalDayCarb / latestRec.NormalDayCalories).toPrecision(1) +
      "%";
    document.getElementById("coach-NFatAmount").innerHTML =
      latestRec.NormalDayFat;
    document.getElementById("coach-NFatRatio").innerHTML =
      (latestRec.NormalDayFat / latestRec.NormalDayCalories).toPrecision(1) +
      "%";
    document.getElementById("coach-NCalories").innerHTML =
      latestRec.NormalDayCalories;
    document.getElementById("coach-NCaloriesRatio").innerHTML =
      (
        (latestRec.NormalDayProtein +
          latestRec.NormalDayCarb +
          latestRec.NormalDayFat) /
        latestRec.NormalDayCalories
      ).toPrecision(1) + "%";
    //High Day
    document.getElementById("coach-HProteinAmount").innerHTML =
      latestRec.HighDayProtein;
    document.getElementById("coach-HProteinRatio").innerHTML =
      (latestRec.HighDayProtein / latestRec.HighDayCalories).toPrecision(1) +
      "%";
    document.getElementById("coach-HCarbAmount").innerHTML =
      latestRec.HighDayCarb;
    document.getElementById("coach-HCarbRatio").innerHTML =
      (latestRec.HighDayCarb / latestRec.HighDayCalories).toPrecision(1) + "%";
    document.getElementById("coach-HFatAmount").innerHTML =
      latestRec.HighDayFat;
    document.getElementById("coach-HFatRatio").innerHTML =
      (latestRec.HighDayFat / latestRec.HighDayCalories).toPrecision(1) + "%";
    document.getElementById("coach-HCalories").innerHTML =
      latestRec.HighDayCalories;
    document.getElementById("coach-HCaloriesRatio").innerHTML =
      (
        (latestRec.HighDayProtein +
          latestRec.HighDayCarb +
          latestRec.HighDayFat) /
        latestRec.HighDayCalories
      ).toPrecision(1) + "%";
    //Low Day
    document.getElementById("coach-LProteinAmount").innerHTML =
      latestRec.LowDayProtein;
    document.getElementById("coach-LProteinRatio").innerHTML =
      (latestRec.LowDayProtein / latestRec.LowDayCalories).toPrecision(1) + "%";
    document.getElementById("coach-LCarbAmount").innerHTML =
      latestRec.LowDayCarb;
    document.getElementById("coach-LCarbRatio").innerHTML =
      (latestRec.LowDayCarb / latestRec.LowDayCalories).toPrecision(1) + "%";
    document.getElementById("coach-LFatAmount").innerHTML = latestRec.LowDayFat;
    document.getElementById("coach-LFatRatio").innerHTML =
      (latestRec.LowDayFat / latestRec.LowDayCalories).toPrecision(1) + "%";
    document.getElementById("coach-LCalories").innerHTML =
      latestRec.LowDayCalories;
    document.getElementById("coach-LCaloriesRatio").innerHTML =
      (
        (latestRec.LowDayProtein + latestRec.LowDayCarb + latestRec.LowDayFat) /
        latestRec.LowDayCalories
      ).toPrecision(1) + "%";
    //HIIT
    document.getElementById("coach-HIITCurrentCardioSession").innerHTML =
      latestRec.HIITCurrentCardioSession;
    document.getElementById("coach-HIITCurrentCardioIntervals").innerHTML =
      latestRec.HIITCurrentCardioIntervals;
  }

  // === BASELINE CHART ===
  var baselineRec = userData.Recommendation[0];
  //Normal Day
  document.getElementById("NProteinAmount").innerHTML =
    baselineRec.NormalDayProtein;
  document.getElementById("NProteinRatio").innerHTML =
    (baselineRec.NormalDayProtein / baselineRec.NormalDayCalories).toPrecision(
      1
    ) + "%";
  document.getElementById("NCarbAmount").innerHTML = baselineRec.NormalDayCarb;
  document.getElementById("NCarbRatio").innerHTML =
    (baselineRec.NormalDayCarb / baselineRec.NormalDayCalories).toPrecision(1) +
    "%";
  document.getElementById("NFatAmount").innerHTML = baselineRec.NormalDayFat;
  document.getElementById("NFatRatio").innerHTML =
    (baselineRec.NormalDayFat / baselineRec.NormalDayCalories).toPrecision(1) +
    "%";
  document.getElementById("NCalories").innerHTML =
    baselineRec.NormalDayCalories;
  document.getElementById("NCaloriesRatio").innerHTML =
    (
      (baselineRec.NormalDayProtein +
        baselineRec.NormalDayCarb +
        baselineRec.NormalDayFat) /
      baselineRec.NormalDayCalories
    ).toPrecision(1) + "%";
  //High Day
  document.getElementById("HProteinAmount").innerHTML =
    baselineRec.HighDayProtein;
  document.getElementById("HProteinRatio").innerHTML =
    (baselineRec.HighDayProtein / baselineRec.HighDayCalories).toPrecision(1) +
    "%";
  document.getElementById("HCarbAmount").innerHTML = baselineRec.HighDayCarb;
  document.getElementById("HCarbRatio").innerHTML =
    (baselineRec.HighDayCarb / baselineRec.HighDayCalories).toPrecision(1) +
    "%";
  document.getElementById("HFatAmount").innerHTML = baselineRec.HighDayFat;
  document.getElementById("HFatRatio").innerHTML =
    (baselineRec.HighDayFat / baselineRec.HighDayCalories).toPrecision(1) + "%";
  document.getElementById("HCalories").innerHTML = baselineRec.HighDayCalories;
  document.getElementById("HCaloriesRatio").innerHTML =
    (
      (baselineRec.HighDayProtein +
        baselineRec.HighDayCarb +
        baselineRec.HighDayFat) /
      baselineRec.HighDayCalories
    ).toPrecision(1) + "%";
  //Low Day
  document.getElementById("LProteinAmount").innerHTML =
    baselineRec.LowDayProtein;
  document.getElementById("LProteinRatio").innerHTML =
    (baselineRec.LowDayProtein / baselineRec.LowDayCalories).toPrecision(1) +
    "%";
  document.getElementById("LCarbAmount").innerHTML = baselineRec.LowDayCarb;
  document.getElementById("LCarbRatio").innerHTML =
    (baselineRec.LowDayCarb / baselineRec.LowDayCalories).toPrecision(1) + "%";
  document.getElementById("LFatAmount").innerHTML = baselineRec.LowDayFat;
  document.getElementById("LFatRatio").innerHTML =
    (baselineRec.LowDayFat / baselineRec.LowDayCalories).toPrecision(1) + "%";
  document.getElementById("LCalories").innerHTML = baselineRec.LowDayCalories;
  document.getElementById("LCaloriesRatio").innerHTML =
    (
      (baselineRec.LowDayProtein +
        baselineRec.LowDayCarb +
        baselineRec.LowDayFat) /
      baselineRec.LowDayCalories
    ).toPrecision(1) + "%";
}

//==================================================
// Helper Functions

// returns true if user has no baseline recommendation (if user never completed profile)
function isNewUser(userData) {
  if (
    userData.Recommendation == null ||
    !userData.Recommendation[0].NormalDayProtein
  ) {
    return true;
  } else {
    return false;
  }
}

// returns latest recommendation object that has an "ModifiedDate"
function getLatestRecommendation(userData) {
  var latestRec = userData.Recommendation.filter(
    (value) => Object.keys(value).length !== 0
  ).slice(-1)[0];
  if (latestRec == null || latestRec.HIITCurrentCardioSession == null) {
    return null;
  } else if (latestRec.ModifiedDate) {
    return latestRec;
  } else {
    return null;
  }
}

// returns the int for the current week as shown in history page
function getCurrentWeek(startingDate) {
  return Math.floor(
    (new Date() - new Date(startingDate + "T00:00")) / 604800000
  );
}

// returns the int for the current day as shown in history page
function getCurrentDay(startingDate) {
  return Math.floor(
    ((new Date() - new Date(startingDate + "T00:00")) / 86400000) % 7
  );
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
  var waistCirc = document.getElementById("waistCirc");
  if (waistCirc) {
    formJSON.WaistCirc = Number(document.getElementById("waistCirc").value);
  }
  return formJSON;
}

//submit form as JSON on "save" button click
function submitForm() {
  //check that form is valid before sending
  if (document.getElementById("DailyUpdateForm").checkValidity() == false) {
    console.log("Form Invalid.");
    return;
  }
  //serialize form to JSON
  var dataObject = serializeDailyUpdate(
    document.getElementById("DailyUpdateForm")
  );
  var jsonData = JSON.stringify(dataObject);
  //POST JSON to api
  fetch(
    "http://localhost:8888/userDaily/" +
      getCurrentWeek(userData.StartDate) +
      "/" +
      getCurrentDay(userData.StartDate),
    {
      method: "post",
      body: jsonData,
    }
  ).then(function (response) {
    //show note that save was successful
    document.getElementById("SaveConfirmationText").innerHTML =
      "&nbsp; &nbsp; &nbsp; Saved!"; //TODO: add note for user too early/late, disable save button
  });
}
