package main

import "fmt"

type loan struct {
	p, t, r float64
}

func (l loan) simpleInterest() float64 {
	return (l.p * l.t * l.r) / 100
}

func (l loan) payableAmount() float64 {
	return l.p + l.simpleInterest()
}

func main() {
	l1 := loan{
		p: 10000,
		t: 1,
		r: 12,
	}

	fmt.Println(l1)
	fmt.Println(l1.simpleInterest())
	fmt.Println(l1.payableAmount())
}
