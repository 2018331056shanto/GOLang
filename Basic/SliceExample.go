package main

import "fmt"

func main() {

	/*
		x {
			pointer: nil
			len: 0
			cap: 0
		}

	*/

	var x []int

	/*

		Initially
		x {
			pointer: nil
			len: 0
			cap: 0
		}
			Now we are sending x to append() and we are passing value of x to append() which is
			 x {
			pointer: nil
			len: 0
			cap: 0
		}
			Now as append itself a function so during execution phase a new Stack Frame is created. And in that Stack Frame
			a variable named Slice is sotred in that Local Scope of Stack Frame.
			When and the value from x is stored there. And now it checks if the pointer is nil or not.
			As it is nil so it creates an array of size 1 and store the value 1 and as append() is a closure function so go compiler
			make escape analysis and the array is created in the heap memory and the pointer of that array is stored in a new variable slice .
			Escape analysis is a process where the Go compiler determines whether a variable can be allocated on the stack or needs to be allocated on the heap.
			As when append() function is executed a new Stack Frame is created and if we store the array in the stack memory then when the function execution is finished
			the stack frame is popped off and the memory is freed. So, we need to store the array in the heap memory so that it can be accessed later.
			And now the slice variable is returned to the main function and the value of x is updated with the value of slice.
			So now the value of x is
			x {
				pointer: 0x12345678 // some address
				len: 1
				cap: 1
			}
	*/
	x = append(x, 1)

	/*
		Now slice  x {
			pointer: 0x12345678 // some address
			len: 1
			cap: 1

		}
			and again we are sending x to append() and we are passing value of x to append() which is
			x {
				pointer: 26 // some address
				len: 1
				cap: 1
			}
			Now as append itself a function so during execution phase a new Stack Frame is created.
			And in that Stack Frame a variable named Slice is sotred in that Local Scope of Stack Frame.
			When and the value from x is stored there. And now it checks if the pointer is nil or not.
			As it is not nil so it checks the length and capacity of the slice.
			As the length is 1 and capacity is 1 so it creates a new array of size 2 and copy the value of the old array to the new array and store the value 2 in the new array.
			And as append() is a closure function so go compiler make escape analysis and the array is created in the heap memory and the pointer of that array
			is stored in a new variable slice.
			Escape analysis is a process where the Go compiler determines whether a variable can be allocated on the stack or needs to be allocated on the heap.
			As when append() function is executed a new Stack Frame is created and if we store the array in the stack memory then when the function
			execution is finished the stack frame is popped off and the memory is freed. So, we need to store the array in the heap memory so that it can be accessed later.
			And now the slice variable is returned to the main function and the value of x is
			x {
				pointer: 27 // some address
				len: 2
				cap: 2
			}
	*/
	x = append(x, 2)
	/*
		and again we are sending x to append() and we are passing value of x to append() which is
				x {
					pointer: 27 // some address
					len: 2
					cap: 2
				}
				Now as append itself a function so during execution phase a new Stack Frame is created.
				And in that Stack Frame a variable named Slice is sotred in that Local Scope of Stack Frame.
				When and the value from x is stored there. And now it checks if the pointer is nil or not.
				As it is not nil so it checks the length and capacity of the slice.
				As the length is 2 and capacity is 2 so it creates a new array of size cap*2 and copy the value of the old array to the new array and store the value[1,2] in the new array.
				And as append() is a closure function so go compiler make escape analysis and the array is created in the heap memory and the pointer of that array
				is stored in a new variable slice.
				Escape analysis is a process where the Go compiler determines whether a variable can be allocated on the stack or needs to be allocated on the heap.
				As when append() function is executed a new Stack Frame is created and if we store the array in the stack memory then when the function
				execution is finished the stack frame is popped off and the memory is freed. So, we need to store the array in the heap memory so that it can be accessed later.
				And now the slice variable is returned to the main function and the value of x is
				x {
					pointer: 29 // some address
					len: 3
					cap: 4
				}
	*/
	x = append(x, 3)

	/*
		Now the value of x which is x {
					pointer: 29 // some address
					len: 3
					cap: 4
				}
				and we are assigning the value of x to y. So, y will have the same pointer as x.
				So, y will have the value of x which is x {
					pointer: 29 // some address
					len: 3
					cap: 4
				}
	*/
	y := x
	/*
		and again we are sending x to append() and we are passing value of x to append() which is
				x {
					pointer: 29 // some address
					len: 3
					cap: 4
				}
				Now as append itself a function so during execution phase a new Stack Frame is created.
				And in that Stack Frame a variable named Slice is sotred in that Local Scope of Stack Frame.
				When and the value from x is stored there. And now it checks if the pointer is nil or not.
				As it is not nil so it checks the length and capacity of the slice.
				As the length is 3 and capacity is 4 now no new array is created instead it will just add the value 4 to the existing array.
				And as append() is a closure function so go compiler make escape analysis and the array
				is created in the heap memory and the pointer of that array is stored in a new variable slice.

				Escape analysis is a process where the Go compiler determines whether a variable can be allocated on the stack or needs to be allocated on the heap.
				As when append() function is executed a new Stack Frame is created and if we store the array in the stack memory then when the function
				execution is finished the stack frame is popped off and the memory is freed. So, we need to store the array in the heap memory so that it can be accessed later.
				And now the slice variable is returned to the main function and the value of x is
				x {
					pointer: 29 // some address
					len: 4
					cap: 4
				}

	*/
	x = append(x, 4)
	/*
		Currently the value of y is{
				pointer: 29 // some address
				len: 3
				cap: 4
			}

		And now we are sending the value of y to append() function similarly a new stack fram will get created a new variable s is created
		in that stack frame and the value of that variable  s will be {{
				pointer: 29 // some address
				len: 3
				cap: 4
			}
			Now it checks if the pointer is nil or not. As it is not nil so now it will check for len and capacity. As we can see that in our
			variable s len is less than cap. So in normal scenario as we have still one slot free instead of creating a new array it will just
			append an element in the array which is in the Heap Memory
			But here's the catch, in previous instruciton we have already added 1 element in the same Heap Memory after inserting the value of
			x became
			x {
				pointer: 29 // some address
				len: 4
				cap: 4
			}

			So there is no space available to append in y=append(y,5)instruction though y thinks it has one element left to insert
			So in this case the memory will be overriden and the value which is placed by x, 4 previously will now be overriden by y, 5
			so now y{
				pointer: 29 // some address
				len: 4
				cap: 4
			}

	*/
	y = append(y, 5)

	/*
	 Now we are updating the first pointer address value
	*/
	x[0] = 10
	fmt.Println("x:", x)
	fmt.Println("y:", y)

	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Capacity first", cap(slice))
	slice = append(slice, 6)
	fmt.Println("Capacity second", cap(slice))

	slice = append(slice, 7)
	fmt.Println("Capacity third", cap(slice))

	a := slice[4:]
	fmt.Println("Capacity fourth", cap(slice))

	z := changeSlice(a)
	fmt.Println("Capacity Sixth", cap(slice))

	fmt.Println(slice)
	fmt.Println(z)

	printNumbers(12, 3, 4, 1, 12, 1, 21, 12, 112, 1, 3)

}

func changeSlice(a []int) []int {

	a[0] = 10
	a = append(a, 11)
	fmt.Println("Capacity fifth", cap(a))

	return a
}

// a= [5,6,7]

// y=[10,6,7,11]

// x=1,2,3,4,10,6,7,
// y=10,6,7,11

/*
Variadic Function
In most situation when we declare a function we take parameters in various numbers. Parameters can be 1,2,3 so on
Normally we take parameters like below

func add( a,b,c int){
}

How many parameters our function will take, we define this in our function definition.
But there may be some situation where the nubmer of argument a code will send in our function is not defined 
In those situation we use Variadic Function

func name_of_function (varible_name ...type){

}
And here the variable_name is a Slice 


*/


func printNumbers(numbers ...int) {

	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))

}
