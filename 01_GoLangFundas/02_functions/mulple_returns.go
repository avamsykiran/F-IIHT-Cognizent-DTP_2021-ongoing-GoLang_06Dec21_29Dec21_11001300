package main

import "fmt"

func arithmeticOperations(n1, n2 int) (int, int, int, int, int) {
	return n1 + n2, n1 - n2, n1 * n2, n1 / n2, n1 % n2
}

func main() {
	sum, dif, prd, qut, rem := arithmeticOperations(25, 7)

	fmt.Printf("Sum:%d\tDif:%d\tPrd:%d\tQut:%d\tRem:%d", sum, dif, prd, qut, rem)
}
