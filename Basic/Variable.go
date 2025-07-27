package main

import (
	"fmt"
)

func main() { //main function starts from here

	var name string

	name = "Ashraful Islam Shanto"

	fmt.Println(name)

	var (
		defaultInt     int
		defaultFloat   float64
		defaultString  string
		defaultBool    bool
		defaultPointer *int
	)

	fmt.Println("Default values of Go variable types:")
	fmt.Printf("int: %d\n", defaultInt)
	fmt.Printf("float64: %f\n", defaultFloat)
	fmt.Printf("string: '%s'\n", defaultString)
	fmt.Printf("bool: %t\n", defaultBool)
	fmt.Printf("pointer: %v\n", defaultPointer)

	var (
		age           int
		education     string
		instituteName string = "SUST"
		graduated     bool   = true
	)

	age = 10
	education = "BSc in Computer Science and Engineering"
	fmt.Printf("%d \n %s\n %s\n %t \n", age, education, instituteName, graduated)

	/*
		When we wish to declare same type of variable in a single line we can do so using this
	*/

	var maritialStatus, district, upozila string

	maritialStatus = "married"
	district = "Sherpur"
	upozila = "Sherpur Sadar"
	fmt.Printf("%s \n %s \n %s \n", maritialStatus, district, upozila)

	/*

		We can initialize the value in a variable in a single line using this manner
	*/

	var (
		school, colleage, gradeInSSC, gradeInHSC = "Sherpur Govt Victoria Academy", "Notre Dame College", 5, 5.00
	)

	fmt.Printf("%s\n %s\n %d \n %f \n ", school, colleage, gradeInSSC, gradeInHSC)

	/*
		When we want to initalize some value in a variable in the time of declartion we can use this format
		variable_name:=Value
		By this manner we can easily declare and initialize the value in a single line. There is no need to decalre the value
		Prior
		Go is intelligent enough to identify the type of based on the value we have initialized in the variable
	*/

	nationality := "Bangladeshi"
	hasPassport := false
	hasNIDS := true

	fmt.Printf("%s \n %t \n %t \n", nationality, hasPassport, hasNIDS)

	/*
		Suppose in the above example we have declared a variable called nationality and we want to change the value of that variable
		We can do so using the same format as we have used to declare the variable
	*/

	nationality = "Bangladeshi-American"
	hasPassport = true
	hasNIDS = false
	fmt.Printf("%s \n %t \n %t \n", nationality, hasPassport, hasNIDS)
	/*
		Once we declare a variable with certain type we can not change the type of that variable
		For example if we declare a variable with int type then we can not assign a string
		or a float value to that variable
		We can only assign a value of that type or a type which can be converted to
		that type without any loss of data
	*/

	/*
		This line will cause a compile-time error because we are trying to assign an int to a string variable

	*/
	// nationality=10;
}
