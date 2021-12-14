package main

import "fmt"

func f1(a int) {
	fmt.Println("On f1 call ", a)
	a++
	fmt.Println("On f1 complete ", a)
}

func f2(a *int) {
	fmt.Println("On f2 call ", *a)
	*a++
	fmt.Println("On f2 complete ", *a)
}

func main() {
	var x int = 45
	f1(x)
	fmt.Println("after f1 call from main ", x)
	f2(&x)
	fmt.Println("after f2 call from main ", x)
}
