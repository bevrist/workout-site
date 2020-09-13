FROM golang:latest
RUN apt-get update && apt-get install uglifyjs -y
WORKDIR /app
RUN go get -v github.com/gorilla/mux
COPY ./frontend/frontend.go .
COPY ./common ../common
# RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o frontend .


FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /app .
CMD ["./frontend"]

ENV FRONTEND_LISTEN_ADDRESS=0.0.0.0:80
ENV BACKEND_ADDRESS=localhost:8090
ENV AUTH_ADDRESS=localhost:8070
EXPOSE 8888
