package main

import (
	"bufio"
	"fmt"
	"os"

	employee "example.com/employee"
)

func main() {
	var empls []employee.Employee

	empls = append(empls, employee.Employee{Name: "Mike", Salary: 10000},
		employee.Employee{Name: "John", Salary: 9000})

	for _, e := range empls {
		fmt.Printf("%s - %f", e.Salary)
	}

	input := bufio.NewScanner(os.Stdin)
	input.Scan()
}
