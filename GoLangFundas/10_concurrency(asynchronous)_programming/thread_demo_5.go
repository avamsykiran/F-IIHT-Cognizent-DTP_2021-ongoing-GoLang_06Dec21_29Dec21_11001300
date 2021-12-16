package main

import (
	"fmt"
	"time"
)

func generateSeries(threadName string, start int, end int) <-chan string {
	ch := make(chan string)

	go func() {

		for i := start; i <= end; i++ {
			output := fmt.Sprint(threadName, ">> ", i)
			ch <- output
			time.Sleep(time.Second)
		}

		close(ch)
	}()
	return ch
}

func main() {

	thd1 := generateSeries("thd1", 1, 10)
	thd2 := generateSeries("thd2", 100, 110)

	for {
		fmt.Println(<-thd1)
		fmt.Println(<-thd2)
	}
}
