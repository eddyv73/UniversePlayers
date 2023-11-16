package main

import (
	"fmt"
	"net/http"
	"os"
)

const serverPort = 80

func main() {
	sendRequest()
}

// send request to ifconfi.me
func sendRequest() {
	requestURL := fmt.Sprintf("http://ifconfig.me:%d", serverPort)
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("client: got response!\n")
	//print response
	res.Write(os.Stdout)
	fmt.Printf("client: status code: %d\n", res.StatusCode)

}
