package main

import "fmt"

type General struct {
	Name    string
	Age     int
	Address string
	Salary  float64
}

func printDetails(p General) {
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
	fmt.Println("Address:", p.Address)
	fmt.Println("Salary:", p.Salary)
}

// Receiver function
// A receiver function is a method that is associated with a specific type.
// It allows you to define functions that operate on instances of that type.
// In Go, you can define methods on types by specifying a receiver parameter in the function signature
// The receiver parameter is defined before the function name and acts as the first parameter of the method
// It allows you to access the fields and methods of the type within the method body.
// The receiver can be a value receiver or a pointer receiver.
// A value receiver receives a copy of the value, while a pointer receiver receives a reference to the value.
// This allows you to modify the original value if you use a pointer receiver.
func (p General) printPersonDetails() {
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
	fmt.Println("Address:", p.Address)
	fmt.Println("Salary:", p.Salary)
}

func main() {
	p1 := General{
		Name:    "John Doe",
		Age:     30,
		Address: "123 Main St",
		Salary:  50000.00,
	}
	p2 := General{
		Name:    "Jane Smith",
		Age:     28,
		Address: "456 Elm St",
		Salary:  60000.00,
	}

	p1.printPersonDetails()
	p2.printPersonDetails()
	printDetails(p1)
	printDetails(p2)

	fmt.Println("Name:", p1.Name)
	fmt.Println("Age:", p1.Age)
	fmt.Println("Address:", p1.Address)
	fmt.Println("Salary:", p1.Salary)
}
