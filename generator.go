package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
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
			if s[1] == "NGROK_DOMAIN" {
				NGROK_DOMAIN = s[2]
			}
			if s[1] == "HTTP_PORT" {
				HTTP_PORT = s[2]
			}
			if s[1] == "TUNNEL_ADDR_PORT" {
				TUNNEL_ADDR_PORT = s[2]
			}
		}

	}

	fmt.Println(NGROK_DOMAIN, HTTP_PORT, TUNNEL_ADDR_PORT)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var output = "\n"
	output += "server {\n"
	output += "        listen          80;\n"
	output += "        listen          [::]:80;\n"
	output += fmt.Sprintf("        server_name     %s www.%s;\n", NGROK_DOMAIN, NGROK_DOMAIN)
	output += "        location / {\n"
	output += fmt.Sprintf("                proxy_pass http://www.%s:%s; \n", NGROK_DOMAIN, HTTP_PORT)
	output += "        }    \n"
	output += "}    \n"

	output += "\n\n---------------------------------------------------------------------------------------\n\n"

	output += fmt.Sprintf("docker build -t %s:latest .\n", NGROK_DOMAIN)
	output += "\n---------------------------------------------------------------------------------------\n\n"
	output += fmt.Sprintf("docker run -t --name %s -p %s:%s -p %s:%s %s:latest\n", NGROK_DOMAIN, HTTP_PORT, HTTP_PORT, TUNNEL_ADDR_PORT, TUNNEL_ADDR_PORT, NGROK_DOMAIN)
	output += "\n---------------------------------------------------------------------------------------\n\n"
	output += fmt.Sprintf("%s:%s\n", NGROK_DOMAIN, TUNNEL_ADDR_PORT)

	err = ioutil.WriteFile("output", []byte(output), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}

}
