package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func server(port string, signal chan struct{}) {
	// listen on a port
	ln, err := net.Listen("tcp", ":"+port)

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
	close(signal)
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

func client(port string, signal chan struct{}) {
	// connect to the server
	c, err := net.Dial("tcp", "127.0.0.1:"+port)

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
	close(signal)
}

func main() {
	serverSignal := make(chan struct{})
	clientSignal := make(chan struct{})

	var serverPort, clientPort string

	fmt.Println("Port To Listen On: ")
	fmt.Scanln(&serverPort)
	go server(serverPort, serverSignal)

	fmt.Println("Port To Listen From: ")
	fmt.Scanln(&clientPort)
	go client(clientPort, clientSignal)

	<-serverSignal
	<-clientSignal
}
