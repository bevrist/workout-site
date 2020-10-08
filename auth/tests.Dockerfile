FROM golang:1

RUN go get -v firebase.google.com/go firebase.google.com/go/auth github.com/gorilla/mux google.golang.org/api/option
WORKDIR /auth
COPY ./auth .
COPY ./common ../common
CMD ["go", "test"]

ENV AUTH_SERVICE_ADDRESS="localhost:8070"
