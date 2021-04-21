FROM golang:1

# ENV GO111MODULE=off
RUN go get -v github.com/go-redis/redis/v8
RUN go get -v github.com/gorilla/mux github.com/markbates/goth github.com/markbates/goth/gothic
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
# This SESSION_SECRET isnt actually secret, its just to supress a goth warning for an unused feature
ENV SESSION_SECRET="workout-app"
CMD ["/root/auth"]

# ENV AUTH_LISTEN_ADDRESS=0.0.0.0:8070
# ENV REDIS_CONNECTION_STRING="redis://user:pass:@localhost:6379/0"
# ENV PROVIDER_SECRET=""
# ENV PROVIDER_KEY=""
# ENV WEBSITE_BASE_URL="http://localhost:3000"
EXPOSE 8070
