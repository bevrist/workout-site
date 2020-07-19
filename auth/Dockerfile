FROM golang:latest
WORKDIR /app
COPY ./auth .
COPY ./common ../common
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o auth .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /app .
ENTRYPOINT ["./auth"]

EXPOSE 8070
