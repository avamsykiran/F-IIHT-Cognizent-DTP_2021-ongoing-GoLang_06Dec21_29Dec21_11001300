package main

import "fmt"

func greetUser(unm string) string {
	return "Hello " + unm
}

func main(){
	fmt.Println(greetUser("Vamsy"))
}