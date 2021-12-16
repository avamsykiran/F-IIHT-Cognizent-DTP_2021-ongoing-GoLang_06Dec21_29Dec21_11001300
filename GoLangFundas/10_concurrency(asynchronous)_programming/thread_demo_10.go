package main

import "fmt"

var msgSource = []string{
	"string1", "string2", "string3", "string4", "string5",
	"string6", "string7", "string8", "string9", "string10",
}

func producer(ch chan<- string) {
	for _, msg := range msgSource {
		fmt.Println(msg, " is produced")
		ch <- msg
	}
	close(ch)
}

func consumer(ch <-chan string, areAllConsumed chan struct{}) {

	for msg := range ch {
		fmt.Println(msg, " is consumed")
	}

	close(areAllConsumed)
}

func main() {

	ch := make(chan string)
	done := make(chan struct{})

	go producer(ch)
	go consumer(ch, done)

	<-done

}
