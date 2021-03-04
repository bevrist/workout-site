FROM golang:1

ENV GO111MODULE=off
RUN go get -v firebase.google.com/go firebase.google.com/go/auth github.com/gorilla/mux google.golang.org/api/option
WORKDIR /auth
COPY ./auth .
COPY ./common ../common

RUN printf '#!/bin/bash \n\
    go test -coverprofile=coverage.out \n\
    go tool cover -func=coverage.out \n\
    ' > /entrypoint.sh && chmod +x /entrypoint.sh
CMD ["/entrypoint.sh"]

ENV AUTH_SERVICE_ADDRESS="localhost:8070"
