package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

const NGROK_DOMAIN = "123444"
const HTTP_PORT = "9035"
const TUNNEL_ADDR_PORT = "9036"

func main() {
	fmt.Println(NGROK_DOMAIN, HTTP_PORT, TUNNEL_ADDR_PORT)
	var path = "certificate\\" + NGROK_DOMAIN
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

	// creating certificate files using cmd
	app := "mkdir"

	arg0 := "myfile"

	cmd := exec.Command(app, arg0)
	stdout, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Print(string(stdout))

}

func sub() {

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path) // for example /home/user

	os.Chdir("./certificate")
	a, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(a) // for example /home/user

	app := "echo"

	// arg0 := "-e"
	arg1 := "Hello world"
	// arg2 := "\n\tfrom"
	// arg3 := "golang"

	cmd := exec.Command(app, arg1)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print(string(stdout))
}
