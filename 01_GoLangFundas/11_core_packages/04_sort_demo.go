package main

import (
	"fmt"
	"sort"
)

type contact struct {
	name   string
	mobile string
	age    int
}

type byName []contact

func (this byName) Len() int {
	return len(this)
}

func (this byName) Less(i, j int) bool {
	return this[i].name < this[j].name
}

func (this byName) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func main() {
	contacts := []contact{
		{"Vamsy", "9052224753", 34},
		{"Srinivas", "9948016004", 54},
		{"Anitha", "9948016114", 14},
		{"Zhahnavi", "9900016004", 34},
		{"TejaPal", "1148016004", 64},
	}

	sort.Sort(byName(contacts))

	for _, c := range contacts {
		fmt.Println(c)
	}
}
