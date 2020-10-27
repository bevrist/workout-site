FROM golang:1
RUN go get -v github.com/gorilla/mux
RUN go get -v github.com/google/go-cmp/cmp
# ADD https://github.com/golang/go/raw/master/lib/time/zoneinfo.zip /zoneinfo.zip
# ENV ZONEINFO /zoneinfo.zip
WORKDIR /frontend-web
COPY ./common ../common
COPY ./frontend-api .
CMD ["go", "test"]

ENV FRONTEND_WEB_SITE_URL=localhost:80
ENV FRONTEND_WEB_LISTEN_ADDRESS=0.0.0.0:80
ENV BACKEND_ADDRESS=localhost:8090
ENV AUTH_ADDRESS=localhost:8070