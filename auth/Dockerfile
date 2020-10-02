FROM golang:latest

RUN go get -v firebase.google.com/go firebase.google.com/go/auth github.com/gorilla/mux google.golang.org/api/option
WORKDIR /app
COPY ./common ../common
COPY ./auth/auth.go .
# RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o auth .
# uncomment this line to enable credential file for debug
# COPY ./auth/ .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /app .
CMD ["./auth"]

ENV AUTH_LISTEN_ADDRESS=0.0.0.0:8070
ENV AUTH_FIREBASE_CREDENTIALS='{}'
EXPOSE 8070
