package main

import "fmt"

// Standard or Named Function
func add(x, y int) {

	z := x + y
	fmt.Println(z)
}

func testFunction(test func(a, b int)) {

	fmt.Println("Hello from hello function")
	test(10, 20)
}
func main() {

	add(4, 7)

	function := func(a, b int) {
		c := a + b
		fmt.Println(c)
	}

	function(10, 20) // 30

	testFunction(function) // Hello from hello function
}

func init() {
	fmt.Println("I am the first function that is executed first ")
}
