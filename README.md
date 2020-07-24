# ngrok-server

### Dockerfile

Used to create docker image

### generator.go

To generate executation command

### output

executation command

### certificate/

storing certificates

### config/

- ngrok.conf - recording the nginx rediret for ngrok
- running-ports - recording the port usage

### ngrok/

- ngrok source code
- ngrok client - windows 32/64
- ngrok server

---

## Usage

1. update the value in Dockerfile 
```
ENV NGROK_DOMAIN spicyworld.api.redpayments.com.au
ENV HTTP_PORT 9043
ENV TUNNEL_ADDR_PORT 9044
```

2. append running-ports

3. ``` go run ./generator.go```
   The script will read Dockerfile and generate **output**
   
4. `output` contains the nginx setting, docker build command, docker run command, the server address.
5. Now push it to server
6.  
   - docker build 
   - docker run
   - vim /etc/nginx/conf.d/ngrok.conf
   - nginx -t
   - systemctl restart nginx