FROM ubuntu:20.04 as builder
ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update && apt-get install npm -y
RUN npm install --global google-closure-compiler html-minifier csso-cli
WORKDIR /app
COPY ./www .
# replace all commented redirects (// window.location.replace) with uncommented versions
# RUN find ./ -type f -exec sed -i -e "s@// window.location.replace@window.location.replace@g" {} \;
# replace all hardcoded test addresses (http://localhost:5500) with relative addresses
RUN find ./ -type f -exec sed -i -e "s,http://localhost:5500/,/,g" {} \;
# replace all hardcoded test addresses (http://localhost:8888) with relative API addresses /api/page
RUN find ./ -type f -exec sed -i -e "s,http://localhost:8888/,/api/,g" {} \;
# compile all web files into optimized and minified versions
# RUN find ./ -type f -name "*.js" -exec google-closure-compiler --js {} --js_output_file {}.min \; -exec mv {}.min {} \; \
#     && find ./ -type f -name "*.css" -exec csso {} -o {}.min \; -exec mv {}.min {} \; \
#     && find ./ -type f -name "*.html" -exec html-minifier --collapse-whitespace --remove-comments --remove-optional-tags --remove-script-type-attributes --remove-tag-whitespace --use-short-doctype --minify-css true --minify-js true {} -o {}.min \; -exec mv {}.min {} \;


FROM alpine:latest
# install nginx
RUN mkdir -p /run/nginx
RUN apk add nginx tini
# remove default site
RUN rm -rf /var/www/*
RUN rm -rf /etc/nginx/conf.d/*
# copy static site and config
COPY --from=builder /app /var/www/
COPY ./nginx.conf /etc/nginx/conf.d/nginx.conf
# link logs to docker output
RUN ln -sf /dev/stdout /var/log/nginx/access.log \
    && ln -sf /dev/stderr /var/log/nginx/error.log

RUN printf '#!/bin/sh \n\
    if [[ -n "${FRONTEND_API_URL}" ]]; then \n\
        sed -i -e "s,<API_ADDRESS>,$FRONTEND_API_URL,g" /etc/nginx/conf.d/nginx.conf \n\
    fi \n\
    echo "frontend-web listening at: 0.0.0.0:80" \n\
    echo "forwarding api requests to: $FRONTEND_API_URL" \n\
    nginx -g "daemon off;" \n\
    ' > /start.sh; chmod +x /start.sh

CMD ["tini", "--", "/start.sh"]

EXPOSE 80
ENV FRONTEND_API_URL=http://localhost:8888
