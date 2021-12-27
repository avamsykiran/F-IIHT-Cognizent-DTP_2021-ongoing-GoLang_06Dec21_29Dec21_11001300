package main

import "fmt"

func greetGenerator(greeting string) func(string) string {
	return func(unm string) string {
		return greeting + " " + unm
	}
}

func main() {

	greetUserInEng := greetGenerator("Hello")
	greetUserInHin := greetGenerator("Namaskar")
	greetUserInTel := greetGenerator("Namaskaram")
	greetUserInTam := greetGenerator("Vanakkam")

	fmt.Println(greetUserInEng("Vamsy"))
	fmt.Println(greetUserInHin("Vamsy"))
	fmt.Println(greetUserInTel("Vamsy"))
	fmt.Println(greetUserInTam("Vamsy"))
}
