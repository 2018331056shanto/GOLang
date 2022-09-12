package main
// an interface stores 2 values
// 1.Dynamic Type 2. Dynamic Concrete Value
// to access dynamic value we have to do an assertion
import(
	"fmt"
	"math"
)
type Shape interface{
	 area() float64//returns value float 
	 perimeter() float64//returns value float
}
type Rectangle struct{
	
	//struct is like Class in java and Struct in C/C++
	width,height float64

}
type Circel struct{
	radius float64
}
func(r Rectangle)area() float64{
	return r.height*r.width
}
func(c Circel) area() float64{
	return math.Pi *c.radius*c.radius
}

func (r Rectangle) perimeter() float64{
	return 2*r.height+2*r.width
}
// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func(c Circel) perimeter() float64{
	return 2*math.Pi*c.radius
}
func Calculate(s Shape){
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
}

func main(){
	r:=Rectangle{height: 10,width: 15}
	c:=Circel{radius: 55}
	Calculate(r)
	Calculate(c)
}

