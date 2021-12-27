package main

import "fmt"
import "os"

func main() {

	var fnm string
	fmt.Print("file name? ")
	fmt.Scanln(&fnm)

	file, err := os.Open(fnm)

	defer file.Close()

	if err == nil {
		stat, err := file.Stat()
		if err == nil {
			bs := make([]byte, stat.Size())
			_, err = file.Read(bs)
			if err == nil {
				fmt.Println(string(bs))
			}
		}
	}

	if err != nil {
		fmt.Println(err)
	}
}
