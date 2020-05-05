package main

import (
	"fmt"
	"net/http"
)

func main() {
	routes := []string{"/", "/api/kabdckabdck", "/api/special"}
	url := "http://127.0.0.1:5500"
	responses := map[string]string{}

	for _, route := range routes {
		resp, err := http.Get(url + route)
		if err != nil {
			fmt.Println("Error:", err)
		}
		body := make([]byte, 1000)
		n, _ := resp.Body.Read(body)
		responses[route] = string(body)[:n]
	}

	for k, v := range responses {
		fmt.Println("Route:", k)
		fmt.Println("Response:", v, "\n")
	}
}
