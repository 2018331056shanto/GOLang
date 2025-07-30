package main

import "fmt"

func a() {

	i := 0 //value of i=0

	fmt.Println("First ", i) //0
	/*
		what defer does that it doesn't run this line of code immediately. When our code execution encounters defer Go holds this instruction

		it stores fmt.Println("Second",current_value_of__i/0) and doesn't call now. It returns just befor exiting or returning from the function


	*/
	defer fmt.Println("Second ", i) //

	i = i + 1 //1 increment the value of i

	fmt.Println("Third", i) //1 print this instruction immediately
	defer fmt.Println("Fourth ", i)
	defer fmt.Println("Fifth ", i)

	return

}

//Named Return Values

func sum(a int, b int) (resutl int) {

	resutl = a + b
	return 
}
//Defer with Named Value Returned Function

func calulate()(resutl int){

	fmt.Println("First ",resutl)
	defer func(){

		resutl=resutl+10
		fmt.Println("defer ",resutl)
	}()
	resutl=5

	fmt.Println("Second ",resutl )
	return
}

func calc() int{

	resutl:=0

	fmt.Println("First ",resutl)
	defer func(){

		resutl=resutl+10
		fmt.Println("defer ",resutl)
	}()
	resutl=5

	fmt.Println("Second ",resutl )
	return resutl
}

func main() {

	// a()

	// res := sum(9, 10)
	// fmt.Println(res)
	a:=calulate()
	b:=calc()


	fmt.Println(a)
	fmt.Println(b)

}

/*
Output:

First  0
Third 1
Second  0

*/

/*

1. Compilation Phase
	During Compilation Phase Go reads our code from first to last and  all the functions,sturct & constants  in this case main() & a() fucntion will get stored in the Code Segment of the compiled file

2. Exectution Phase
	During Execution Phase Go again reads our code from top to bottom line by line. And at first all the instructions which are stored in the Code Segment
	area of Compiled File gets loaded inside the Code Segment of RAM.
	Now it will check for main() in the Code Segment and after finding it will Create a new Stack Frame inside the Stack Segment.

	Now it will start execution main() function line by line and it encounters a() and first it will search for Global Memory or Data Segment
	As a() is not in Data Segment so it will find a() in Code Segment.
	Now another Stack Frame will get created and all the instuctions inside a() will get executed.

	First variable i will get memory allocation inside stack frame

	Now "First 0" will get printed

	Now it comes to defer and the whole statment will get stored with the value of i this statement is not executed yet. It just hold the executeion of this statement
	and stores it somewhere else

	Now the value of i will get incremented by 1.

	And then the value of i which is 1 will get printed now.

	And at last the function a() will return.

	In normal scenario when a function completed it's execution it's stack frame get's destroyed and removed from Stack Segment.

	But now in our case before exiting the Function Go Runtime will call the stored function which is fmt.Println("Second",0)

	So another stack frame get's created for Println() function and after completing Printing that Stack Frame get's removed and finally our a()
	function returns and it's Stack Frame gets destroyed.


*/

/*

	What happens when we use defer with named returned function and normal function 

	In case of named returned value function
	1. All code executes
	2. defer functions are stored in heap or stack in a linked list format, and in stack fram there is something named Defer List POinter
		which store the memory address of derfer functions memory address
	3.  Execute all defer functions
	5. return the named variable

	In case of normal function 
	
	1. All code executes
	2. defer functions are stored in heap or stack in a linked list format, and in stack fram there is something named Defer List POinter
		which store the memory address of derfer functions memory address
	3. return value are stored in register and it it's the value which was present at that moment
	4. Executes all defer function


	Normally in Closure we take the inner function and all the variables that are used in that inner function 
	 store a copy of the outer function in Heap Memory Area. So even though the Stack Frame for the Outer Function 
	 is destroyed but the funciton and variables that are used in inner function are stored in heap and we can use that 

	But in case of defer when both the Outer Function and Inner Function both are present at the same time, in that case 
	closure are formed and instead of value of a variable which are used in the inner function but it was defined in outer function
	we store the reference in that closure variable
	 
*/
