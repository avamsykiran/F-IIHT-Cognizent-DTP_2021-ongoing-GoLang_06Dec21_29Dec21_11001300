package main

import "fmt"
import "errors"

func simpleInterest(p, t, r int) int {

	var errMsg string = "";

	if p < 0 {
		errMsg = "Invalid Loan Amount"
	}
	if r < 0 {
		errMsg += "\nInvalid rate of interest")
	}
	if t < 0 {
		errMsg += "Invalid time period"
	}

	if errMsg !="" {
		panic(errMsg);
	}

	return (p * t * r) / 100
}

func payableAmount(p, t, r int) (lA int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprint(e))
		}
	}()
	return p + simpleInterest(p, t, r), nil
}

func main() {
	fmt.Println(payableAmount(1000, 10, 2))
	fmt.Println(payableAmount(1000, -10, 2))
	fmt.Println(payableAmount(1000, 5, 2))
}
