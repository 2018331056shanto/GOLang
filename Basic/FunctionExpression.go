package main

import "fmt"

func add(a int, b int) {

	c := a + b
	fmt.Println("From Inside Global Scope")
	fmt.Println(c)

}

func sum() {

	add(2, 3)
}

func main() {

	sum()
	add(20999,2222)
	add := func(a int, b int) {

		c := a + b
		fmt.Println("From Inside Local Scope")
		fmt.Println(c)

	}

	add(2, 3)
}
func init() {

	fmt.Println("I am the first function that is executed first ")
}
