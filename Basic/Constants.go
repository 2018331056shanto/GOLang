package main

import (
	"fmt"
)


const pi = 3.1415

func main() {
	// Constants are declared like variables, but with the const keyword
	// const can only be a Character, String, Boolean and Numeric value
	// const can not be declared usin := syntax
	const (
		Country    = "Bangladesh"
		Religion   = "Islam"
		Population = 160000000
	)

	fmt.Printf("%s\n %s\n %d\n", Country, Religion, Population)
	// The fmt.Sprintf() function in Go language formats according to a format specifier and returns the resulting string
	format := fmt.Sprintf("hello my name is %s I am %d years old", "Shanto", 24)
	fmt.Println(format)
	
}
