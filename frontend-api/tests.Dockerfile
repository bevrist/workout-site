FROM golang:latest
RUN go get -v github.com/gorilla/mux
WORKDIR /frontend-api
COPY ./common ../common
COPY ./frontend-api .
CMD ["go", "test"]

ENV FRONTEND_API_SERVICE_ADDRESS="localhost:8888"
