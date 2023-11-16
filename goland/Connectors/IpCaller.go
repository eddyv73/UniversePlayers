package Connectors

import (
	"fmt"
	"net/http"
	"os"
)

func GetIp(endpoint string) string {
	fmt.Printf("client: sending request to %s\n", endpoint)

	requestURL := fmt.Sprintf(endpoint)
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("client: got response!\n")
	//print response
	res.Write(os.Stdout)
	fmt.Printf("client: status code: %d\n", res.StatusCode)
	requestInfo := res.Status

	return requestInfo
}
