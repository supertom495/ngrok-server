FROM ubuntu:18.04

LABEL maintainer="supertom495@gmail.com"

ENV NGROK_DOMAIN test.xiyantong.pw
ENV HTTP_PORT 9025
ENV TUNNEL_ADDR_PORT 9026

RUN apt-get update && \
    apt-get install openssl -y
RUN cd /root && openssl rand -out .rnd 1804495
RUN mkdir ngrok && cd ngrok

WORKDIR /root/ngrok

COPY ./ngrok/bin/ngrokd /root/ngrok
RUN openssl genrsa -out rootCA.key 2048 && \
    openssl req -x509 -new -nodes -key rootCA.key -subj /CN=$NGROK_DOMAIN -days 5000 -out rootCA.pem && \
    openssl genrsa -out device.key 2048 && \
    openssl req -new -key device.key -subj /CN=$NGROK_DOMAIN -out device.csr && \
    openssl x509 -req -in device.csr -CA rootCA.pem -CAkey rootCA.key -CAcreateserial -out device.crt -days 5000


EXPOSE $HTTP_PORT
EXPOSE $TUNNEL_ADDR_PORT

# CMD ["-tlsKey=/root/ngrok/device.key", "-tlsCrt=/root/ngrok/device.crt", "-domain=\"test.xiyantong.pw\"", "-httpAddr=\":9025\"", "-httpsAddr=\":\"", "-tunnelAddr=\":9026\""] 
# CMD [ \
#         "-tlsKey=/root/ngrok/device.key", \
#         "-tlsCrt=/root/ngrok/device.crt", \
#         "-domain=${NGROK_DOMAIN}", \
#         "-httpAddr=:${HTTP_PORT}", \
#         "-httpsAddr=:", \
#         "-tunnelAddr=:${TUNNEL_ADDR_PORT}" \
#     ] 

CMD -tlsKey=/root/ngrok/device.key -tlsCrt=/root/ngrok/device.crt -domain=${NGROK_DOMAIN} -httpAddr=:${HTTP_PORT} -httpsAddr=: -tunnelAddr=:${TUNNEL_ADDR_PORT}

ENTRYPOINT ["/root/ngrok/ngrokd"]
# ENTRYPOINT ["/bin/bash"]