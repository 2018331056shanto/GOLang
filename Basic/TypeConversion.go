package main

import "fmt"
var(
	x int =43
	y float32=344.22
)
func main(){

	z:=uint(x)
	i := 42
	f := float64(i)
	u := uint(f)
	fmt.Printf("%v \n %v \n %v\n",i,f,u)
	fmt.Println(z)
	fmt.Println(int(y))
	fmt.Println(float32(x))
}