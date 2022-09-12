// A type assertion takes a value and tries to create
//  another version in the specified explicit type
package main
// type assertion provides access to an interface value's underlying concrete type
// go has no Generic like Java so type assertion is used for one of these ways 
// when we use an interface and we want to get the actual value of the sturct insted of the interface method we use type assertion
import(
	"fmt"
	"math"
)
type Shape1 interface{
	 area1() float64//returns value float 
	 perimeter1() float64//returns value float
}
type Rectangle1 struct{
	
	//struct is like Class in java and Struct in C/C++
	width,height float64
	sides int

}
type Polygon interface{
	getSisde() int
}
type Circel1 struct{
	radius float64
}
func(r Rectangle1)area1() float64{
	return r.height*r.width
}
func(c Circel1) area1() float64{
	return math.Pi *c.radius*c.radius
}

func (r Rectangle1) perimeter1() float64{
	return 2*r.height+2*r.width
}
func (r Rectangle1) getSisde() int{
	return r.sides
}
// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func(c Circel1) perimeter1() float64{
	return 2*math.Pi*c.radius
}

func Calculate1(s Shape1){

	// fmt.Println("My type %T\n",s)
	// fmt.Println("My value %v\n",s)
	fmt.Println(s.area1())
	fmt.Println(s.perimeter1())
	rect:=s.(Polygon) //or rect:=s.(Rectangle1)
	//type assertion 
	// in interface Shape1 we have two method area1()
	// and perimeter1() so if we want to pass any object instence we have
	// to first ensure that our instance hase these method implemented if we have
	// extra method or parameter in that instance and we want to access in those methods and parameters
	// then we have to do type assertion which underlying value we want to access
	fmt.Println("TYpe %T\n",rect)
	fmt.Println(rect.getSisde())
}

func main(){
	r:=Rectangle1{height: 10,width: 15,sides: 4}
	// c:=Circel1{radius: 55}
	Calculate1(r)
	// Calculate1(c)
}

