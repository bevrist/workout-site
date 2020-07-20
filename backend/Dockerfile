FROM golang:latest
WORKDIR /app
COPY ./backend .
COPY ./common ../common
RUN go get -d -v ./...
# RUN go get -v github.com/gorilla/mux github.com/lib/pq
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o backend .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /app .
CMD ["./backend"]

EXPOSE 8090