# TODO


- delete all references to Recommmendation[].Week in tests
- modify all service Dockerfiles to use `scratch` instead of `alpine`
- add better error messages
    - add error ID to all error messages and ensure error is forwarded to user
- add prometheus logging?
- grafana dashboard monitoring


## auth


## database


## Backend
- add modifiedDate and Week parameters to "UpdateRecommendations"

## frontend-api
- finish unit tests
- update /GenerateUserBaseline tests
- do data validation on post requests

## frontend-web
- put start date in profile page
- make "daily update" page
- put recommendation and baseline on "daily update" page
- change profile weight/leanmass/waistCirc to "starting waistCirc..."
- change baseline page to "Recommendation page" 
- make "daily data" page for entering day data
- history page shows weekly data and on weeks  recommendation was made shows the recomendation below that week
- make a dedicated weekly recommendation page (for admin)
- make baseline page show current recommendation and "starting baseline" (recommendation shows "last edited" date)
- coach adjustment: show baseline info, last coach recomendation, last user week, and form for new recomendation
- make 2 versions of weekly history page, edit version that has the current editable tables, and normal version that has simple tables
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



# User Usage
User First Time Use:
enter profile info -> {app generates baseline data} -> enter first day daily data -> leave

User Daily Usage:
view baseline and updated recommendation -> enter daily data -> view/edit weekly data(history) | leave

Anthony Dad Usage:
view user history (looking at latest weekly data) -> update user weekly recommendation -> leave
