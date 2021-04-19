# Workout-site

This project is a monorepo of all of the services which power this site. 
A workout tracking application for helping users track optimal intake and workout sessions when working with a trainer.


# Running Instructions:
> Run all commands from root folder of project  

## Docker Compose
to run a simple local instance, run: `docker-compose down && docker-compose up --build`  


## Individual Containers
**Frontend-Web:** `docker build -t frontend-web ./frontend-web && docker run --rm -it -e FRONTEND_API_URL="http://localhost:8888" -p 80:80 frontend-web`

**Frontend-api:** `docker build -t frontend-api -f ./frontend-api/Dockerfile . && docker run -p 8080:8080 --net=host --rm -it -e FRONTEND_API_LISTEN_ADDRESS="0.0.0.0:8888" -e BACKEND_ADDRESS="localhost:8090" -e AUTH_ADDRESS="localhost:8070" frontend-api`

**Auth:** `docker build -t auth -f ./auth/Dockerfile . && docker run -p 8070:8070 --net=host --rm -it -e AUTH_LISTEN_ADDRESS="0.0.0.0:8070" -e REDIS_CONNECTION_STRING="redis://localhost:6379/0" -e PROVIDER_SECRET="" -e PROVIDER_KEY="" -e ADMINS='testUID,test3' auth`

**Backend:** `docker build -t backend -f ./backend/Dockerfile . && docker run -p 8090:8090 --net=host --rm -it -e BACKEND_LISTEN_ADDRESS="0.0.0.0:8090" -e DATABASE_ADDRESS="localhost:8050" backend`

**Database:** `docker build -t database -f ./database/Dockerfile . && docker run -p 8050:8050 --net=host --rm -it -e DATABASE_LISTEN_ADDRESS="0.0.0.0:8050" -e DATABASE_ADDRESS="localhost:27017" -e DATABASE_USERNAME="adminz" -e DATABASE_PASSWORD="cheeksbutt" database`

---
**mongoDB:** `docker build -t mongodb-test -f ./database/mongoDB/Dockerfile . && docker run --rm -it -e MONGO_INITDB_ROOT_USERNAME=adminz -e MONGO_INITDB_ROOT_PASSWORD=cheeksbutt -p 27017:27017 --net=host mongodb-test`

**Redis:** `docker build -t redis-test auth/redis && docker run --rm -it -p 6379:6379 redis-test redis-server`
---

# Tests:
### Run all Integration Tests with `run_integration_tests.sh`  

---
## Individual Tests:
> All tests expect the corresponding service to be running

**Mongo-Express Debug Container:** `docker run -p 8081:8081 -e ME_CONFIG_MONGODB_ADMINUSERNAME=adminz -e ME_CONFIG_MONGODB_ADMINPASSWORD=cheeksbutt -e ME_CONFIG_MONGODB_SERVER=localhost mongo-express`

**Database:** `docker build -t database-test -f ./database/tests.Dockerfile . && docker run --rm -it --net=host -e DATABASE_ADDRESS=localhost:27017 database-test`  
> Depends on mongoDB

**Auth:** `docker build -t auth-test -f ./auth/tests.Dockerfile . && docker run --rm -it --net=host auth-test`  

**Backend:** `docker build -t backend-test -f ./backend/tests.Dockerfile . && docker run --rm -it --net=host backend-test`

**Frontend-api:** `docker build -t frontend-api-test -f ./frontend-api/tests.Dockerfile . && docker run --rm -it --net=host frontend-api-test`  

**Frontend-web:** `docker build -t frontend-web-test -f ./frontend-web/tests.Dockerfile . && docker run --rm -it --net=host frontend-web-test` 
