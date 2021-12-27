package main

import (
	"fmt"
	"time"
)

func generateSeries(threadName string, start int, end int) {
	for i := start; i <= end; i++ {
		fmt.Println(threadName, ">> ", i)
		time.Sleep(time.Second)
	}
}

func main() {
	go generateSeries("thd1", 1, 10)
	go generateSeries("thd2", 100, 110)
	time.Sleep(10 * time.Second)
}
