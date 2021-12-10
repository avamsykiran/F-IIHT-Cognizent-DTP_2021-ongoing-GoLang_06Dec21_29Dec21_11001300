package main

import (
	"fmt"
	"reflect"
)

type pen struct {
	nib      float32
	inkColor string
}

func main() {
	p1 := pen{0.4, "blue"}

	fmt.Printf("%T\n", p1)
	fmt.Println(reflect.TypeOf(p1))

	x := 34
	y := 3.14
	z := true
	l := "Hello"
	m := '@'

	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)
	fmt.Printf("%T\n", l)
	fmt.Printf("%T\n", m)
}
