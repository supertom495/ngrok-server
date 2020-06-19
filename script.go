package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

const NGROK_DOMAIN = "www.google.com"
const HTTP_PORT = "9035"
const TUNNEL_ADDR_PORT = "9036"

func main() {
	fmt.Println(NGROK_DOMAIN, HTTP_PORT, TUNNEL_ADDR_PORT)
	var path = "certificate/" + NGROK_DOMAIN
	_, err := os.Stat(path)

	// creating certificate directory
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(path, 0755)
		if errDir != nil {
			log.Fatal(err)
			return
		}
	}

	// change directory
	os.Chdir("./" + path)
	certificatePath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(certificatePath) // for example /home/user

	// generate certificate files using cmd
	generateCertificate()

	// generate Dockerfile
	os.Chdir("./../../")

	var dockerfile = ""
	dockerfile += "FROM ubuntu:18.04\n\n"
	dockerfile += "LABEL maintainer=\"supertom495@gmail.com\"\n\n"
	dockerfile += "COPY ./ngrok/bin/ngrokd /usr/local/bin\n\n"
	dockerfile += fmt.Sprintf("COPY ./certificate/%s /root/certificate\n\n", NGROK_DOMAIN)
	dockerfile += fmt.Sprintf("EXPOSE %s\n\n", HTTP_PORT)
	dockerfile += fmt.Sprintf("EXPOSE %s\n\n", TUNNEL_ADDR_PORT)
	dockerfile += fmt.Sprintf("CMD [\"/usr/local/bin/ngrokd\"]")

	err = ioutil.WriteFile("Dockerfile", []byte(dockerfile), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}

	// update ngrok.conf

}

func generateCertificate() {
	genRootCA()
	reqRootCA()
	genDevice()
	reqDevice()
	reqCrt()
}

func genRootCA() {
	// 1
	app := "openssl"
	arg0 := "genrsa"
	arg1 := "-out"
	arg2 := "rootCA.key"
	arg3 := "2048"
	cmd := exec.Command(app, arg0, arg1, arg2, arg3)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
	fmt.Println("Result: " + out.String())
}

func genDevice() {
	// 3
	app := "openssl"
	arg0 := "genrsa"
	arg1 := "-out"
	arg2 := "device.key"
	arg3 := "2048"
	cmd := exec.Command(app, arg0, arg1, arg2, arg3)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
	fmt.Println("Result: " + out.String())
}

func reqRootCA() {
	// 2
	app := "openssl"
	arg0 := "req"
	arg1 := "-x509"
	arg2 := "-new"
	arg3 := "-nodes"
	arg4 := "-key"
	arg5 := "rootCA.key"
	arg6 := "-subj"
	arg7 := "/CN=" + NGROK_DOMAIN
	arg8 := "-days"
	arg9 := "5000"
	arg10 := "-out"
	arg11 := "rootCA.pem"

	cmd := exec.Command(app, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
	fmt.Println("Result: " + out.String())
}

func reqDevice() {
	// 4
	app := "openssl"
	arg0 := "req"
	arg1 := "-new"
	arg2 := "-key"
	arg3 := "device.key"
	arg4 := "-subj"
	arg5 := "/CN=" + NGROK_DOMAIN
	arg6 := "-out"
	arg7 := "device.csr"

	cmd := exec.Command(app, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
	fmt.Println("Result: " + out.String())
}

func reqCrt() {
	// 5
	app := "openssl"
	arg0 := "x509"
	arg1 := "-req"
	arg2 := "-in"
	arg3 := "device.csr"
	arg4 := "-CA"
	arg5 := "rootCA.pem"
	arg6 := "-CAkey"
	arg7 := "rootCA.key"
	arg8 := "-CAcreateserial"
	arg9 := "-out"
	arg10 := "device.crt"
	arg11 := "-days"
	arg12 := "5000"

	cmd := exec.Command(app, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
	fmt.Println("Result: " + out.String())
}
