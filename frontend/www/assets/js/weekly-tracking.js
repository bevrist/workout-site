// add Session-Token to hidden form field
document.getElementById("Form-Session-Token").value = getCookie("Session-Token");

// this function clones a unique form and renames items in form
function cloneForm(weekNum) {
    var myDiv = document.getElementById("myContainer");
    var divClone = myDiv.cloneNode(true);
    document.getElementById("myHeading").id = "Week" + weekNum;
    document.getElementById("Week" + weekNum).innerHTML = "Week " + weekNum;
    document.getElementById("firstName").id = "aaaaaName" + weekNum;
    document.getElementById("lastName").id = "bbbbbName" + weekNum;
    document.getElementById("myForm").id = "form" + weekNum;
    document.getElementById("form" + weekNum).action= "http://localhost:8080/userWeeklyTracking/" + weekNum;
    document.getElementById("myContainer").id = "week" + weekNum;
    document.getElementById("formLink").id= "formLink" + weekNum;
    document.getElementById("formLink" + weekNum).name= "asdf" + weekNum;
    document.body.appendChild(divClone);
}

cloneForm(1);
cloneForm(2);
cloneForm(3);
cloneForm(4);
cloneForm(5);
cloneForm(6);
cloneForm(7);
cloneForm(8);
cloneForm(9);
cloneForm(10);
cloneForm(11);
cloneForm(12);
cloneForm(13);
cloneForm(14);
cloneForm(15);
cloneForm(16);
cloneForm(17);
cloneForm(18);
cloneForm(19);
cloneForm(20);
document.getElementById("myContainer").remove(); // delete empty clone form at end

// jump to particular location on page //TODO: calculate current week and redirect here
location.href = "#asdf9";


//get UserProfile json object, call other functions on complete
fetch("http://localhost:8080/userWeeklyTracking/" + num, {
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

  if (userProfileData.Gender == "male") {
    document.getElementById("genderMale").selected = "selected"
  }
  else if (userProfileData.Gender == "female") {
    document.getElementById("genderFemale").selected = "selected"
  }
  else {
    document.getElementById("genderNone").selected = "selected"
  }
}