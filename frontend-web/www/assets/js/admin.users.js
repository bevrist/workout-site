// get user list from api and store JSON in "userList"
// var xmlHttp = new XMLHttpRequest();
// xmlHttp.open("GET", "http://localhost:8888/admin/listUsers", false);
// xmlHttp.setRequestHeader("Session-Token",getCookie("Session-Token"));
// // xmlHttp.setRequestHeader("Session-Token", myToken); //FIXME use correct session-token
// xmlHttp.send(null);
// var userList = JSON.parse(xmlHttp.responseText);

// TODO: show error on 403 error
fetch("http://localhost:8888/admin/listUsers").then(function (response) {
  if (response.status === 401) {
    //redirect to sign in on auth failure
    console.log("auth fail. redirecting...")
    window.location.href = "http://localhost:5500/auth";
  } else {
    response.json().then(function (userData) {
      updateForm(userData);
    });
  }
});

document.getElementById("user-list-table").innerHTML = "";
function updateForm(userList) {
  for (var i = 0; i < userList.length; i++) {
    userItemEntry =
      `<tr><td>` +
      userList[i].FirstName +
      ` ` +
      userList[i].LastName +
      `</td><td>` +
      userList[i].StartDate +
      `</td><td><a href="http://localhost:5500/admin/users/view?uid=` +
      userList[i].uid +
      `" target="_blank">View History</a></td><td><a href="http://localhost:5500/admin/users/recommendation?uid=` +
      userList[i].uid +
      `" target="_blank">New Rec</a></td></tr>`;
    document.getElementById("user-list-table").innerHTML += userItemEntry;
  }
}
