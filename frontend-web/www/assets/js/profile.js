// add Session-Token to hidden form field
document.getElementById("Form-Session-Token").value = getCookie("Session-Token");

//get UserProfile json object, call other functions on complete
fetch("http://localhost:8888/userInfo", {
  headers: {
    'Session-Token': getCookie("Session-Token"),
  }
}).then(function (response) {
  response.json().then(function (data) {
    updateForm(data);
  });
});

// update form with existing user profile info (if present)
function updateForm(userProfileData) {
  console.log("updateForm...");
  document.getElementById("firstName").value = userProfileData.FirstName;
  document.getElementById("lastName").value = userProfileData.LastName;
  document.getElementById("weight").value = userProfileData.Weight;
  document.getElementById("heightInches").value = userProfileData.HeightInches;
  document.getElementById("waistCirc").value = userProfileData.WaistCirc;
  document.getElementById("leanBodyMass").value = userProfileData.LeanBodyMass;
  document.getElementById("age").value = userProfileData.Age;

  switch (userProfileData.Gender) {
    case "male":
      document.getElementById("genderMale").selected = "selected"
      break;
    case "female":
      document.getElementById("genderFemale").selected = "selected"
      break;
    default:
      document.getElementById("genderNone").selected = "selected"
  }
}





//Debug //FIXME remove
function hi(name) {
  fetch("http://localhost:8080/getUserProfile", {
    headers: {
      'Session-Token': name,
    }
  }).then(function (response) {
    response.json().then(function (data) {
      updateForm(data);
      console.log(data);
    });
  });
}