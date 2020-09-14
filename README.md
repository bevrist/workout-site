# Workout-site

# Running Instructions:
> Run all commands from root folder of project  

## Docker Compose
to run with docker compose, simply run `docker-compose up --build`  


## Individual Containers
**Frontend:** `docker build -t frontend -f ./frontend/Dockerfile . && docker run -p 8080:8080 --rm -it -e FRONTEND_LISTEN_ADDRESS="0.0.0.0:8080" -e BACKEND_ADDRESS="localhost:8090" -e AUTH_ADDRESS="localhost:8070" -e FRONTEND_WEBSITE_URL="localhost:8080" frontend`

**Auth:** `docker build -t auth -f ./auth/Dockerfile . && docker run -p 8070:8070 --rm -it -e AUTH_LISTEN_ADDRESS="0.0.0.0:8070" auth`

**Backend:** `docker build -t backend -f ./backend/Dockerfile . && docker run -p 8090:8090 --rm -it -e BACKEND_LISTEN_ADDRESS="0.0.0.0:8090" backend`

**Database:** `docker build -t database -f ./database/Dockerfile . && docker run -p 8050:8050 --rm -it -e DATABASE_LISTEN_ADDRESS="0.0.0.0:8050" -e DATABASE_ADDRESS="localhost:27017" database`

> you can use `host.docker.internal` in place of **`localhost`** to connect between containers on non linux platforms

---

# Tests:
### Run all Integration Tests with `run_integration_tests.sh`  
---
## Individual Tests:
**mongoDB:** `docker build -t mongodb-test -f ./database/mongoDB/Dockerfile . && docker run --rm -it -e MONGO_INITDB_ROOT_USERNAME=adminz -e MONGO_INITDB_ROOT_PASSWORD=cheeksbutt -p 27017:27017 mongodb-test`

**Mongo-Express Debug Container:** `docker run -p 8081:8081 -e ME_CONFIG_MONGODB_ADMINUSERNAME=adminz -e ME_CONFIG_MONGODB_ADMINPASSWORD=cheeksbutt -e ME_CONFIG_MONGODB_SERVER=host.docker.internal mongo-express`

**Database:** `docker build -t database-test -f ./database/tests.Dockerfile . && docker run --rm -it -e DATABASE_ADDRESS=localhost:27017 database-test`  
> Depends on mongoDB

**Auth:** `docker build -t auth-test -f ./auth/tests.Dockerfile . && docker run --rm -it auth-test`  

**Frontend:** `docker build -t frontend-test -f ./frontend/tests.Dockerfile . && docker run --rm -it frontend-test`  
