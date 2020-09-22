#!/bin/bash

# exit when any command fails
set -e

# --- build all containers ---
# service containers
docker build -t auth:latest -f ./auth/Dockerfile . &
docker build -t database:latest -f ./database/Dockerfile . &
wait
# testing containers
docker build -t auth-test:latest -f ./auth/tests.Dockerfile . &
docker build -t database-test:latest -f ./database/tests.Dockerfile . &
docker build -t mongodb-mock-database:latest -f ./database/mongoDB/Dockerfile . &
docker build -t frontend-test:latest -f ./frontend/tests.Dockerfile . &
wait

# --- run all tests ---
echo "running tests..."
# database test
echo "preparing database integration test..."
docker stop mongodb-mock-database || :
docker stop database-service || :
docker run --rm -d --name=mongodb-mock-database -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=adminz -e MONGO_INITDB_ROOT_PASSWORD=cheeksbutt mongodb-mock-database:latest
sleep 2 && docker run --rm -d --name=database-service -p 8050:8050 --net=host database:latest
echo "testing database..."
sleep 5 && docker run --rm -i --name=database-test --net=host database-test:latest
docker stop mongodb-mock-database
docker stop database-service
# backend test
# TODO: complete backend test
# auth test
echo "preparing auth integration test..."
# stop hanging auth-service instances
docker stop auth-service || :
docker run -d --name=auth-service -p 8070:8070 --rm -e AUTH_LISTEN_ADDRESS="0.0.0.0:8070" -e AUTH_FIREBASE_CREDENTIALS='{test}' auth:latest
echo "testing auth..."
sleep 5 && docker run --rm -i --name=auth-test --net=host auth-test:latest
docker stop auth-service
# frontend test
# echo "testing frontend..."
# FIXME get frontend tests working again
# docker run --rm -i --net=host frontend-test
# frontend-web test
# TODO: complete frontend-web test
