package main
import (
	"fmt"
	"math"
)


func print(x int, y int)  {
 z:=x + y
 fmt.Println(z)
}
func sqrt(x,y int)  (float64,float64) {
	p:=math.Sqrt(float64(x))
	q:=math.Sqrt(float64(y))
	return  p,q;
}
func location(city string) (region, continent string) {
	switch city {
	case "Los Angeles", "LA", "Santa Monica":
		region, continent = "California", "North America"
	case "New York", "NYC":
		region, continent = "New York", "North America"
	default:
		region, continent = "Unknown", "Unknown"
	}
	return //returning region and continent
}
type Vertex struct {
	X, Y float64
}

func (v Vertex) Math()float64{
	 return math.Sqrt(v.X*v.X + v.Y*v.Y)

}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Named return values 
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
func main(){
	print(666,888)
	p,q:=sqrt(44,55)
	fmt.Printf("%f \n %f\n",p,q)
	region, continent := location("Santa Monica")
	fmt.Printf("Matt lives in %s, %s\n", region, continent)
	v := Vertex{2, 4}
	
	v.Scale(10)
	fmt.Println(v)
	fmt.Println(v.Math())
	fmt.Println(split(17))
}