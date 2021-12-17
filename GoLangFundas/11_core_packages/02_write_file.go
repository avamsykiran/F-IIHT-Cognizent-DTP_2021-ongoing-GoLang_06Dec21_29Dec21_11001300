package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var fnm string
	fmt.Print("file name? ")
	fmt.Scanln(&fnm)

	file, err := os.Create(fnm)

	defer file.Close()

	if err == nil {
		var data, line string
		fmt.Println("Enter data and a #$* to indicate EOF")
		for {
			fmt.Scanln(&line)
			if strings.HasPrefix(line, "#$*") {
				break
			} else {
				data += "\n" + line
			}
		}
		file.WriteString(data)
	} else {
		fmt.Println(err)
	}
}
