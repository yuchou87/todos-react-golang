FROM nginx:latest

LABEL maintainer="yuchou87 <yuchou87@gmail.com>"

WORKDIR /usr/share/nginx/html

COPY ./build/ /usr/share/nginx/html
COPY default.conf /etc/nginx/conf.d/default.conf