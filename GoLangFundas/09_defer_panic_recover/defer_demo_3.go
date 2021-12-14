package main

import "fmt"

func quitiont(divident, divisor int) (q int, errMsg string) {
	defer func() {
		if err := recover(); err != nil {
			errMsg = "Zero is not a qualified divisor"
		}
	}()

	return divident / divisor, ""
}

func main() {
	fmt.Println(quitiont(10, 5))
	fmt.Println(quitiont(10, 0))
	fmt.Println(quitiont(10, 3))
}
