package main

import "fmt"

func main() {

	var i int = 0

Loop:
	fmt.Println(i)
	i++
	if i <= 10 {
		goto Loop
	}
}
