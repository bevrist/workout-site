# Workout-site

# Running Instructions:
> Run all commands from root folder of project 

Auth: `docker build -t auth -f ./auth/Dockerfile .`  
`docker run -p 8070:8070 --rm -it auth`

Backend: ``

Frontend: `docker build -t frontend -f ./frontend/Dockerfile .`  
`docker run -p 8080:8080 --rm -it frontend`

<!-- TODO: add env flags for frontend and backend -->