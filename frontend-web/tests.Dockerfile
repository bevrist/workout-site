FROM golang:1
RUN go get -v github.com/gorilla/mux
RUN go get -v github.com/google/go-cmp/cmp
WORKDIR /frontend-web
COPY ./common ../common
COPY ./frontend-web .
CMD ["go", "test"]

ENV FRONTEND_WEB_ADDRESS=http://localhost
