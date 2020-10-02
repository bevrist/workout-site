# Workout-site

# Running Instructions:
> Run all commands from root folder of project  

## Docker Compose
to run with docker compose, simply run `docker-compose up --build`  


## Individual Containers
**Frontend-api:** `docker build -t frontend-api -f ./frontend-api/Dockerfile . && docker run -p 8080:8080 --net=host --rm -it -e FRONTEND_API_LISTEN_ADDRESS="0.0.0.0:8888" -e BACKEND_ADDRESS="localhost:8090" -e AUTH_ADDRESS="localhost:8070" frontend-api`

**Auth:** `docker build -t auth -f ./auth/Dockerfile . && docker run -p 8070:8070 --net=host --rm -it -e AUTH_LISTEN_ADDRESS="0.0.0.0:8070" -e AUTH_FIREBASE_CREDENTIALS='{}' auth`

**Backend:** `docker build -t backend -f ./backend/Dockerfile . && docker run -p 8090:8090 --net=host --rm -it -e BACKEND_LISTEN_ADDRESS="0.0.0.0:8090" -e DATABASE_ADDRESS="localhost:8050" backend`

**Database:** `docker build -t database -f ./database/Dockerfile . && docker run -p 8050:8050 --net=host --rm -it -e DATABASE_LISTEN_ADDRESS="0.0.0.0:8050" -e DATABASE_ADDRESS="localhost:27017" -e DATABASE_USERNAME="adminz" -e DATABASE_PASSWORD="cheeksbutt" database`

---
**mongoDB:** `docker build -t mongodb-test -f ./database/mongoDB/Dockerfile . && docker run --rm -it -e MONGO_INITDB_ROOT_USERNAME=adminz -e MONGO_INITDB_ROOT_PASSWORD=cheeksbutt -p 27017:27017 --net=host mongodb-test`
---

# Tests:
### Run all Integration Tests with `run_integration_tests.sh`  
---
## Individual Tests:
**Mongo-Express Debug Container:** `docker run -p 8081:8081 -e ME_CONFIG_MONGODB_ADMINUSERNAME=adminz -e ME_CONFIG_MONGODB_ADMINPASSWORD=cheeksbutt -e ME_CONFIG_MONGODB_SERVER=localhost mongo-express`

**Database:** `docker build -t database-test -f ./database/tests.Dockerfile . && docker run --rm -it --net=host -e DATABASE_ADDRESS=localhost:27017 database-test`  
> Depends on mongoDB

**Auth:** `docker build -t auth-test -f ./auth/tests.Dockerfile . && docker run --rm -it --net=host auth-test`  

**Frontend:** `docker build -t frontend-test -f ./frontend/tests.Dockerfile . && docker run --rm -it --net=host frontend-test`  
