package main

import (
	"fmt"

	"greetapp.cts.com/goodbye"
	"greetapp.cts.com/hello"
)

func main() {
	fmt.Println(hello.InTelugu())
	fmt.Println(goodbye.InTelugu())

	fmt.Println(hello.InHindi())
	fmt.Println(goodbye.InHindi())

	fmt.Println(hello.InEnglish())
	fmt.Println(goodbye.InEnglish())
}
