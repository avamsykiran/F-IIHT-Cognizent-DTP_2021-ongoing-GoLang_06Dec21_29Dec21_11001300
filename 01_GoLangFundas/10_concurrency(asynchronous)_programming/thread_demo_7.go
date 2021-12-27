package main

import (
	"fmt"
	"time"
)

func generateSeries(start int, end int) <-chan int {
	ch := make(chan int)

	go func() {

		for i := start; i <= end; i++ {
			ch <- i
			time.Sleep(time.Second)
		}

		close(ch)
	}()
	return ch
}

func printer(taskName string, ch <-chan int) chan struct{} {
	signal := make(chan struct{})

	go func() {
		for data := range ch {
			fmt.Println(taskName, ">>", data)
			time.Sleep(time.Second)
		}
		close(signal)
	}()

	return signal
}

func main() {
	thd1 := generateSeries(1, 10)
	thd2 := generateSeries(100, 110)

	sig1 := printer("Thread1", thd1)
	sig2 := printer("Thread2", thd2)

	<-sig1
	<-sig2
}
