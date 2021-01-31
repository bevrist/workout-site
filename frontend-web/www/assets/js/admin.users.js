// get user list from api and store JSON in "userList"
var xmlHttp = new XMLHttpRequest();
xmlHttp.open("GET", "http://localhost:8888/admin/listUsers", false);
xmlHttp.setRequestHeader("Session-Token",getCookie("Session-Token"));
// xmlHttp.setRequestHeader("Session-Token", myToken); //FIXME use correct session-token
xmlHttp.send(null);
var userList = JSON.parse(xmlHttp.responseText);

document.getElementById("user-list-table").innerHTML = "";

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
