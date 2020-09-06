#!/bin/bash

LOCAL_ADDRESS=host.docker.internal

# stop any running instances of mongodb-test container
docker stop $(docker ps | grep mongodb-test | cut -f 1 -d " ") &>/dev/null

# exit when any command fails
set -e

# --- build all containers ---
docker build -t mongodb-test -f ./database/mongoDB/Dockerfile . &
docker build -t database-test -f ./database/tests.Dockerfile . &
docker build -t auth-test -f ./auth/tests.Dockerfile . &
docker build -t frontend-test -f ./frontend/tests.Dockerfile . &
wait

# --- run all tests ---
echo running tests...
# database test
docker run --rm -d -e MONGO_INITDB_ROOT_USERNAME=adminz -e MONGO_INITDB_ROOT_PASSWORD=cheeksbutt -p 27017:27017 mongodb-test && sleep 1
docker run --rm -it -e DATABASE_ADDRESS="$LOCAL_ADDRESS":27017 database-test
docker stop $(docker ps | grep mongodb-test | cut -f 1 -d " ") &>/dev/null
# backend test
# TODO: complete backend test
# auth test
docker run --rm -it auth-test
# frontend test
docker run --rm -it frontend-test


# --- cleanup ---
# terminate mongoDB container
docker stop $(docker ps | grep mongodb-test | cut -f 1 -d " ") &>/dev/null
