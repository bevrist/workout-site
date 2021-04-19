FROM golang:1

# ENV GO111MODULE=off
RUN wget -O /zoneinfo.zip https://github.com/golang/go/raw/master/lib/time/zoneinfo.zip
RUN go get -v github.com/google/go-cmp/cmp
RUN go get -v github.com/go-redis/redis/v8
RUN go get -v github.com/gorilla/mux github.com/markbates/goth github.com/markbates/goth/gothic
RUN go get -v github.com/markbates/goth/providers/google github.com/markbates/goth/providers/github
ENV ZONEINFO /zoneinfo.zip
WORKDIR /auth
COPY ./auth .
COPY ./common ../common

RUN printf '#!/bin/bash \n\
    go test -coverprofile=coverage.out \n\
    go tool cover -func=coverage.out \n\
    ' > /entrypoint.sh && chmod +x /entrypoint.sh
CMD ["/entrypoint.sh"]

ENV AUTH_SERVICE_ADDRESS="localhost:8070"
ENV AUTH_LISTEN_ADDRESS="localhost:8070"
ENV REDIS_CONNECTION_STRING="redis://localhost:6379/0"
ENV ADMINS='testUID,test3'
ENV TEST=1
