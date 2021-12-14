package main

import "fmt"

func quickSort(arr []int, low, high int) {
	if low < high {
		var mid = partition(arr, low, high)
		quickSort(arr, low, mid)
		quickSort(arr, mid+1, high)
	}
}

func partition(arr []int, low, high int) int {
	var midVal = arr[low]
	var i = low
	var j = high

	for i < j {
		for arr[i] <= midVal && i < high {
			i++
		}
		for arr[j] > midVal && j > low {
			j--
		}
		if i < j {
			var temp = arr[i]
			arr[i] = arr[j]
			arr[j] = temp
		}
	}

	arr[low] = arr[j]
	arr[j] = midVal
	return j
}

func main() {
	var arr = []int{23, 5, 6, 99, 56, 12, 3, 34}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
