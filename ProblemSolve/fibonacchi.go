package main

import "fmt"


func calculateFibonacchi(n int) int{

	if n<=2 {
		return 1

	}
	
	return calculateFibonacchi(n-1)+calculateFibonacchi(n-2)

}

func main(){

  var n int

  fmt.Scan(&n)


  fib:=calculateFibonacchi(n)

  fmt.Println(fib)

}
