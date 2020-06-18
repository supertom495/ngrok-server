FROM ubuntu:18.04

RUN echo "hello world"


ADD ./ngrok/bin/ngrokd /usr/local/bin

ENV NGROK_DOMAIN harry.xiyantong.pw
ENV HTTP_PORT 9027
ENV TUNNEL_ADDR_PORT 9028

WORKDIR /usr/local/bin

MAINTAINER Xiyan

CMD ["/usr/local/bin/ngrokd", "-domain=\"adsf\"", "-httpAddr=\":asdfe\""]
