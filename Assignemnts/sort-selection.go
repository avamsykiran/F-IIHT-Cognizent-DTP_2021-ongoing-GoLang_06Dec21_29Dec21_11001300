package main

import "fmt"

func selectionSort(selectionArr []int) []int {

	for i := 0; i < len(selectionArr)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(selectionArr); j++ {
			if selectionArr[j] < selectionArr[minIdx] {
				minIdx = j
			}
		}

		if minIdx != i {
			selectionArr[i], selectionArr[minIdx] = selectionArr[minIdx], selectionArr[i]
		}
	}
	return selectionArr
}

func main() {
	selectionArr := []int{10, 4, 7, 34, 12, 3, 4}
	fmt.Println(selectionSort(selectionArr))
}
