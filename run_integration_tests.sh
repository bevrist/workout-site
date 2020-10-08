#!/bin/bash

# exit when any command fails
set -e

# --- build all container images ---
# service images
docker build -t auth:$(git rev-parse --short HEAD) -f ./auth/Dockerfile .
docker build -t database:$(git rev-parse --short HEAD) -f ./database/Dockerfile .
docker build -t backend:$(git rev-parse --short HEAD) -f ./backend/Dockerfile .
docker build -t frontend-api:$(git rev-parse --short HEAD) -f ./frontend-api/Dockerfile .
# testing images
docker build -t auth-test:$(git rev-parse --short HEAD) -f ./auth/tests.Dockerfile .
docker build -t database-test:$(git rev-parse --short HEAD) -f ./database/tests.Dockerfile .
docker build -t mongodb-mock-database:$(git rev-parse --short HEAD) -f ./database/mongoDB/Dockerfile .
# docker build -t backend-test:$(git rev-parse --short HEAD) -f ./backend/tests.Dockerfile .
docker build -t frontend-api-test:$(git rev-parse --short HEAD) -f ./frontend-api/tests.Dockerfile .


# --- run all tests ---
echo "running tests..."

# database test
echo "cleaning up hanging containers..."
docker stop mongodb-mock-database || :
docker stop database-service || :
sleep 5
echo "preparing database integration test..."
docker run --rm -d --name=mongodb-mock-database -p 27017:27017 --net=host -e MONGO_INITDB_ROOT_USERNAME=adminz -e MONGO_INITDB_ROOT_PASSWORD=cheeksbutt mongodb-mock-database:$(git rev-parse --short HEAD)
sleep 15 && docker run -i --rm --name=database-service -p 8050:8050 --net=host database:$(git rev-parse --short HEAD) &
echo "testing database..."
sleep 5 && docker run --rm -i --name=database-test --net=host database-test:$(git rev-parse --short HEAD)
echo "cleaning up..."
docker stop mongodb-mock-database
docker stop database-service
echo ""

# auth test
echo "cleaning up hanging containers..."
docker stop auth-service || :
sleep 5
echo "preparing auth integration test..."
docker run -i --name=auth-service -p 8070:8070 --net=host --rm -e AUTH_FIREBASE_CREDENTIALS='{test}' auth:$(git rev-parse --short HEAD) &
echo "testing auth..."
sleep 5 && docker run --rm -i --name=auth-test --net=host auth-test:$(git rev-parse --short HEAD)
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
docker stop frontend-api-service || :
sleep 5
echo "preparing frontend-api integration test..."
docker run -i --name=auth-service -p 8070:8070 --net=host --rm -e AUTH_FIREBASE_CREDENTIALS='{test}' auth:$(git rev-parse --short HEAD) &
docker run --rm -d --name=mongodb-mock-database -p 27017:27017 --net=host -e MONGO_INITDB_ROOT_USERNAME=adminz -e MONGO_INITDB_ROOT_PASSWORD=cheeksbutt mongodb-mock-database:$(git rev-parse --short HEAD)
sleep 5 && docker run -d --rm --name=database-service -p 8050:8050 --net=host database:$(git rev-parse --short HEAD)
sleep 5 && docker run -d --rm --name=backend-service -p 8050:8050 --net=host backend:$(git rev-parse --short HEAD)
sleep 5 && docker run -i --rm --name=frontend-api-service -p 8888:8888 --net=host frontend-api:$(git rev-parse --short HEAD) &
echo "testing frontend-api..."
sleep 5 && docker run --rm -i --name=frontend-api-test --net=host frontend-api-test:$(git rev-parse --short HEAD)
echo "cleaning up..."
docker stop frontend-api-service
docker stop auth-service
docker stop backend-service
docker stop database-service
docker stop mongodb-mock-database
echo ""

# frontend-web test
# TODO: complete frontend-web test







# # frontend-api test
# echo "cleaning up hanging containers..."
# docker stop auth-service || :
# docker stop mongodb-mock-database || :
# docker stop database-service || :
# docker stop backend-service || :
# docker stop frontend-api-service || :
# sleep 5
# echo "preparing frontend-api integration test..."
# docker run -i --name=auth-service -p 8070:8070 --net=host --rm -e AUTH_FIREBASE_CREDENTIALS='{test}' auth:$(git rev-parse --short HEAD) &
# docker run --rm -d --name=mongodb-mock-database -p 27017:27017 --net=host -e MONGO_INITDB_ROOT_USERNAME=adminz -e MONGO_INITDB_ROOT_PASSWORD=cheeksbutt mongodb-mock-database:$(git rev-parse --short HEAD)
# sleep 5 && docker run -i --rm --name=database-service -p 8050:8050 --net=host database:$(git rev-parse --short HEAD) &
# sleep 5 && docker run -i --rm --name=backend-service -p 8050:8050 --net=host backend:$(git rev-parse --short HEAD) &
# sleep 5 && docker run -i --rm --name=frontend-api-service -p 8888:8888 --net=host frontend-api:$(git rev-parse --short HEAD) &
# echo "testing frontend-api..."
# sleep 5 && docker run --rm -i --name=frontend-api-test --net=host frontend-api-test:$(git rev-parse --short HEAD)
# echo "cleaning up..."
# docker stop frontend-api-service
# docker stop auth-service
# docker stop backend-service
# docker stop database-service
# docker stop mongodb-mock-database
# echo ""




printf "\nALL TESTS PASSED!!! \n\n"