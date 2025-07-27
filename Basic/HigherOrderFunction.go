package main

import "fmt"

func applyOperation(x, y int, fun func(x, y, z int) int) {

	result := fun(20, 20, 10)
	fmt.Println(x, y)
	fmt.Println("Result of operation:", result)

}

func returnFun(x, y int, fun func(x, y, z int) int) func(x, y int) int {

	result := fun(x, y, 10)
	fmt.Println("Result of operation:", result)

	return func(x, y int) int {

		return x / y
	}

}

func main() {

	fun := func(x, y, z int) int {
		return x * y * z

	}

	applyOperation(10, 20, fun)

	fun2 := returnFun(10, 20, fun)

	fmt.Println(fun2(10, 2))

}
