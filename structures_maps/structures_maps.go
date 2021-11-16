package main

import "fmt"

func main() {
	type Employee struct {
		ID      int
		Name    string
		Salaray float64
	}

	var developer Employee
	developer.Salaray = 5000.0

	// other way to declare
	boss := Employee{Name: "John", Salaray: 20000.0}
	fmt.Println(boss)

	// Maps
	// map[KeyType]ValueType

	type Coords struct {
		Lat, Long float64
	}

	var m map[string]Coords
	m = make(map[string]Coords)

	m["New York"] = Coords{40.730610, -73.935242}
	m["Moscow"] = Coords{55.751244, 37.618423}

	fmt.Println(m["Moscow"].Lat)
}
