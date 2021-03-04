FROM golang:1

ENV GO111MODULE=off
RUN go get -v github.com/gorilla/mux
WORKDIR /app
COPY ./frontend-api/frontend-api.go .
COPY ./common ../common
# RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o frontend-api .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /app .
CMD ["./frontend-api"]

ENV FRONTEND_LISTEN_ADDRESS=0.0.0.0:80
ENV BACKEND_ADDRESS=localhost:8090
ENV AUTH_ADDRESS=localhost:8070
EXPOSE 8888
