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

func main() {
	thd1 := generateSeries(1, 10)
	thd2 := generateSeries(100, 110)

	for {
		select {
		case data, isOpen := <-thd1:
			fmt.Println("From First Thread>> ", data, isOpen)
			if !isOpen { //if channel is clsoed
				thd1 = nil
			}
		case data, isOpen := <-thd2:
			fmt.Println("From Second Thread>> ", data, isOpen)
			if !isOpen { //if channel is clsoed
				thd2 = nil
			}
		}
		if thd1 == nil && thd2 == nil {
			break
		}
	}
}
