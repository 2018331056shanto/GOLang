package main

import (
	"fmt"
	"math"
)

func main() {
	var t int
	fmt.Scan(&t)
	for t > 0 {
		var a, b, c int
		fmt.Scan(&a, &b, &c)
		// fmt.Println(a,b,c)
		if a-1 < int(math.Abs(float64(c-b))+math.Abs(float64(c)-1)) {
			fmt.Println(1)
		} else if a-1 > int(math.Abs(float64(c-b))+math.Abs(float64(c)-1)) {
			fmt.Println(2)
		} else {
			fmt.Println(3)
		}
		t--
	}
}

// 1 2 1
