#!/bin/bash

# exit when any command fails
set -e

# --- build all containers ---
# service containers
docker build -t auth:latest -f ./auth/Dockerfile .
docker build -t database:latest -f ./database/Dockerfile .
docker build -t backend:latest -f ./backend/Dockerfile .
docker build -t frontend-api:latest -f ./frontend-api/Dockerfile .
# testing containers
docker build -t auth-test:latest -f ./auth/tests.Dockerfile .
docker build -t database-test:latest -f ./database/tests.Dockerfile .
docker build -t mongodb-mock-database:latest -f ./database/mongoDB/Dockerfile .
# docker build -t backend-test:latest -f ./backend/tests.Dockerfile .
docker build -t frontend-api-test:latest -f ./frontend-api/tests.Dockerfile .


# --- run all tests ---
echo "running tests..."

# database test
echo "cleaning up hanging containers..."
docker stop mongodb-mock-database || :
docker stop database-service || :
echo "preparing database integration test..."
docker run --rm -d --name=mongodb-mock-database -p 27017:27017 --net=host -e MONGO_INITDB_ROOT_USERNAME=adminz -e MONGO_INITDB_ROOT_PASSWORD=cheeksbutt mongodb-mock-database:latest
sleep 2 && docker run -i --rm --name=database-service -p 8050:8050 --net=host database:latest &
echo "testing database..."
sleep 2 && docker run --rm -i --name=database-test --net=host database-test:latest
echo "cleaning up..."
docker stop mongodb-mock-database
docker stop database-service
echo ""

# auth test
echo "cleaning up hanging containers..."
docker stop auth-service || :
echo "preparing auth integration test..."
docker run -i --name=auth-service -p 8070:8070 --net=host --rm -e AUTH_FIREBASE_CREDENTIALS='{test}' auth:latest &
echo "testing auth..."
sleep 2 && docker run --rm -i --name=auth-test --net=host auth-test:latest
echo "cleaning up..."
docker stop auth-service
echo ""

# backend test
# TODO: complete backend test

# frontend-api test
echo "cleaning up hanging containers..."
docker stop auth-service || :
docker stop mongodb-mock-database || :
docker stop database-service || :
docker stop backend-service || :
echo "preparing frontend-api integration test..."
docker run -i --name=auth-service -p 8070:8070 --net=host --rm -e AUTH_FIREBASE_CREDENTIALS='{test}' auth:latest &
docker run --rm -d --name=mongodb-mock-database -p 27017:27017 --net=host -e MONGO_INITDB_ROOT_USERNAME=adminz -e MONGO_INITDB_ROOT_PASSWORD=cheeksbutt mongodb-mock-database:latest
sleep 2 && docker run -d --rm --name=database-service -p 8050:8050 --net=host database:latest
sleep 2 && docker run -d --rm --name=backend-service -p 8050:8050 --net=host backend:latest
sleep 2 && docker run -i --rm --name=frontend-api-service -p 8888:8888 --net=host frontend-api:latest &
echo "testing frontend-api..."
sleep 2 && docker run --rm -i --name=frontend-api-test --net=host frontend-api-test:latest
echo "cleaning up..."
docker stop auth-service
docker stop backend-service
docker stop database-service
docker stop mongodb-mock-database
echo ""

# frontend-web test
# TODO: complete frontend-web test
