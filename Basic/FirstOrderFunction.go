package main

import "fmt"

func add(x int, y int) {
	z := x + y

	fmt.Println(z)
}

func main() {

	add(10, 20)
}
