FROM golang:latest

RUN go get -v firebase.google.com/go firebase.google.com/go/auth github.com/gorilla/mux google.golang.org/api/option
WORKDIR /auth
COPY ./auth .
COPY ./common ../common
# RUN go get -d -v ./...
RUN echo "go run ./auth.go & sleep 1; go test" > start.sh; chmod 777 start.sh
CMD ["/bin/bash", "./start.sh"]

EXPOSE 8070
