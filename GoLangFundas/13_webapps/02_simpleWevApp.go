package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {

		responseBody := "<h1>This is a simple golang web</h1>"

		responseBody += "<hr/>"
		responseBody += "<br /> Method: " + req.Method
		responseBody += "<br /> Path: " + req.URL.Path

		fmt.Fprintf(writer, responseBody)

	})

	//set up the server and lsitin at a specific port 4545
	fmt.Println("Sever is listining at 4545")
	err := http.ListenAndServe(":4545", nil)
	if err != nil {
		fmt.Println(err)
	}
}
