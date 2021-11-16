package main

import (
	"fmt"
)

func main() {
	// arrays
	var x [2]string
	x[0] = "Hello"
	x[1] = "World"

	// alternative way
	var x := [...]string{"Hello", "World"}

	// access to elements
	fmt.Println(x[0])

	for _, v := range x {
		fmt.Println(v)
	}

	// slices
	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}
	// add elements
	days = append(days, "Saturday", "Sunday")

	days2 := days
	days2[0] = "X-day"
	fmt.Println(days[0], "", days2[0]) // same value for both

	days3 := make([]string, len(days))
	copy(days3, days)

	days3[0] = "Monday"

	fmt.Println(days)
	fmt.Println(days3) // diff values for
}
