//TODO: test flow for brand new users
// var myToken = "test"; //FIXME remove

//get user profile info from server and prepopulate form
var xmlHttp = new XMLHttpRequest();
xmlHttp.open( "GET", "http://localhost:8888/userInfo", false );
xmlHttp.setRequestHeader("Session-Token",getCookie("Session-Token"));
// xmlHttp.setRequestHeader("Session-Token",myToken); //FIXME use correct session-token
xmlHttp.send(null);
var userData = JSON.parse(xmlHttp.responseText);

// only update form if userData already exists
if (userData.FirstName != "") {
  updateForm(userData);
}

// //get UserProfile json object, call other functions on complete
// fetch("http://localhost:8888/userInfo", {
//   headers: {
//     'Session-Token': getCookie("Session-Token"),
//   }
// }).then(function (response) {
//   response.json().then(function (data) {
//     updateForm(data);
//   });
// });

// update form with existing user profile info (if present)
function updateForm(userProfileData) {
  // console.log("updateForm...");
  document.getElementById("firstName").value = userProfileData.FirstName;
  document.getElementById("lastName").value = userProfileData.LastName;
  var startDate = new Date(userProfileData.StartDate);
  document.getElementById("StartDate").value = startDate.toISOString().split('T')[0];
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
    console.log("Form Invalid.")
    return
  }
  //serialize form to JSON
  var dataObject = serializeProfile(document.getElementById("ProfileForm"));
  var jsonData = JSON.stringify(dataObject);
  //POST JSON to api
  var xmlHttp = new XMLHttpRequest();
  xmlHttp.open( "POST", "http://localhost:8888/userInfo", false );
  xmlHttp.setRequestHeader("Session-Token",getCookie("Session-Token"));
  // xmlHttp.setRequestHeader("Session-Token",myToken); //FIXME use correct session-token
  xmlHttp.send(jsonData);
  console.log("Server response: " + xmlHttp.responseText);
  //show note that save was successful
  document.getElementById("SaveConfirmationText").innerHTML = "&nbsp; &nbsp; &nbsp; Saved!";
  //TODO: update confirmation to link to daily-update page
}
