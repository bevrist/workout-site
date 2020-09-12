FROM golang:latest
# this test requires mongoDB container to be running
RUN go get -v github.com/gorilla/mux go.mongodb.org/mongo-driver/bson go.mongodb.org/mongo-driver/mongo go.mongodb.org/mongo-driver/mongo/options go.mongodb.org/mongo-driver/mongo/readpref
RUN go get -v github.com/google/go-cmp/cmp
WORKDIR /database
COPY ./database .
COPY ./common ../common
# RUN go get -d -v ./...
RUN echo "go run ./database.go & sleep 1; go test" > start.sh; chmod 777 start.sh
CMD ["/bin/bash", "./start.sh"]

ENV DATABASE_LISTEN_ADDRESS=0.0.0.0:80
ENV DATABASE_ADDRESS=host.docker.internal:27017
