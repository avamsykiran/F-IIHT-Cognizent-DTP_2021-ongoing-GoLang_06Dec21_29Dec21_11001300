package main

import "fmt"

func main() {
	s := "VAMSY"

	for i := 0; i < len(s); i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%c", s[j])
		}
		fmt.Println()
	}
}
