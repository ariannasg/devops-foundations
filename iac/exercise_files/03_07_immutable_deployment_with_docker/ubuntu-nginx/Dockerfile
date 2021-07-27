FROM ubuntu:latest

MAINTAINER Ernest Mueller <ernestmueller@theagileadmin.com>

ENV FOO bar

RUN apt-key adv --keyserver hkp://pgp.mit.edu:80 --recv-keys 573BFD6B3D8FBC641079A6ABABF5BD827BD9BF62 \
    && echo "deb http://nginx.org/packages/mainline/ubuntu/ xenial nginx" >> /etc/apt/sources.list \
    && echo "deb-src http://nginx.org/packages/mainline/ubuntu/ xenial nginx" >> /etc/apt/sources.list \
    && apt-get update \
    && apt-get install -y nginx

EXPOSE 80

CMD ["nginx","-g","daemon off;"]
