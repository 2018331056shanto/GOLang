package main
import (
	"fmt"
)

func main(){

	sum:=0
	for i:=0;i<10;i++{
		sum+=i
	}
	fmt.Println(sum)
	// The init and post statements are optional.
	x := 1
	for  ;x< 1000; {
		fmt.Println(x)
		x+= x
	}
	fmt.Println(x)
	// For is Go's "while"
	y:=1
	for y<100{
		y+=y
	}
	fmt.Println(y)


}
