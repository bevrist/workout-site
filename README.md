# Workout-site

# Running Instructions:
> Run all commands from root folder of project  

## Docker Compose
to run with docker compose, simply run `docker-compose up --build`  


## Individual Containers
Auth: `docker build -t auth -f ./auth/Dockerfile .`  
`docker run -p 8070:8070 --rm -it -e AUTH_LISTEN_ADDRESS="0.0.0.0:8070" auth`

Backend: `docker build -t backend -f ./backend/Dockerfile .`  
`docker run -p 8090:8090 --rm -it -e BACKEND_LISTEN_ADDRESS="0.0.0.0:8090" backend`

Frontend: `docker build -t frontend -f ./frontend/Dockerfile .`  
`docker run -p 8080:8080 --rm -it -e FRONTEND_LISTEN_ADDRESS="0.0.0.0:8080" -e BACKEND_ADDRESS="localhost:8090" -e AUTH_ADDRESS="localhost:8070" -e FRONTEND_WEBSITE_URL="localhost:8080" frontend`

> you can use `host.docker.internal` in place of localhost to connect between containers on non linux platforms

# Tests:
Auth: `docker build -t auth-test -f ./auth/tests.Dockerfile .; docker run --rm -it auth-test`  

Frontend: `docker build -t frontend-test -f ./frontend/tests.Dockerfile .; docker run --rm -it frontend-test`  