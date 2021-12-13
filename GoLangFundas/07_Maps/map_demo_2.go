package main

import "fmt"

type employee struct {
	empId     int
	firstName string
	lastName  string
	basic     float32
}

func createNewEmployee() employee {
	var e employee

	fmt.Print("EmpId: ")
	fmt.Scanln(&e.empId)
	fmt.Print("First Name: ")
	fmt.Scanln(&e.firstName)
	fmt.Print("Last Name: ")
	fmt.Scanln(&e.lastName)
	fmt.Print("Basic: ")
	fmt.Scanln(&e.basic)

	return e
}

func acceptChoice() int {
	fmt.Println("1\tList Employees")
	fmt.Println("2\tAdd Employees")
	fmt.Println("3\tDelete Employees")
	fmt.Println("4\tSearch Employees")
	fmt.Println("5\tExit")

	var ch int
	fmt.Print("Choice: ")
	fmt.Scanln(&ch)

	return ch
}

func print(m map[int]employee) {
	fmt.Println("\nEmp#\tFirst Name\tLast Name\tBasic")
	for _, emp := range m {
		fmt.Printf("\n%d\t%s\t%s\t%f", emp.empId, emp.firstName, emp.lastName, emp.basic)
	}
}

func main() {

	emps := make(map[int]employee)

	ch := 0

	for ch != 5 {
		ch = acceptChoice()

		switch ch {
		case 1:
			print(emps)
		case 2:
			e := createNewEmployee()
			emps[e.empId] = e
			fmt.Println("Employee Added")
		case 3:
			fmt.Print("Emp Id: ")
			var key int
			fmt.Scanln(&key)
			delete(emps, key)
			fmt.Println("Employee Deleted")
		case 4:
			fmt.Print("Emp Id: ")
			var empId int
			fmt.Scanln(&empId)
			emp, isFound := emps[empId]
			if isFound {
				fmt.Println(emp)
			} else {
				fmt.Println("Employee Not Found")
			}
		case 5:
			fmt.Println("Program Terminated")
		default:
			fmt.Println("Unknown Menu Choice")
		}
	}
}
