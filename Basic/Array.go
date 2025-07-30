package main

import (
	"fmt"
)

// The type [n]T is an array of n values of type T.
// var a [10]int
func main() {
	var a [2]string         //array of size 2
	a[0] = "Hello"          //Zero index of "a" has "Hello"
	a[1] = "World"          //1st index of "a" has "World"
	fmt.Println(a[0], a[1]) // will print Hello World
	fmt.Println(a)
	//  using Printf and used the %q “verb” to print each element quoted.
	fmt.Printf("%q\n", a)
	primes := [6]int{2, 3, 5, 7, 9}
	fmt.Println(primes)
	// fmt.Printf("%q\n",primes)

	// Finally, you can use an ellipsis to use an implicit length when you pass the values:
	// the size of the array is dynamic like vector in C++
	b := [...]string{"a", "b", "c", "d", "e"}
	fmt.Printf("%q\n", b)
	// Multi Dimensional Arrya
	y := [...][3]int{{1, 2, 4}, {3, 4, 2}}
	fmt.Println(y)

	var arr0 [5]int // All elements default to 0

	var arr1 = [5]int{1, 2, 3, 4, 5}

	arr2 := [5]int{1, 2, 3, 4, 5}
	arr3 := [...]int{10, 20, 30} // Length = 3
	// nitialize with index-value pairs
	arr4 := [5]int{1: 100, 3: 200} // [0 100 0 200 0]
	var matrix [2][3]int = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	matrix1 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	arr5 := [2]*int{new(int), new(int)}
	*arr5[0] = 10
	*arr5[1] = 20

	fmt.Println(arr0) // [0 0 0 0 0]
	fmt.Println(arr1) // [1 2 3 4 5]
	fmt.Println(arr2) // [1 2 3 4 5]
	fmt.Println(arr3) // [10 20 30]
	fmt.Println(arr4) // [0 100 0 200 0]
	fmt.Println(matrix) // [[1 2 3] [4 5 6]]
	fmt.Println(matrix1) // [[1 2 3] [4 5 6]]
	fmt.Println(arr5) // [0x... 0x...]

}
