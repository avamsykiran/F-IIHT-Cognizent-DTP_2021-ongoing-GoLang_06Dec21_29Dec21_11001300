package main

import "fmt"

func f() {
	fmt.Println("F called ")
	defer fmt.Println("This statement is defered")
	fmt.Println("F is complete")
}

func main() {
	f()
}
