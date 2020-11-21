//attempt redirect on page load
tryRedirect();

// create sign in with google button, handle sign in
var provider = new firebase.auth.GoogleAuthProvider();
document.getElementById("SignInWithGoogle").setAttribute("onclick", "signInWithGoogle()");
function signInWithGoogle() {
  firebase.auth().signInWithRedirect(provider);
}

firebase.auth().getRedirectResult().then(function (result) {
  // set "session-Token" cookie
  document.cookie = "Session-Token=" + result.user._lat + ";SameSite=Strict;path=/";

  tryRedirect();

}).catch(function (error) {
  // Handle Errors here.
  var errorCode = error.code;
  var errorMessage = error.message;
});

function tryRedirect() {
  if (getCookie("Session-Token") == null || getCookie("Session-Token") == "") {
    console.log("no cookie found...");
    return;
  }
  //check if user profile info is populated
  fetch("http://localhost:8888/userInfo", {
    headers: {
      'Session-Token': getCookie("Session-Token"),
    }
  }).then(function (response) {
    response.json().then(function (data) {
      //if first name is empty, redirect to profile page
      if (data.FirstName == "") {
        console.log("FirstName blank, redirecting to profile...")
        // window.location.replace('http://localhost:5500/profile');
      }
      else {
        //else redirect to daily-update page
        console.log("user profile data found, redirecting to daily-update...")
        // window.location.replace('http://localhost:5500/daily-update');
      }
    });
  });
  // console.log("redirecting to daily-update... I DONT KNOW WHY THIS SHOULDNT HAPPEN") //FIXME remove?
  // window.location.replace('http://localhost:5500/'); //FIXME remove?

  // //check if user profile info is populated
  // var xmlHttp = new XMLHttpRequest();
  // xmlHttp.open( "GET", "http://localhost:8888/userInfo", false );
  // xmlHttp.setRequestHeader("Session-Token",getCookie("Session-Token"));
  // // xmlHttp.setRequestHeader("Session-Token",myToken); //FIXME use correct session-token
  // xmlHttp.send(null);
  // var userData = JSON.parse(xmlHttp.responseText);
}

// enable auth signOut button
function authSignOut() {
  console.log("Signing out...")
  firebase.auth().signOut().then(function () {
    // Sign-out successful.
    // Invalidate "Session-Token" cookie
    document.cookie = "Session-Token=;expires=Thu, 01 Jan 1970 00:00:00 UTC;SameSite=Strict;path=/";
    console.log("Signed out button, redirecting to homepage...")
    // window.location.replace('http://localhost:5500/');
  }).catch(function (error) {
    console.log("Sign out error occurred")
    // An error happened.
  });
}
