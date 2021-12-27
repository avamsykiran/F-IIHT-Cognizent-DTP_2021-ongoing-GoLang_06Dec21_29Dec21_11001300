package main

import (
	"fmt"
	"time"
)

type passenger struct {
	pid    int
	seatNo int
}

type bus struct {
	maxSeats      int
	currentVacant int
}

func allocateSeat(pgrs []passenger, b *bus) chan struct{} {
	ch := make(chan struct{})

	go func() {
		for i := 0; i < len(pgrs); i++ {
			if (*b).currentVacant <= (*b).maxSeats {
				pgrs[i].seatNo = (*b).currentVacant
				time.Sleep(time.Millisecond) //create some latency.....
				(*b).currentVacant++
			} else {
				break
			}
		}
		close(ch)
	}()
	return ch
}

func main() {

	b := bus{15, 1}

	var pgrs [25]passenger

	for i := 0; i < len(pgrs); i++ {
		pgrs[i] = passenger{101 + i, 0}
	}

	ch1 := allocateSeat(pgrs[:10], &b)
	ch2 := allocateSeat(pgrs[10:15], &b)
	ch3 := allocateSeat(pgrs[15:], &b)

	<-ch1
	<-ch2
	<-ch3

	for _, p := range pgrs {
		fmt.Println(p.pid, "\t", p.seatNo)
	}
}
