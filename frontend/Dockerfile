FROM golang:latest
WORKDIR /app
# RUN go get -v github.com/gorilla/mux
COPY ./frontend .
COPY ./common ../common
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o frontend .


FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /app .
# before run substitute all webpage references to localhost:8080 with the env var FRONTEND_LISTEN_ADDRESS (if present)
RUN echo 'if [[ -n "${FRONTEND_LISTEN_ADDRESS}" ]]; then find ./www/ -type f -exec sed -i -e "s/localhost:8080/$FRONTEND_LISTEN_ADDRESS/g" {} \;; fi; ./frontend' > ./start.sh; chmod +x ./start.sh

CMD ["/bin/sh", "./start.sh"]

EXPOSE 8080
