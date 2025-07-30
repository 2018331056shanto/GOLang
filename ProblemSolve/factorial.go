package main 

import "fmt"

func calculateFactorial(n int) int{

	res:=1

	for i:=2;i<=n;i++{

		res*=i
	}

	return res

}

func main(){


	var n int
	
	fmt.Scan(&n)

	factorial:=calculateFactorial(n)

	fmt.Println(factorial)
}
