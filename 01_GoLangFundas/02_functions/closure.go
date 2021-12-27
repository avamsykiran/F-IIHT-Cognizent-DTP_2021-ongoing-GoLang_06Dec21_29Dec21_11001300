package main

import "fmt"

func greetGenerator(greetNote string) func(string) (string, int) {
	count := 0

	return func(unm string) (string, int) {
		count++
		return greetNote + unm, count
	}
}

func main() {
	greetUserInHindi := greetGenerator("Namsaskar ")
	greetUserInEnglish := greetGenerator("Hello ")
	greetUserInTamil := greetGenerator("Vanakkam ")

	fmt.Println(greetUserInHindi("Vamsy"))
	fmt.Println(greetUserInEnglish("Vamsy"))
	fmt.Println(greetUserInTamil("Vamsy"))

	fmt.Println(greetUserInHindi("Vamsy"))
	fmt.Println(greetUserInEnglish("Vamsy"))
	fmt.Println(greetUserInTamil("Vamsy"))
}
