package main

import "fmt"

type dummy struct {
	string
	bool
	int
}

func main() {
	d := dummy{"Apple", true, 450}

	d.bool = false

	fmt.Println(d)

	st := struct {
		field1 string
		field2 int
		field3 bool
	}{
		field1: "soem value",
		field2: 56,
		field3: false,
	}

	fmt.Println(st)
}
