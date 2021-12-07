package main

import "fmt"

func main() {

	greetUser := func(unm string) string {
		return "Hello " + unm
	}

	fmt.Println(greetUser("Vamsy"))
}
