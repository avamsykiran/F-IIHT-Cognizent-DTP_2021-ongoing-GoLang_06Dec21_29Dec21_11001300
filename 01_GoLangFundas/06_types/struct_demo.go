package main

import "fmt"

type pen struct {
	nib      float32
	inkColor string
}

func main() {
	p1 := pen{0.4, "blue"}

	fmt.Println(p1)
	fmt.Println(p1.nib)
	fmt.Println(p1.inkColor)
}
