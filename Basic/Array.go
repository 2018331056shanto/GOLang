package main
import(
	"fmt"
)
// The type [n]T is an array of n values of type T.
// var a [10]int
func main(){
	var a [2]string   //array of size 2
    a[0] = "Hello"    //Zero index of "a" has "Hello"
    a[1] = "World"    //1st index of "a" has "World"
    fmt.Println(a[0], a[1]) // will print Hello World
    fmt.Println(a)    
	//  using Printf and used the %q “verb” to print each element quoted.
	fmt.Printf("%q\n",a)
	primes:=[6]int {2,3,5,7,9}
	fmt.Println(primes)
	// fmt.Printf("%q\n",primes)
	
	// Finally, you can use an ellipsis to use an implicit length when you pass the values:
	// the size of the array is dynamic like vector in C++
	b:=[...]string{"a","b","c","d","e"}
	fmt.Printf("%q\n",b)
	// Multi Dimensional Arrya
	 y:=[...][3] int {{1,2,4},{3,4,2}}
	fmt.Println(y)	
}