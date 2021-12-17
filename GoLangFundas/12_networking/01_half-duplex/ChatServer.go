package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func server() {
	// listen on a port
	ln, err := net.Listen("tcp", ":9999")

	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		// accept a connection
		c, err := ln.Accept()

		if err != nil {
			fmt.Println(err)
			continue
		}

		// handle the connection
		ch := make(chan struct{})
		go handleServerConnection(c, ch)
		<-ch
		break
	}
}

func handleServerConnection(c net.Conn, ch chan struct{}) {
	// receive the message
	var msg string
	for {
		receiver := gob.NewDecoder(c)
		err := receiver.Decode(&msg)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Received: ", msg)
			if msg == "quit" {
				break
			}
		}
	}
	c.Close()
	close(ch)
}

func main() {
	server()
}
