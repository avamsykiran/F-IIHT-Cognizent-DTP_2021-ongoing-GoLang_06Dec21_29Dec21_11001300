package main

import "fmt"

func main(){
	var x int =1;

	switch(x) {
	case 0:
		fmt.Println("Good Morning")
	case 1:
		fmt.Println("Good Afternoon")
		fallthrough
	case 2:
		fmt.Println("Good Evening")
	default:
		fmt.Println("Good Night")
	}

}