# TODO
 
# AUTH REVAMP
- [ ] verify admin pages fail to load for non admin users
- [ ] fix website to use correct auth login/logout urls (google)

## REDIS
- [ ] make redis container
- [ ] add redis instructions to readme
- [ ] add redis data persistence


==========================================================
- complete FIXMES
- complete TODOS

- code coverage on tests

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



## database
- perform data validation on incoming data
    - validate StartDate and modifiedDate to be iso 8601
    - check all string inputs for day and recommend are lowercase

## frontend-api
- finish unit tests
- update /GenerateUserBaseline tests
- do data validation on post requests (marshal to client object then send to other services)
    - test marshaling removes invalid extra data

