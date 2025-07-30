package main

import "fmt"

var a = 10

const p = 15

/*

	So here we can see that we are returning a function. Now the inner funciton's existance will not remain if the Outer parent function do not exist
	When we are invoking outer() an new Stack Frame Gets created inside the main() stack which remains in Heap and a new stack frame get's created.
	And inide the stack fream all the lines gets executed. And normally a variable get's declared in the stack frame area with named money. And a inner function
	is also accessing the value of money which is stored in the Stack Frame. So now if we think when we are finish with execution of outer() function
	The stack frame get's popped off. So now think we are returning a funcing as return type and in that returned function we are using the variable of outer function
	which was sotred in the stack frame of outer function. Now after returning a inner function when we are invoking it later in the code
	and when a new execution frame get's created our inner funciton will search for money but it's not in there so normally we should have seen error.

	But we are not gettign error, though the stack frame which sotred the variable money is not there still we are able to access the variable money. THis is because of
	Clousure.
	In compile time our Go compiler see's that we are returning a inner function and we are using data which are in the parent outer function. So go compiler does
	Escape Analysis and form a closure. It takes the inner function and all the variable or function that has reference inside the inner function GO compiler take those
	and bind them together and places all the variables or funcitons that has reference inside the inner function in the Heap Area.

	So even though the outer function does not exists but still we can access the values which were in outer function from Heap.
*/

func outer(money int) func() int {

	inner := func() int {

		fmt.Println("From inside Closure", money)
		money = money + 1
		return (money + p + a)
	}

	return inner
}

func main() {

	closure := outer(10)

	value := closure()

	fmt.Println(value)

	value = closure()
	fmt.Println(value)

	value = closure()
	fmt.Println(value)

}
