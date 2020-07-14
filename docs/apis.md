# API Documentation 

# Internet Accessible API's
## Frontend v1
<!-- markdown-swagger -->
 Endpoint       | Method | Description
 -------------- | ------ | ----------------------------------------
 `/submitForm`  | POST   | Receives input from user
 `/getUserData` | GET    | Returns calculations and other user data
 `/apiVersion`  | GET    | Returns the symantec version number of the api
<!-- /markdown-swagger -->

# Backend API's
## Backend v1
<!-- markdown-swagger -->
 Endpoint          | Method | Description
 ----------------- | ------ | ----------------------------------------------
 `/userInfo/{UID}` | GET    | Get information related to user ID
 `/userInfo/{UID}` | POST   | Update information related to user ID
 `/apiVersion`     | GET    | Returns the symantec version number of the api
<!-- /markdown-swagger -->

## Auth v1
<!-- markdown-swagger -->
 Endpoint                 | Method | Description
 ------------------------ | ------ | -------------------------------------------------
 `/getUID/{SessionToken}` | GET    | Get the UID associated with a valid session token
 `/apiVersion`            | GET    | Returns the symantec version number of the api
<!-- /markdown-swagger -->
