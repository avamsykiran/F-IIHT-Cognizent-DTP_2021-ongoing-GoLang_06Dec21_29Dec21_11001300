package main

import "fmt"

type address struct {
	doorNum, street, city string
}

type employee struct {
	empId            int
	name             string
	basic            float32
	currentAddress   address
	parminantAddress address
}

type manager struct {
	employee
	allowence float32
}

func main() {
	emp := employee{
		empId: 101,
		name:  "Vamsy Kiran",
		basic: 45678.0,
		currentAddress: address{
			doorNum: "4-16/1",
			street:  "Gurramvari Colony",
			city:    "Visakhapatanm",
		},
		parminantAddress: address{
			doorNum: "4-16/1",
			street:  "Gurramvari Colony",
			city:    "Visakhapatanm",
		},
	}

	mang := manager{
		employee: employee{
			empId: 102,
			name:  "Indhikaa Valli",
			basic: 145678.0,
			currentAddress: address{
				doorNum: "4-16/1",
				street:  "Gurramvari Colony",
				city:    "Visakhapatanm",
			},
			parminantAddress: address{
				doorNum: "4-16/1",
				street:  "Gurramvari Colony",
				city:    "Visakhapatanm",
			},
		},
		allowence: 5678.0,
	}

	fmt.Println(emp)
	fmt.Println(mang)
}
