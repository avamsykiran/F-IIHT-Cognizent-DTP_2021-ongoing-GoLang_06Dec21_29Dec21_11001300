package main

import "fmt"

func main() {
	var nums [5]int

	var sum, min, max int = 0, 0, 0

	fmt.Println("Enter 5 valeus: ")

	for i := 0; i < 5; i++ {
		fmt.Scanln(&nums[i])
	}

	min = nums[0]
	max = nums[0]

	for _, n := range nums {
		sum += n
		if min > n {
			min = n
		}
		if max < n {
			max = n
		}
	}

	fmt.Println("Sum: ", sum)
	fmt.Println("Min: ", min)
	fmt.Println("Max: ", max)
}
