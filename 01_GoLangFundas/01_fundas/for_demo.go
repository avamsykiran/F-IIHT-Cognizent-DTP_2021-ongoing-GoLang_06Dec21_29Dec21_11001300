package main

import "fmt"

func main() {

	var i int

	for i = 1; i <= 10; i++ {
		fmt.Println(i)
	}

	for i <= 20 {
		fmt.Println(i)
		i++
	}
}
