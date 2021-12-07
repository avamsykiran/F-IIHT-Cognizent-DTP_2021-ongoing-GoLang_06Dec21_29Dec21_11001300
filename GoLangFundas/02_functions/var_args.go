package main

import "fmt"

func sum(nums ...int) int {
	s := 0

	for _, val := range nums {
		s += val
	}

	return s
}

func min(nums ...int) int {
	m := nums[0]
	for _, val := range nums {
		if val < m {
			m = val
		}
	}
	return m
}

func main() {
	fmt.Printf("\nSum %d", sum(10, 20))
	fmt.Printf("\nSum %d", sum(10, 20, 30))
	fmt.Printf("\nSum %d", sum(10, 20, 30, 40, 50, 60))

	fmt.Printf("\nMin %d", min(20, 10))               //10
	fmt.Printf("\nMin %d", min(20, 10, 30))           //10
	fmt.Printf("\nMin %d", min(20, 10, 30, 5, 50, 8)) //5
}
