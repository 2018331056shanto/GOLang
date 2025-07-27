package main

import "fmt"

func addE(x int, y int) {
	z := x + y
	fmt.Println(z)
}

func main() {

	addE(10, 20)

	//IIFE (Immediately Invoked Function Expression)
	func(a, b int) {
		c := a + b
		fmt.Println(c)
	}(10, 20)

}

func init() {
	fmt.Println("I am the first function that is executed first ")
}
