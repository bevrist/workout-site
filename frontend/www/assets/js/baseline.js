//get UserBaseline json object, call other functions on complete
fetch("http://localhost:8080/getUserBaseline", {
  headers: {
    'Session-Token': getCookie("Session-Token"),
  }
}).then(function (response) {
  response.json().then(function (data) {
    updateCharts(data);
  });
});

// updates baseline charts with user data
function updateCharts(userBaselineData) {
  // redirect to profile page on empty data
  if (userBaselineData.NormalDay == null) {
    console.log("data blank, redirecting to profile...")
    // window.location.replace('http://localhost:5500/profile');  //FIXME
  }
  //Normal Day
  document.getElementById("NProteinAmount").innerHTML = userBaselineData.NProteinAmount;
  document.getElementById("NProteinRatio").innerHTML = userBaselineData.NProteinRatio;
  document.getElementById("NCarbAmount").innerHTML = userBaselineData.NCarbAmount;
  document.getElementById("NCarbRatio").innerHTML = userBaselineData.NCarbRatio;
  document.getElementById("NFatAmount").innerHTML = userBaselineData.NFatAmount;
  document.getElementById("NFatRatio").innerHTML = userBaselineData.NFatRatio;
  document.getElementById("NormalDay").innerHTML = userBaselineData.NormalDay;
  //High Day
  document.getElementById("HProteinAmount").innerHTML = userBaselineData.HProteinAmount;
  document.getElementById("HProteinRatio").innerHTML = userBaselineData.HProteinRatio;
  document.getElementById("HCarbAmount").innerHTML = userBaselineData.HCarbAmount;
  document.getElementById("HCarbRatio").innerHTML = userBaselineData.HCarbRatio;
  document.getElementById("HFatAmount").innerHTML = userBaselineData.HFatAmount;
  document.getElementById("HFatRatio").innerHTML = userBaselineData.HFatRatio;
  document.getElementById("HighDay").innerHTML = userBaselineData.HighDay;
  //Low Day
  document.getElementById("LProteinAmount").innerHTML = userBaselineData.LProteinAmount;
  document.getElementById("LProteinRatio").innerHTML = userBaselineData.LProteinRatio;
  document.getElementById("LCarbAmount").innerHTML = userBaselineData.LCarbAmount;
  document.getElementById("LCarbRatio").innerHTML = userBaselineData.LCarbRatio;
  document.getElementById("LFatAmount").innerHTML = userBaselineData.LFatAmount;
  document.getElementById("LFatRatio").innerHTML = userBaselineData.LFatRatio;
  document.getElementById("LowDay").innerHTML = userBaselineData.LowDay;
}





//FIXME remove debug
function hi(name) {
  //get UserBaseline json object, call other functions on complete
  fetch("http://localhost:8080/getUserBaseline", {
    headers: {
      'Session-Token': name,
    }
  }).then(function (response) {
    response.json().then(function (data) {
      updateCharts(data);
    });
  });
}