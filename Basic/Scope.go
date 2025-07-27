package main

import (
	"fmt"

	"example.com/mathlib"
)

// In Go, we can declare variables outside any function — at package level.
// This is unlike Java (where everything is inside a class) and similar to JavaScript (where top-level variables are allowed).

var a, b = 20, 30

// func add(x int, y int) {
// 	z := x + y
// 	// l := x + p // ❌ Error: 'p' is undefined in this scope
// 	fmt.Println(z)
// }

func main() {
	// p := 30
	// q := 40
	// add(p, q) // 70
	// add(a, b) // 50
	// add(a, p) // 50

	// // add(a, z) // ❌ Error: 'z' is not in scope

	// x := 18

	// if x >= 18 {
	// 	l := 20
	// 	fmt.Println("You are eligible to vote")
	// 	fmt.Printf("I will cast %d vote\n", l)
	// } else {
	// 	fmt.Println("You are not eligible to vote")
	// }
	// // fmt.Println(l) // ❌ Error: 'l' is not defined in this scope

	mathlib.Add(4, 7)      // 50
	fmt.Println(mathlib.C) // Hello, World!
}

/*

📌 Scope in Go:

- **Package-level variables** (like a and b) are accessible from any function in the same package.
- **Function-level (local) variables** are only accessible within the function they are declared in (e.g., z, p, q).
- **Block-level variables** (declared inside if, for, etc.) are limited to that block.

🧠 Execution Model (compared to JavaScript):

Go is a compiled language, not interpreted like JavaScript. So the Go compiler does **not** behave exactly like the JS engine's hoisting or runtime memory model.

But conceptually for learning:

1. Variables like `a`, `b`, and functions like `add()` and `main()` are registered in memory at compile time.
2. Execution starts with the `main()` function — it is the **entry point** of any Go program.
3. When `main()` is called:
   - A new stack frame is created for it.
   - Local variables `p` and `q` are allocated on the stack.
4. When `add(p, q)` is called:
   - First the `add()` function is looked up in the stack frame of `main()`. And if it is not found, it checks the Global Memory.
   - The `add()` function is executed:
   - A new stack frame is created for `add()`.
   - Parameters `x` and `y` are assigned the values of `p` and `q`.
   - Variable `z` is computed and printed.
   - After `add()` finishes, control returns to `main()`.
5. After the control returns to `main()`, the stack frame for `add()` is popped off, and execution continues in `main()`.
6. The same process happens for `add(a, b)` and `add(a, p)`.
	- A new Stack frame is created for each call.
	- The parameters `x` and `y` are assigned the values of `a`, `b`, or `p`.
	- It checks for the variable in the local scope first if it's not found, it checks the Global Memory.
	- Then the addition is performed and assigned to `z`, which is printed.
	- After the function execution, the stack frame is popped off.
After all function calls, the `main()` function completes, and the program exits.

In case of `add(a, z)`, it results in an error because `z` is not defined in the scope of `main()`. The Go compiler checks the local scope first, then the package scope, and finally the global scope for variable definitions.
- This is similar to how JavaScript works, but in Go, the stack frames are managed at compile time, and the execution is strictly linear without the dynamic nature of JavaScript's runtime environment
*/

/*

There are 3 types of scopes in Go:
1. Global Scope
2. Package Scope
	-
3. Local Scope
	- Block Scope
		Anything inside `{}`is like the Block Scope in Go. Any variable declared inside a Block is called the Local Variable.
		In main function we have defined a Block. So any variable declared inside the main function is a Local Variable.
		In compilation time compiler allocate memory for global variable, functions in the Global Area of RAM.

		But when we declare a variable inside a function like main() a new Stack Frame is created and a new memory localtion is allocated for that Stack Frame.
		So any variable declared inside the main() function will not get any allocated memory in the Global Area of RAM instaed it will get allocated in the Stack Frame of main() function
		and this is Local Scope.
		Now in the above case when if condition is true then i am assigning a new variable `p` inside the if block.

		Now here comes Scope when the whole file will run, the go runtime will allocate memory for all the globar variables and functions in the Global Area of RAM.

		So the initialization is completed and now the main function will run. When the main() function will run it will create a new Stack Frame and it will get allocated memory in the Stack Area of RAM.
		Anything inside the main() is called Local Scope.

		Now the variable `l` is declared inside the if block so it will be allocated memory in the Stack Frame of main() function.
		To get the condition evaluated it needs to check the value of `x` . So it will first check the value of `x` in the Stack Frame of main() function.
		If it is not found then it will check the Global Area of RAM.
		So the value of `x` is found in the Stack Frame of main() function and it is evaluated to true.

		Now the program will enter inside the if block and an new variable `l` is declared here.
		But again it's a Scope as the variable `l` is declared inside the Block Scope of the if statement. At this moment a new Scope will get created
		inside the main which is already a Local Scope. So we are creating a Local Scope inside a Local Scope.

		So Go sees `l` as a scoped only to the Inner Block. `l` is still allocated memory in the same Stack Frame of the main() function but
		it's lifetime is limited by the Inner Block only.
		When the Block exits `l` goes out of Scope and it's memory is eligible to be overwritten
		Now when the variable 'l' is going to be printed it will first look for the variable `l`
		in the Inner Block Scope. If it was not found in the Inner Block Scope then it would have checked the
		Local Scope of the main() function. If it was not found in the Local Scope of the main() function then it would have checked the Global Area of RAM
		and again if it was not found in the Global Scope then it would have thrown an error.

		In this case  variable `l` is found in the Inner Block Scope and it is printed.






*/
