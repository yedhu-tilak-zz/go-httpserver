package main

import (
	"io/ioutil"
	"net/http"
)

func handleAPIRoutes(res http.ResponseWriter, req *http.Request) {
	url := req.URL.Path

	res.Write([]byte("You have put a " + req.Method + " at " + url))
	if len(url) > len("/api/") {
		res.Write([]byte("\nParam is :" + url[len("/api/"):]))
	}
}

func handleAPISpecial(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte(req.URL.Path + ": You have accessed special API"))
}

func handleHome(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "text/html")

	body, _ := ioutil.ReadFile("index.html")
	res.Write(body)
}

func main() {

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hello from Go HTTP Server\n"))
		res.Write([]byte(req.URL.Path))
	})

	http.HandleFunc("/api/", handleAPIRoutes)
	http.HandleFunc("/api/special/", handleAPISpecial)
	http.HandleFunc("/home", handleHome)

	http.ListenAndServe("0.0.0.0:5500", nil)
}
