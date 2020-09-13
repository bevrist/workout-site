FROM golang:latest
# auth dependencies
RUN go get -v firebase.google.com/go firebase.google.com/go/auth github.com/gorilla/mux google.golang.org/api/option
# frontend dependencies
RUN go get -v github.com/gorilla/mux

WORKDIR /app
COPY ./common ./common
COPY ./auth ./auth
COPY ./frontend ./frontend

RUN echo "\
    cd /app/auth; \
    go run ./auth.go & \
    cd /app/frontend; \
    go run ./frontend.go & \
    sleep 1; go test" > start.sh; chmod 777 start.sh

CMD ["/bin/bash", "./start.sh"]

# frontend env vars
ENV FRONTEND_WEBSITE_URL=localhost:80
ENV FRONTEND_LISTEN_ADDRESS=0.0.0.0:80
ENV BACKEND_ADDRESS=localhost:8090
ENV AUTH_ADDRESS=localhost:8070
# auth env vars
ENV AUTH_LISTEN_ADDRESS=0.0.0.0:8070
