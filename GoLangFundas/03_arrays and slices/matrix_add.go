package main

import "fmt"

func main() {
	var mat1 [2][3]int
	var mat2 [2][3]int
	var sum [2][3]int

	fmt.Println("Enter matrix 1: ")
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Scanln(&mat1[i][j])
		}
	}

	fmt.Println("Enter matrix 2: ")
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Scanln(&mat2[i][j])
		}
	}

	for i := 0; i < 2; i++ {
		fmt.Print("\n")
		for j := 0; j < 3; j++ {
			sum[i][j] = mat1[i][j] + mat2[i][j]
			fmt.Print("\t", sum[i][j])
		}
	}

}
