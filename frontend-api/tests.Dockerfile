FROM golang:latest
RUN go get -v github.com/gorilla/mux
WORKDIR /frontend-api
COPY ./common ../common
COPY ./frontend-api .
CMD ["go", "test"]

ENV FRONTEND_LISTEN_ADDRESS=0.0.0.0:80
ENV BACKEND_ADDRESS=localhost:8090
