// To define methods on a type you don’t “own”, you need to define an alias​ for the type you want to extend.
//  type alias declaration, which introduces an alternate name for an existing type.
package main

import (
	"fmt"
	
)

type person struct{
	x,y string
}
type anotherPerson person//here we will not get Greet() function
type againPerson=person// this is type aliasing now we have Greet() function also 
func (p person) Greet(){
	fmt.Println("Hello welcome!")
}
func(a againPerson) Celebrate(){
	fmt.Println("Hurrah! we won the game")
}


func main() {
	x:=person{
		x: "Shanto",
		y: "Titu",
	}
	x.Greet()

	y:=anotherPerson{
		x: "Tufail",
		y: "Arnob",
	}
	// we will not get Greet() function as it is not added when we make another type 
	fmt.Println(y.x)
	z:=againPerson{
		x: "Shanto",
		y: "Ala",
	}
	z.Greet()
	z.Celebrate()
}