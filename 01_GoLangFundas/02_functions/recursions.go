package main

import "fmt"

func factorial(n int) int64 {
	var result int64 = 1

	if n > 1 {
		result = int64(n) * factorial(n-1)
	}

	return result
}

func main() {
	fmt.Println(factorial(5))
}
