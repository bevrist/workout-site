# TODO


Fix generateUserBaseline to correctly generate a full recomendation object based off sample data `MISSING CALORIES`
- {"FirstName":"brett","LastName":"evrist","Weight":12,"WaistCirc":32,"HeightInches":12,"LeanBodyMass":21,"Age":23,"StartDate":"2020-11-21","Gender":"male"}

- code coverage
- remove "XMLHttpRequest" in javascript

- make daily update not function when outside active running time (message for being early or late)

- do data validation on incoming data
- double check all API docs
- finish frontend-api tests

- delete all references to Recommmendation[].Week in tests
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
- test new user sign up flow
- test new user redirect on baseline page
- test new user redirect only happens for new users
- test user profile update and saves
- test user baseline update on profile change
### weekly tracking
    - load jump to starting week on page load


## Admin
- create admin page for admin users to view all other users' info
    - admin should be able to view list of all users and click on one button to open that users history in new window
    - click another button to get form for providing user a recommendation for this week
