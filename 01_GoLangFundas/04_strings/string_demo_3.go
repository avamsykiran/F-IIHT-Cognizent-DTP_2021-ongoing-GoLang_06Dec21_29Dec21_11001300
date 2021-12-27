package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Split("This is a long sentence", " "))
	fmt.Println(strings.Join([]string{"rat", "cat", "dog"}, "->"))
	fmt.Println(strings.Contains("This is the soruce string", "the"))
	fmt.Println(strings.Index("This is the soruce string", "the"))
	fmt.Println(strings.HasPrefix("LN0012", "LN"))
	fmt.Println(strings.Replace("This is the soruce string", "s", "$", 2))
	fmt.Println(strings.ReplaceAll("This is the soruce string", "s", "$"))
}
