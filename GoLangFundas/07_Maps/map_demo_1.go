package main

import "fmt"

func main() {
	map1 := make(map[string]string)

	map1["9052224753"] = "Vamsy Kiran"
	map1["9052224755"] = "Varun Tej"

	fmt.Println(map1)

	fmt.Println("Mobile\t\tName")
	for mobile, name := range map1 {
		fmt.Println(mobile, "\t", name)
	}

	map2 := map[string]float32{"Mangoes": 124.50, "Apples": 30.0, "Grapes": 145.90}
	fmt.Println(map2)

}
