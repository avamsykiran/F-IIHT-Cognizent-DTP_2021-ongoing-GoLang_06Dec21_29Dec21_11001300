package main

import "fmt"

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {

		for j := 0; j < len(arr)-i-1; j++ {

			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}

		}
	}

	return arr
}

func main() {
	arr := []int{76, 45, 12, 98, 10, 3}
	fmt.Println(arr)
	arr = bubbleSort(arr)
	fmt.Println(arr)
}
