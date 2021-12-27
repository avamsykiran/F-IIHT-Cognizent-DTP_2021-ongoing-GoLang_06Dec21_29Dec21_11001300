package main

import "fmt"

func swap(n1, n2 int) {
	fmt.Printf("\nFrom Inside Swap, Values before swap: \t%d\t%d", n1, n2)
	t := n1
	n1 = n2
	n2 = t
	fmt.Printf("\nFrom Inside Swap, Values after swap: \t%d\t%d", n1, n2)
}

func swap2(n1, n2 *int) {
	fmt.Printf("\nFrom Inside Swap2, Values before swap: \t%d\t%d", *n1, *n2)
	t := *n1
	*n1 = *n2
	*n2 = t
	fmt.Printf("\nFrom Inside Swap2, Values after swap: \t%d\t%d", *n1, *n2)
}

func main() {
	var a, b int = 12, 56

	fmt.Printf("\nFrom Inside Main, Values before swap: \t%d\t%d", a, b)
	swap(a, b) //call by value
	fmt.Printf("\nFrom Inside Main, Values after swap: \t%d\t%d", a, b)
	swap2(&a, &b) //call by ref
	fmt.Printf("\nFrom Inside Main, Values after swap2: \t%d\t%d", a, b)
}
