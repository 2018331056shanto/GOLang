package main

import (
	"fmt"
)

type Usr struct {
	Name string
	Age  int
}

func aa(x *int) {
	*x = 100

	fmt.Println("Value of x inside function:", *x)
}

func printDetails(u *Usr) {

	u.Name = "Shamima"
	u.Age = 25
	fmt.Println("Name:", u.Name)
	fmt.Println("Age:", u.Age)
}

func print(numbers *[3]int) {
	fmt.Println("Address of array:", numbers)
	fmt.Println("Value of array:", *numbers)
}

func main() {

	x := 20

	fmt.Println("Value of x:", x)
	fmt.Println(&x)

	aa(&x)

	fmt.Println("Address of x:", x)

	p1 := Usr{
		Name: "Ashraful",
		Age:  25,
	}
	fmt.Println("Name:", p1.Name)
	fmt.Println("Age:", p1.Age)

	printDetails(&p1)
	fmt.Println("Name:", p1.Name)
	fmt.Println("Age:", p1.Age)

	p := &x

	fmt.Println("Address of x:", p)
	fmt.Println("Value of x:", *p)

	*p = 50
	fmt.Println("Value of x after modification:", *p)
	fmt.Println("Value of x after modification:", x)

	arr := [3]int{1, 2, 3}
	fmt.Println("Array:", arr)
	print(&arr)

}

// when we pass an array or any data type to a function, it is passed by value by default.
// This means that a copy of the value is made and passed to the function.
// Copy of the value is made, not the original value. And it's really time consuming.
// If we have a large number of data suppose array and when we are passing it into a function
//this means we are copying each and every element of the array to the function.
// So, to avoid this we can use pointer.
// A pointer is a variable that stores the memory address of another variable.
//So instead of copying every element we just pass the address of the variable.
