#!/bin/bash

# exit when any command fails
set -e

# --- build all container images ---
# service images
docker build -t auth:$(git describe --always) -f ./auth/Dockerfile .
docker build -t database:$(git describe --always) -f ./database/Dockerfile .
docker build -t backend:$(git describe --always) -f ./backend/Dockerfile .
docker build -t frontend-api:$(git describe --always) -f ./frontend-api/Dockerfile .
docker build -t frontend-web:$(git describe --always) ./frontend-web
# testing images
docker build -t auth-test:$(git describe --always) -f ./auth/tests.Dockerfile .
docker build -t database-test:$(git describe --always) -f ./database/tests.Dockerfile .
docker build -t mongodb-mock-database:$(git describe --always) -f ./database/mongoDB/Dockerfile .
docker build -t backend-test:$(git describe --always) -f ./backend/tests.Dockerfile .
docker build -t frontend-api-test:$(git describe --always) -f ./frontend-api/tests.Dockerfile .
docker build -t frontend-web-test:$(git describe --always) -f ./frontend-web/tests.Dockerfile .


# --- run all tests ---
echo "running tests..."

# database test
echo "cleaning up hanging containers..."
docker stop mongodb-mock-database || :
docker stop database-service || :
docker network rm database_net || :
sleep 1
echo "preparing database integration test..."
docker network create --driver bridge database_net
docker run --rm -d --name=mongodb-mock-database --net=database_net -e MONGO_INITDB_ROOT_USERNAME=adminz -e MONGO_INITDB_ROOT_PASSWORD=cheeksbutt mongodb-mock-database:$(git describe --always)
sleep 1 && docker run -i --rm --name=database-service --net=database_net -e DATABASE_ADDRESS=mongodb-mock-database:27017 -e DATABASE_LISTEN_ADDRESS=0.0.0.0:80 database:$(git describe --always) &
echo "testing database..."
sleep 1 && docker run --rm -i --name=database-test --net=database_net -e DATABASE_SERVICE_ADDRESS=database-service database-test:$(git describe --always)
echo "cleaning up..."
docker stop mongodb-mock-database
docker stop database-service
docker network rm database_net
echo ""

# auth test
echo "cleaning up hanging containers..."
docker stop auth-service || :
docker network rm auth_net || :
sleep 1
echo "preparing auth integration test..."
docker network create --driver bridge auth_net
docker run --rm -i --name=auth-service --net=auth_net -e TEST=1 -e AUTH_LISTEN_ADDRESS="0.0.0.0:80" auth:$(git describe --always) &
echo "testing auth..."
sleep 1 && docker run --rm -i --name=auth-test --net=auth_net -e AUTH_SERVICE_ADDRESS=auth-service auth-test:$(git describe --always)
echo "cleaning up..."
docker stop auth-service
docker network rm auth_net
echo ""

# backend test
echo "cleaning up hanging containers..."
docker stop mongodb-mock-database || :
docker stop database-service || :
docker stop backend-service || :
docker network rm backend_net ||:
sleep 1
echo "preparing backend integration test..."
docker network create --driver bridge backend_net
docker run --rm -d --name=mongodb-mock-database --net=backend_net -e MONGO_INITDB_ROOT_USERNAME=adminz -e MONGO_INITDB_ROOT_PASSWORD=cheeksbutt mongodb-mock-database:$(git describe --always)
sleep 1 && docker run -d --rm --name=database-service --net=backend_net -e DATABASE_ADDRESS=mongodb-mock-database:27017 -e DATABASE_LISTEN_ADDRESS=0.0.0.0:80 database:$(git describe --always)
sleep 1 && docker run -i --rm --name=backend-service --net=backend_net -e DATABASE_ADDRESS=database-service -e BACKEND_LISTEN_ADDRESS=0.0.0.0:80 backend:$(git describe --always) &
echo "testing backend..."
sleep 1 && docker run --rm -i --name=backend-test --net=backend_net -e BACKEND_API_SERVICE_ADDRESS=backend-service backend-test:$(git describe --always)
echo "cleaning up..."
docker stop backend-service
docker stop database-service
docker stop mongodb-mock-database
docker network rm backend_net
echo ""

# frontend-api test
echo "cleaning up hanging containers..."
docker stop auth-service || :
docker stop mongodb-mock-database || :
docker stop database-service || :
docker stop backend-service || :
docker stop frontend-api-service || :
docker network rm frontend-api_net ||:
sleep 1
echo "preparing frontend-api integration test..."
docker network create --driver bridge frontend-api_net
docker run --rm -d --name=mongodb-mock-database --net=frontend-api_net -e MONGO_INITDB_ROOT_USERNAME=adminz -e MONGO_INITDB_ROOT_PASSWORD=cheeksbutt mongodb-mock-database:$(git describe --always)
docker run --rm -d --name=auth-service --net=frontend-api_net -e TEST=1 -e AUTH_LISTEN_ADDRESS="0.0.0.0:80" auth:$(git describe --always)
sleep 1 && docker run -d --rm --name=database-service --net=frontend-api_net -e DATABASE_ADDRESS=mongodb-mock-database:27017 -e DATABASE_LISTEN_ADDRESS=0.0.0.0:80 database:$(git describe --always)
sleep 1 && docker run -d --rm --name=backend-service --net=frontend-api_net -e DATABASE_ADDRESS=database-service -e BACKEND_LISTEN_ADDRESS=0.0.0.0:80 backend:$(git describe --always)
sleep 1 && docker run -i --rm --name=frontend-api-service --net=frontend-api_net -e BACKEND_ADDRESS=backend-service -e AUTH_ADDRESS=auth-service -e FRONTEND_API_LISTEN_ADDRESS=0.0.0.0:80 frontend-api:$(git describe --always) &
echo "testing frontend-api..."
sleep 1 && docker run --rm -i --name=frontend-api-test --net=frontend-api_net -e FRONTEND_API_SERVICE_ADDRESS=frontend-api-service frontend-api-test:$(git describe --always)
echo "cleaning up..."
docker stop frontend-api-service
docker stop auth-service
docker stop backend-service
docker stop database-service
docker stop mongodb-mock-database
docker network rm frontend-api_net
echo ""

# frontend-web test
echo "cleaning up hanging containers..."
docker stop frontend-api-service || :
docker stop frontend-web-service || :
docker network rm frontend-web_net ||:
sleep 1
echo "preparing frontend-web integration test..."
docker network create --driver bridge frontend-web_net
sleep 1 && docker run -d --rm --name=frontend-api-service --net=frontend-web_net -e BACKEND_ADDRESS=backend-service -e AUTH_ADDRESS=auth-service -e FRONTEND_API_LISTEN_ADDRESS=0.0.0.0:80 frontend-api:$(git describe --always)
sleep 1 && docker run -d --rm --name=frontend-web-service --net=frontend-web_net -e FRONTEND_API_URL=http://frontend-api-service:8888 frontend-web:$(git describe --always)
echo "testing frontend-web..."
sleep 1 && docker run --rm -i --name=frontend-web-test --net=frontend-web_net -e FRONTEND_WEB_ADDRESS=http://frontend-web-service frontend-web-test:$(git describe --always)
echo "cleaning up..."
docker stop frontend-web-service
docker stop frontend-api-service
docker network rm frontend-web_net
echo ""

printf "\nALL TESTS PASSED!!! \n\n"