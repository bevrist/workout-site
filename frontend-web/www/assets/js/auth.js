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
        // window.location.replace('http://localhost:5500/profile');  //FIXME
      }
      else {
        //else redirect to baseline page
        console.log("user profile data found, redirecting to baseline...")
        // window.location.replace('http://localhost:5500/baseline'); //FIXME
      }
    });
  });
  console.log("redirecting to baseline... I DONT KNOW WHY THIS SHOULDNT HAPPEN")
  // window.location.replace('http://localhost:5500/baseline'); //FIXME remove?

}).catch(function (error) {
  // Handle Errors here.
  var errorCode = error.code;
  var errorMessage = error.message;
});
