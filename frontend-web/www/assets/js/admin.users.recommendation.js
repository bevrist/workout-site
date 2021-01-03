const queryString = window.location.search;
const urlParams = new URLSearchParams(queryString);
const uid = urlParams.get('uid')
console.log(uid);

//TODO: if uid url parameter is empty redirect to /admin/users
