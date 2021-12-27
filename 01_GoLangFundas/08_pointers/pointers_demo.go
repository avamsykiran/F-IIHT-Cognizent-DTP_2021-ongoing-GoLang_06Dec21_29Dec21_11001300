package main

import "fmt"

func main(){

	n := 43;

	fmt.Println(n)
	fmt.Println(&n)

	ptr := &n
	fmt.Println(ptr)
	fmt.Println(*ptr)
	fmt.Println(&ptr)
}