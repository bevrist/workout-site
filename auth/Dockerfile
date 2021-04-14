FROM golang:1

# ENV GO111MODULE=off
RUN go get -v github.com/go-redis/redis/v8
RUN go get -v github.com/gorilla/mux github.com/gorilla/pat github.com/markbates/goth github.com/markbates/goth/gothic
RUN go get -v github.com/markbates/goth/providers/google github.com/markbates/goth/providers/github
WORKDIR /app
COPY ./common ../common
COPY ./auth/ .
# RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o auth .


FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /app .
CMD ["./auth"]

ENV AUTH_LISTEN_ADDRESS=0.0.0.0:8070
# ENV AUTH_FIREBASE_CREDENTIALS='{}'
# ENV TEST=1
EXPOSE 8070
