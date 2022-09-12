// Go supports the "new" expression to allocate a zeroed 
// value of the requested type and to return a pointer to it.

// new 
// 1) returns a pointer 2)newly allocated 3)Zerod value
// make
// 1) slice, maps, channels 2) allocated memory 3)initializes(puts 0 r empty string int values)
package main
import (
	"fmt"
)
func main(){
	var age *[]int= new([]int)// returns a pointeer
	fmt.Println(age)
	fmt.Println(*age)
	var bday []int =make([]int,10)
	fmt.Println(bday)
	fmt.Println(len(bday))

}