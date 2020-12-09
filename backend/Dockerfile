FROM golang:1
WORKDIR /app
RUN go get -v github.com/gorilla/mux
COPY ./common ../common
COPY ./backend/backend.go .
# RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o backend .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
RUN wget -O /zoneinfo.zip https://github.com/golang/go/raw/master/lib/time/zoneinfo.zip
ENV ZONEINFO /zoneinfo.zip
COPY --from=0 /app .
CMD ["./backend"]

EXPOSE 8090