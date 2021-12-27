package main

import (
	"fmt"
	"strings"
)

func main(){
	s1 := "Vamsy"

	fmt.Println(strings.ToLower(s1))
	fmt.Println(strings.ToUpper(s1))
	fmt.Println(strings.ToTitle(s1))
	fmt.Println(strings.Repeat(s1,10))
}