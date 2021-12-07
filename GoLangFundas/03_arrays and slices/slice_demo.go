package main

import "fmt"

func main() {
	var nums = [10]int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}

	var slice1 = nums[:4]
	var slice2 = nums[4:7]
	var slice3 = nums[7:]

	fmt.Println(nums, "Len: ", len(nums), "Cap: ", cap(nums))
	fmt.Println(slice1, "Len: ", len(slice1), "Cap: ", cap(slice1))
	fmt.Println(slice2, "Len: ", len(slice2), "Cap: ", cap(slice2))
	fmt.Println(slice3, "Len: ", len(slice3), "Cap: ", cap(slice3))
}
