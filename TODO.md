# TODO

- add ability to get list of all users to api
- Test with real firebase users
- complete FIXMES
- complete TODOS

- code coverage on tests
- remove "XMLHttpRequest" in javascript

- make daily update not function when outside active running time (message for being early or late)

- do data validation on incoming data
- double check all API docs
- finish frontend-api tests

- delete all references to Recommmendation[].Week in tests
- add gRPC connection type as alternative to REST
- modify all service Dockerfiles to use `scratch` instead of `alpine`
- add better error messages
    - add error ID to all error messages and ensure error is forwarded to user
- add prometheus logging?
- grafana dashboard monitoring


## auth
- remove debug for credentials env var

## database
- perform data validation on incoming data
    - validate StartDate and modifiedDate to be iso 8601
    - check all string inputs for day and recommend are lowercase

## Backend

## frontend-api
- finish unit tests
- update /GenerateUserBaseline tests
- do data validation on post requests (marshal to client object then send to other services)
    - test marshaling removes invalid extra data

## frontend-web

## Webpage
- verify new user sign up flow
- verify new user redirect on baseline page
- verify new user redirect only happens for new users
- verify user profile update and saves
- verify user baseline update on profile change

## Admin
- create admin page for admin users to view all other users' info
    - admin should be able to view list of all users and click on one button to open that users history in new window
    - click another button to get form for providing user a recommendation for this week
