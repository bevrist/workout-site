// create sign in with google button, handle sign in
var provider = new firebase.auth.GoogleAuthProvider();
document.getElementById("SignInWithGoogle").setAttribute("onclick", "signInWithGoogle()");
function signInWithGoogle() {
  firebase.auth().signInWithRedirect(provider);
}
firebase.auth().getRedirectResult().then(function (result) {
  // set "session-Token" cookie
  document.cookie = "Session-Token=" + result.user._lat + "; SameSite=Strict;";

  //check if user profile info is populated
  fetch("http://localhost:8080/getUserProfile", {
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
  console.log("redirecting to homepage... I DONT KNOW WHY THIS SHOULDNT HAPPEN")
  // window.location.replace('http://localhost:5500/'); //FIXME remove?

}).catch(function (error) {
  // Handle Errors here.
  var errorCode = error.code;
  var errorMessage = error.message;
});

// enable signOut button
function authSignOut() {
  console.log("Signing out...")
  firebase.auth().signOut().then(function () {
    // Sign-out successful.
    // Invalidate "Session-Token" cookie
    document.cookie = "Session-Token=; expires=Thu, 01 Jan 1970 00:00:00 UTC;";
    console.log("Signed out button, redirecting to homepage...")
    // window.location.replace('http://localhost:5500/');
  }).catch(function (error) {
    console.log("Sign out error occurred")
    // An error happened.
  });
}
