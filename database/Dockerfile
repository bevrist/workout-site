FROM golang:latest
WORKDIR /app
RUN go get -v github.com/gorilla/mux go.mongodb.org/mongo-driver/bson go.mongodb.org/mongo-driver/mongo go.mongodb.org/mongo-driver/mongo/options go.mongodb.org/mongo-driver/mongo/readpref
COPY ./database .
COPY ./common ../common
# RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o database .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /app .
CMD ["./database"]

EXPOSE 8050

ENV DATABASE_LISTEN_ADDRESS=0.0.0.0:8050
ENV DATABASE_ADDRESS=localhost:27017
ENV DATABASE_USERNAME=adminz
ENV DATABASE_PASSWORD=cheeksbutt
