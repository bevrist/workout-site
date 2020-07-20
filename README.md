# Workout-site

# Running Instructions:
> Run all commands from root folder of project 

Auth: `docker build -t auth -f ./auth/Dockerfile .`  
`docker run -p 8070:8070 --rm -it -e AUTH_LISTEN_ADDRESS="0.0.0.0:8070" auth`

Backend: `docker build -t backend -f ./backend/Dockerfile .`  
`docker run -p 8090:8090 --rm -it -e BACKEND_LISTEN_ADDRESS="0.0.0.0:8090" backend`

Frontend: `docker build -t frontend -f ./frontend/Dockerfile .`  
`docker run -p 8080:8080 --rm -it -e FRONTEND_LISTEN_ADDRESS="0.0.0.0:8080" -e BACKEND_ADDRESS="localhost:8090" -e AUTH_ADDRESS="localhost:8070" frontend`
