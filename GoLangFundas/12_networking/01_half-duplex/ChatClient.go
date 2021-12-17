package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func client() {
	// connect to the server
	c, err := net.Dial("tcp", "127.0.0.1:9999")

	if err != nil {
		fmt.Println(err)
		return
	}

	var msg string

	for {
		fmt.Println("Enter message: ")
		fmt.Scanln(&msg)
		fmt.Println("Sending", msg)

		err = gob.NewEncoder(c).Encode(msg)
		if err != nil {
			fmt.Println(err)
		}
		if msg == "quit" {
			break
		}
	}
	c.Close()
}

func main() {
	client()
}
