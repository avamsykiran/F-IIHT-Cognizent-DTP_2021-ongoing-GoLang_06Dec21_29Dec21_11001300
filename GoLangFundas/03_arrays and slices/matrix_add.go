package main

import "fmt"

func main() {
	var mat1 [2][2]int
	var mat2 [2][2]int
	var sum [2][2]int
	var prd [2][2]int

	fmt.Println("Enter matrix 1: ")
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Print("Ele#", i, j, ": ")
			fmt.Scanln(&mat1[i][j])
		}
	}

	fmt.Println("Enter matrix 2: ")
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Print("Ele#", i, j, ": ")
			fmt.Scanln(&mat2[i][j])
		}
	}

	fmt.Println("\nSum is ")
	for i := 0; i < 2; i++ {
		fmt.Print("\n")
		for j := 0; j < 2; j++ {
			sum[i][j] = mat1[i][j] + mat2[i][j]
			fmt.Print("\t", sum[i][j])
		}
	}

	fmt.Println("\nPrd is ")
	for i := 0; i < 2; i++ {
		fmt.Print("\n")
		for j := 0; j < 2; j++ {
			prd[i][j] = 0
			for k := 0; k < 2; k++ {
				prd[i][j] += (mat1[i][k] * mat2[k][j])
			}
			fmt.Print("\t", prd[i][j])
		}
	}
}
