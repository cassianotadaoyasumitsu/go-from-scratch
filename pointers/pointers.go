package main

import "fmt"

func main() {
	var X int = 100
	Z := &X
	fmt.Println(Z)  // get memory value
	fmt.Println(*Z) // get X value
}
