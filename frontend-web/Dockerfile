FROM golang:latest
RUN apt-get update && apt-get install uglifyjs -y
WORKDIR /app
RUN go get -v github.com/gorilla/mux
COPY ./frontend/frontend.go .
COPY ./common ../common
# RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o frontend .
COPY ./frontend/www/ ./www
# minify all js and css files
# RUN find ./ -type f -name "*.js" -exec uglifyjs --compress --mangle -o {}.min {} \; -exec rm {} \; -exec mv {}.min {} \;
# RUN find ./ -type f -name "*.css" -exec uglifycss --compress --mangle -o {}.min {} \; -exec rm {} \; -exec mv {}.min {} \;


FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /app .
# before run substitute all webpage references to http://localhost:8080 & http://localhost:5500 with the env var FRONTEND_WEBSITE_URL (if present)
RUN echo 'if [[ -n "${FRONTEND_WEBSITE_URL}" ]]; then find ./www/ -type f -exec sed -i -e "s,http://localhost:8080,$FRONTEND_WEBSITE_URL,g" {} \;; find ./www/ -type f -exec sed -i -e "s,http://localhost:5500,$FRONTEND_WEBSITE_URL,g" {} \;; fi; ./frontend' > ./start.sh; chmod +x ./start.sh

CMD ["/bin/sh", "./start.sh"]

ENV FRONTEND_WEBSITE_URL=localhost:80

ENV FRONTEND_LISTEN_ADDRESS=0.0.0.0:80
ENV BACKEND_ADDRESS=localhost:8090
ENV AUTH_ADDRESS=localhost:8070
EXPOSE 8080
