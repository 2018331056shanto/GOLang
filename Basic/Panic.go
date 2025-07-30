package main

import (
	"fmt"
)

/*
There is no Exception in Go
Exceptional situations are considered normal in go
Errors are returned to handle sucn situation
Sometimes a Go Program cannot figure out what to do and cannot continue to run. The program panics in this case
The panic is a buil-in funciton in go which stop the normal execution of current go routine.
When a function call panics, normal execdution of that function stops immediately.
*/
func main() {

	a, b := 10, 0

	/*
		When our application gives runtime exception our go runtime do not know what to do then our applicaition starts panic.
		This panic flows up to the go runtime decide to stop or kill the program and gives us error message


	*/

	// fmt.Println(a / b)

	/*
		We can explicitly use panic() function to make our application panic.
		When a panic happens and any function which might have defer within that same function will be executed before the panic actually takes place



	*/

	fmt.Println(a, b)

	defer fmt.Println("2 defferd")
	panic("Cnanot continue")
	fmt.Println("Post panic")

	mx:=max(1,111)
	mn:=min(1,111)

	fmt.Println(mx,mn)
}


