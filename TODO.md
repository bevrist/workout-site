# TODO

- break frontend website into frontend-web, make frontend only a REST API
- verify all startup failures result in container termination 

## Database
- make errors not crash app
- update unit tests to marshal output to json instead of direct string compare

## auth
- find way to pass credentials using env vars

## database
- change reasonable post errors to non-fatal

## frontend
- refactor from structs.UserInfo to structs.Client
- modify frontend endpoints
    - consolidate submitUserProfile & getUserProfile
- write tests for frontend

## Webpage
- test new user sign up flow
- test new user redirect on baseline page
- test new user redirect only happens for new users
- test user profile update and saves
- test user baseline update on profile change
### weekly tracking
    - load starting week on page load

## Admin
- create admin page for admin users to view all other users?

## all
- add better error messages
    - add error ID to all error messages and ensure error is forwarded to user
- add prometheus logging?
- grafana monitoring
