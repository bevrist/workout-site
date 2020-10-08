FROM golang:1
RUN go get -v github.com/gorilla/mux go.mongodb.org/mongo-driver/bson go.mongodb.org/mongo-driver/mongo go.mongodb.org/mongo-driver/mongo/options go.mongodb.org/mongo-driver/mongo/readpref
RUN go get -v github.com/google/go-cmp/cmp
WORKDIR /database
COPY ./common ../common
COPY ./database .
CMD ["go", "test"]

ENV DATABASE_SERVICE_ADDRESS="localhost:8050"
