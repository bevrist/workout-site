const queryString = window.location.search;
const urlParams = new URLSearchParams(queryString);
const userUid = urlParams.get("uid");
//if uid url parameter is empty redirect to /admin/users
if (userUid == "") {
  console.log("no uid URL parameter, redirecting to admin/users...");
  // window.location.replace('http://localhost:5500/admin/users');
}

fetch("http://localhost:8888/admin/userInfo", { headers: { "User-UID": userUid }}).then(function (response) {
  if (response.status === 401) {
    //redirect to sign in on auth failure
    console.log("auth fail. redirecting...")
    window.location.href = "http://localhost:5500/auth";
  } else {
    response.json().then(function (uData) {
      updateForm(uData);
    });
  }
});

// update week-number on form to current week for user
function updateForm(userData) {
  //calculate current week from user profile.startDate
  var startDate = new Date(userData.StartDate);
  var currentDate = new Date();
  weeksSinceStart = (Math.round((currentDate - startDate) / (7 * 24 * 60 * 60 * 1000)))+1;
  document.getElementById("week-number").value = weeksSinceStart;
}

//serialize form fields into json object
function serializeRecForm(form) {
  return {
    WeekNumber: Number(document.getElementById("week-number").value),
    HIITCurrentCardioSession: Number(document.getElementById("hiit-session").value),
    HIITCurrentCardioIntervals: Number(document.getElementById("hiit-intervals").value),
    HighDayProtein: Number(document.getElementById("h-protein").value),
    HighDayCarb: Number(document.getElementById("h-carbs").value),
    HighDayFat: Number(document.getElementById("h-fat").value),
    HighDayCalories: Number(document.getElementById("h-calories").value),
    NormalDayProtein: Number(document.getElementById("m-protein").value),
    NormalDayCarb: Number(document.getElementById("m-carbs").value),
    NormalDayFat: Number(document.getElementById("m-fat").value),
    NormalDayCalories: Number(document.getElementById("m-calories").value),
    LowDayProtein: Number(document.getElementById("l-protein").value),
    LowDayCarb: Number(document.getElementById("l-carbs").value),
    LowDayFat: Number(document.getElementById("l-fat").value),
    LowDayCalories: Number(document.getElementById("l-calories").value),

    ModifiedDate: new Date(), //this probably needs to be iso 8601 format?
  };
}

//submit form as JSON on "save" button click
function submitForm() {
  //check that form is valid before sending
  if (document.getElementById("RecForm").checkValidity() == false) {
    console.log("Form Invalid.");
    return;
  }
  //serialize form to JSON
  var dataObject = serializeRecForm(document.getElementById("RecForm"));
  var jsonData = JSON.stringify(dataObject);
  //POST JSON to api
  fetch("http://localhost:8888/admin/userRecommendation/" + (dataObject.WeekNumber - 1), {
    headers: { "User-UID": userUid },
    method: "post",
    body: jsonData,
  }).then(function (response) {
    console.log("Server response: " + response);
    document.getElementById("SaveConfirmationText").innerHTML = "Saved!";
  });
}
