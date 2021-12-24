package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(writer, "<h1>This is a simple golang web</h1>")
	})

	http.HandleFunc("/aboutUs", func(writer http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(writer, "<h1>We are a team of 8 pursuing GoLang.</h1>")
	})

	//set up the server and listen at a specific port 7777
	fmt.Println("Starting web server @ http://localhost:7777")
	err := http.ListenAndServe(":7777", nil)
	if err != nil {
		fmt.Println(err)
	}
}
