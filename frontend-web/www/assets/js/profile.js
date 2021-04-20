document.title = "Profile";

var userData = "";

//get UserProfile json object from server, call other functions on complete
fetch("http://localhost:8888/userInfo").then(function (response) {
  if (response.status === 401) {
    //redirect to sign in on auth failure
    console.log("auth fail. redirecting...")
    window.location.href = "http://localhost:5500/auth";
  } else {
    response.json().then(function (userData) {
      // only update form if data already exists
      if (userData.FirstName != "") {
        document.title =
          "Profile - " + userData.FirstName + " " + userData.LastName;
        updateForm(userData);
      }
    });
  }
});

// update form with existing user profile info (if present)
function updateForm(userProfileData) {
  // console.log("updateForm...");
  document.getElementById("firstName").value = userProfileData.FirstName;
  document.getElementById("lastName").value = userProfileData.LastName;
  var startDate = new Date(userProfileData.StartDate);
  document.getElementById("StartDate").value = startDate
    .toISOString()
    .split("T")[0];
  document.getElementById("weight").value = userProfileData.Weight;
  document.getElementById("heightInches").value = userProfileData.HeightInches;
  document.getElementById("waistCirc").value = userProfileData.WaistCirc;
  document.getElementById("leanBodyMass").value = userProfileData.LeanBodyMass;
  document.getElementById("age").value = userProfileData.Age;

  switch (userProfileData.Gender) {
    case "male":
      document.getElementById("genderMale").selected = "selected";
      break;
    case "female":
      document.getElementById("genderFemale").selected = "selected";
      break;
    default:
      document.getElementById("genderNone").selected = "selected";
  }
}

//serialize form fields into json object
function serializeProfile(form) {
  return {
    FirstName: document.getElementById("firstName").value,
    LastName: document.getElementById("lastName").value,
    Weight: Number(document.getElementById("weight").value),
    WaistCirc: Number(document.getElementById("waistCirc").value),
    HeightInches: Number(document.getElementById("heightInches").value),
    LeanBodyMass: Number(document.getElementById("leanBodyMass").value),
    Age: Number(document.getElementById("age").value),
    StartDate: document.getElementById("StartDate").value,
    Gender: document.getElementById("gender").value,
  };
}

//submit form as JSON on "save" button click
function submitForm() {
  //check that form is valid before sending
  if (document.getElementById("ProfileForm").checkValidity() == false) {
    console.log("Form Invalid.");
    return;
  }
  //serialize form to JSON
  var dataObject = serializeProfile(document.getElementById("ProfileForm"));
  var jsonData = JSON.stringify(dataObject);
  //find endpoint to use based on if userBaseline data exists
  var apiEndpoint = "userInfo";
  if (isNewUser(userData) == true) {
    console.log("generating User Baseline...");
    apiEndpoint = "generateUserBaseline";
  }
  //POST JSON to api
  fetch("http://localhost:8888/" + apiEndpoint, {
    method: "post",
    body: jsonData,
  }).then(function (response) {
    console.log("Server response: " + response);
    document.getElementById("SaveConfirmationText").innerHTML =
      "Saved! Go to <a href=/daily-update>Daily-Update</a>";
  });
}

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
