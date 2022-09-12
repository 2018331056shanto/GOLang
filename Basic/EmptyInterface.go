// An empty interface may hold values of any type
// An empty interface  are used by code that handles values of unknown type
// and you can pass an empty interface type as a function parameter of any type
package main
import(
	"fmt"
)
type emptyInterface interface{

	// an empty interface can hold value of any type
}
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
type Person struct{
	info interface{}
}
func main(){
	var i interface{}
	describe(i)
	i=24
	describe(i)
	i="hello"
	describe(i)
	var empty interface{}
	empty=5
	fmt.Println(empty)
	empty="hello world"
	fmt.Println(empty)
	empty=[] int{1,2,4,5}
	fmt.Println(empty)
	fmt.Println(len(empty.([]int)))//here type assertion has done

	p:=Person{}
	p.info="I am Shanto"
	fmt.Println(p.info)
	p.info=33
	fmt.Println(p.info)
	p.info=[]float64{3.3,5,22}
	fmt.Println(p.info)




}