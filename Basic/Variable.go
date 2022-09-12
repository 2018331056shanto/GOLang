package main

import (
	"fmt"
	"math"
	"net/http"
)

//Package fmt implements formatted I/O with functions analogous to C's printf and scanf
func main() { //main function starts from here

	//  var statement decleares a lost of variables
	//the name of the variable comes first and then the type name
	// var name type
	var int_val int;
	int_val = 10;
	fmt.Println(int_val);
	// multiple variable declearation
	var (
		name string
		age int
		location string
		salary float64
		maritial_status bool
		dimension complex128
		sex string="Male"
	)
	var (
		
		schoole,college,hometown,friendNo="Sherpur Govt Victoria Acedemy","Notre Dame College Mymensingh","Sherpur",50;
		

	)

	name="Shanto"
	age=24
	location="Akhaliya Ghat"
	salary=3333.22
	maritial_status=false
	dimension=1+5i
	// %v	the value in a default format
	// %T	a Go-syntax representation of the type of the value

	fmt.Printf("%s \n %d\n %s\n %s\n %f \n %t\n %v",name,age,sex,location,salary,maritial_status,dimension)
	fmt.Println(schoole,college,hometown,friendNo)
	// Inside a function, the := short assignment statement
	//  can be used in place of a var declaration with implicit type

	goal:="To be happpy"
	passion,dream:="Explore World","Live life with family"

	fmt.Printf("%s\n %s\n %s\n",goal,passion,dream)
	// A variable can contain any type, including functions:
	action:=func (x int, y int)  {
		fmt.Println(x+y)
		
	}
	action(15,20)
	fmt.Println(math.Pi)
	fmt.Printf("HTTP status OK uses code: %d", http.StatusOK)

}
