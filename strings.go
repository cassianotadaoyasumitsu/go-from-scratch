package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	s1 := "Hello!"
	s2 := "Hエllo!" // contain one katakana character!!!
	fmt.Println("Is S1 equal to S2? :", s1 == s2)
	fmt.Printf("The length of s1 string is: %v\n", len(s1))
	fmt.Printf("The length of s2 string is: %v\n", len(s2))

	fmt.Println("The length of s1 string is:", utf8.RuneCountInString(s1))
	fmt.Println("The length of s2 string is:", utf8.RuneCountInString(s2))

	fmt.Println(strings.Contains("Drunken sailor", "sail"))
}
