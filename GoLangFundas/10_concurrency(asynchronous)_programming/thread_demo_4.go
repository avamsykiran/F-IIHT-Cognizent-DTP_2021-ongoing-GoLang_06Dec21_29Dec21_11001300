package main

import (
	"fmt"
	"time"
)

func generateSeries(threadName string, start int, end int) <-chan struct{} {
	ch := make(chan struct{})

	go func() {
		for i := start; i <= end; i++ {
			fmt.Println(threadName, ">> ", i)
			time.Sleep(time.Second)
		}
		close(ch)
	}()

	return ch
}

func main() {
	s1 := generateSeries("thd1", 1, 10)
	s2 := generateSeries("thd2", 100, 110)

	<-s1
	<-s2
}
