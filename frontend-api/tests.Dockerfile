FROM golang:1

ENV GO111MODULE=off
RUN go get -v github.com/gorilla/mux
RUN go get -v github.com/google/go-cmp/cmp
RUN wget -O /zoneinfo.zip https://github.com/golang/go/raw/master/lib/time/zoneinfo.zip
ENV ZONEINFO /zoneinfo.zip
WORKDIR /frontend-api
COPY ./common ../common
COPY ./frontend-api .
CMD ["go", "test"]

ENV FRONTEND_API_SERVICE_ADDRESS="localhost:8888"
