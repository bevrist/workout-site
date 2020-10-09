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
docker network rm database_net || :
sleep 5
echo "preparing database integration test..."
docker network create --driver bridge database_net || :
docker run --rm -d --name=mongodb-mock-database --net=database_net -e MONGO_INITDB_ROOT_USERNAME=adminz -e MONGO_INITDB_ROOT_PASSWORD=cheeksbutt mongodb-mock-database:$(git rev-parse --short HEAD)
sleep 5 && docker run -i --rm --name=database-service --net=database_net -e DATABASE_ADDRESS=mongodb-mock-database:27017 -e DATABASE_LISTEN_ADDRESS=0.0.0.0:80 database:$(git rev-parse --short HEAD) &
echo "testing database..."
sleep 5 && docker run --rm -i --name=database-test --net=database_net -e DATABASE_SERVICE_ADDRESS=database-service database-test:$(git rev-parse --short HEAD)
echo "cleaning up..."
docker stop mongodb-mock-database
docker stop database-service
docker network rm database_net
echo ""

# auth test
echo "cleaning up hanging containers..."
docker stop auth-service || :
docker network rm auth_net || :
sleep 5
echo "preparing auth integration test..."
docker network create --driver bridge auth_net || :
docker run --rm -i --name=auth-service --net=auth_net -e AUTH_FIREBASE_CREDENTIALS='{test}' -e AUTH_LISTEN_ADDRESS=0.0.0.0:80 auth:$(git rev-parse --short HEAD) &
echo "testing auth..."
sleep 5 && docker run --rm -i --name=auth-test --net=auth_net -e AUTH_SERVICE_ADDRESS=auth-service auth-test:$(git rev-parse --short HEAD)
echo "cleaning up..."
docker stop auth-service
docker network rm auth_net
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
docker network rm frontend-api_net ||:
sleep 5
echo "preparing frontend-api integration test..."
docker network create --driver bridge frontend-api_net || :
docker run --rm -d --name=auth-service --net=frontend-api_net -e AUTH_FIREBASE_CREDENTIALS='{test}' -e AUTH_LISTEN_ADDRESS=0.0.0.0:80 auth:$(git rev-parse --short HEAD)
docker run --rm -d --name=mongodb-mock-database --net=frontend-api_net -e MONGO_INITDB_ROOT_USERNAME=adminz -e MONGO_INITDB_ROOT_PASSWORD=cheeksbutt mongodb-mock-database:$(git rev-parse --short HEAD)
sleep 5 && docker run -d --rm --name=database-service --net=frontend-api_net -e DATABASE_ADDRESS=mongodb-mock-database:27017 -e DATABASE_LISTEN_ADDRESS=0.0.0.0:80 database:$(git rev-parse --short HEAD)
sleep 5 && docker run -d --rm --name=backend-service --net=frontend-api_net -e DATABASE_ADDRESS=database-service -e BACKEND_LISTEN_ADDRESS=0.0.0.0:80 backend:$(git rev-parse --short HEAD)
sleep 5 && docker run -i --rm --name=frontend-api-service --net=frontend-api_net -e BACKEND_ADDRESS=backend-service -e AUTH_ADDRESS=auth-service -e FRONTEND_API_LISTEN_ADDRESS=0.0.0.0:80 frontend-api:$(git rev-parse --short HEAD) &
echo "testing frontend-api..."
sleep 5 && docker run --rm -i --name=frontend-api-test --net=frontend-api_net -e FRONTEND_API_SERVICE_ADDRESS=frontend-api-service frontend-api-test:$(git rev-parse --short HEAD)
echo "cleaning up..."
docker stop frontend-api-service
docker stop auth-service
docker stop backend-service
docker stop database-service
docker stop mongodb-mock-database
docker network rm frontend-api_net
echo ""

# frontend-web test
# TODO: complete frontend-web test


printf "\nALL TESTS PASSED!!! \n\n"