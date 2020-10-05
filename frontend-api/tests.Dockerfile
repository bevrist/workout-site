FROM golang:latest
RUN go get -v github.com/gorilla/mux
RUN go get -v github.com/google/go-cmp/cmp
WORKDIR /frontend-api
COPY ./common ../common
COPY ./frontend-api .
CMD ["go", "test"]

ENV FRONTEND_API_SERVICE_ADDRESS="localhost:8888"
