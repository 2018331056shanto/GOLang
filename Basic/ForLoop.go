package main

import (
	"fmt"
)

func main() {

	// sum := 0
	// for i := 0; i < 10; i++ {
	// 	sum += i
	// }
	// fmt.Println(sum)
	// // The init and post statements are optional.
	// x := 1
	// for x < 1000 {
	// 	fmt.Println(x)
	// 	x += x
	// }
	// fmt.Println(x)
	// // For is Go's "while"
	// y := 1
	// for y < 100 {
	// 	y += y
	// }

	// fmt.Println(y)

	var arr = []int{1, 2, 3, 4, 5, 6, 7, 8}

	for n := range arr {

		fmt.Println(n)
	}

	//i=index, n=index_value

	for i, n := range arr {

		fmt.Println(i, n)
	}

	for i, j := 0, 10; i < j; i, j = i+1, j-1 {

		fmt.Println(i, j)
	}

	for {


		fmt.Println("I love GO")
	}

}
