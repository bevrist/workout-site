//FIXME update this
const queryString = window.location.search;
const urlParams = new URLSearchParams(queryString);
const userUid = urlParams.get("uid");
//if uid url parameter is empty redirect to /admin/users
if (userUid == "") {
  console.log("no uid URL parameter, redirecting to admin/users...");
  // window.location.replace('http://localhost:5500/admin/users');
}
//get user data from api and store JSON in "userData"
var xmlHttp = new XMLHttpRequest();
xmlHttp.open("GET", "http://localhost:8888/admin/userInfo", false);
xmlHttp.setRequestHeader("User-UID", userUid);
// xmlHttp.setRequestHeader("Session-Token",getCookie("Session-Token"));
xmlHttp.setRequestHeader("Session-Token", "testUID"); //FIXME use correct session-token
xmlHttp.send(null);
var userData = JSON.parse(xmlHttp.responseText);

updateForm(userData);
// update week-number on form to current week for user
function updateForm(userData) {
  //calculate current week num from user profile.startDate
  var startDate = new Date(userData.StartDate);//TODO: complete this
  var currentDate = new Date();
  weeksSinceStart = (Math.round((currentDate - startDate) / (7 * 24 * 60 * 60 * 1000)))+1;
  document.getElementById("week-number").value = weeksSinceStart;
}

//TODO:-------------------------------------------------------------------
//TODO:update below lines
//TODO:-------------------------------------------------------------------

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
  var xmlHttp = new XMLHttpRequest();
  xmlHttp.open("POST", "http://localhost:8888/" + apiEndpoint, false);
  xmlHttp.setRequestHeader("Session-Token", getCookie("Session-Token"));
  // xmlHttp.setRequestHeader("Session-Token",myToken); //FIXME use correct session-token
  xmlHttp.send(jsonData);
  console.log(jsonData);
  console.log("Server response: " + xmlHttp.responseText);
  //show note that save was successful
  document.getElementById("SaveConfirmationText").innerHTML =
    "Saved! Go to <a href=/daily-update>Daily-Update</a>";
}
