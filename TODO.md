# TODO

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
- add modifiedDate and Week parameters to "UpdateRecommendations"
- TESTS update names of backend-api actions tests

## frontend-api
- finish unit tests
- update /GenerateUserBaseline tests
- do data validation on post requests (marshal to client object then send to other services)
    - test marshaling removes invalid extra data

## frontend-web
- add a note showing what date is being edited on daily-update page
- correct history jump link day
- add WaistCirc to history form
- history page: 
    - give dropdowns a default empty option
    - make highlight on current week
    - highlight missed days in red
    - progress page shows weekly data and on weeks recommendation was made shows the recommendation below that week (like a new table inserted)
- Daily Update page:
    - make "daily update" page for adding that days data
- Profile page:
    - change profile weight/leanmass/waistCirc to "starting waistCirc..."

- Admin:
    - make 2 versions of weekly history page, edit version that has the current editable tables, and normal version that has simple tables (admin)
    - make a dedicated weekly recommendation creation page (for admin)
        - ??? coach adjustment: show baseline info, last coach recomendation, last user week, and form for new recomendation ???

- 3 pages:
    - profile page (for creating new profile)
    - daily update page (with user baseline and recommendation and form for entering daily data)
    - history page:  (showing all previous weeks with ability to issue corrections, and coach recommendations if present)


## Webpage
- test new user sign up flow
- test new user redirect on baseline page
- test new user redirect only happens for new users
- test user profile update and saves
- test user baseline update on profile change
### weekly tracking
    - load jump to starting week on page load


## Admin
- create admin page for admin users to view all other users' info?



# User Usage
User First Time Use:
enter profile info -> {app generates baseline data} -> enter first day daily data -> leave

User Daily Usage:
view baseline and updated recommendation -> enter daily data -> view/edit weekly data(history) | leave

Anthony Dad Usage:
view user history (looking at latest weekly data) -> update user weekly recommendation -> leave
