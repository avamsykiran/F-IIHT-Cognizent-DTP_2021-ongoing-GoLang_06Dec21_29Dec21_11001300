package main

import "fmt"
import "regexp"

func main() {
	pattern := regexp.MustCompile("[6-9][0-9]{9}")

	str := "You may call me at 9052224753"

	fmt.Println(pattern.FindStringIndex(str))
	fmt.Println(pattern.FindString(str))

}
