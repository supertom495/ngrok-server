package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "io/ioutil"
)

func main() {
    file, err := os.Open("Dockerfile")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    NGROK_DOMAIN := ""
    HTTP_PORT := ""
    TUNNEL_ADDR_PORT := ""

    for scanner.Scan() {
        s := strings.Split(scanner.Text(), " ")
        if s[0] == "ENV" {
            if s[1] == "NGROK_DOMAIN"{
                NGROK_DOMAIN = s[2]
            }
            if s[1] == "HTTP_PORT"{
                HTTP_PORT  = s[2]
            }
            if s[1] == "TUNNEL_ADDR_PORT"{
                TUNNEL_ADDR_PORT = s[2]
            }
        }
        
    }

    fmt.Println(NGROK_DOMAIN, HTTP_PORT, TUNNEL_ADDR_PORT)


    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }


	var nginxConf = ""
	nginxConf += "server {\n"
	nginxConf += "        listen          80;\n"
	nginxConf += "        listen          [::]:80;\n"
	nginxConf += fmt.Sprintf("        server_name     %s www.%s;\n", NGROK_DOMAIN, NGROK_DOMAIN)
	nginxConf += "        location / {\n"
	nginxConf += fmt.Sprintf("                proxy_pass http://www.%s:%s; \n", NGROK_DOMAIN, HTTP_PORT)
	nginxConf += "        }    \n"
	nginxConf += "}    \n"

	err = ioutil.WriteFile("nginxConf", []byte(nginxConf), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}


    var docker = ""
    docker += fmt.Sprintf("docker build -t %s:latest .\n", NGROK_DOMAIN)
	docker += fmt.Sprintf("docker run -t --name %s -p %s:%s -p %s:%s %s:latest\n", NGROK_DOMAIN, HTTP_PORT, HTTP_PORT, TUNNEL_ADDR_PORT, TUNNEL_ADDR_PORT, NGROK_DOMAIN)
	err = ioutil.WriteFile("docker", []byte(docker), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}


}