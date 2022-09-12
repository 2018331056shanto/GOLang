package main

// The range form of the for loop iterates over a slice or a map.
import (
	"fmt"
	"math"
)
func main(){
	var xx=[]int{1,3,5,6}
	for i,v:=range xx{

		fmt.Println(i,v*5)
	}
	for _,v:=range xx{
		fmt.Println(math.Pow(2,float64(v)))
	}
	fmt.Println("Continue")
	for _,v:=range xx{
		if math.Pow(2,float64(v))<10{
			continue
		}
		fmt.Println(math.Pow(2,float64(v)))

		
	}
	fmt.Println("Break ")
	for _,v:=range xx{
		if math.Pow(2,float64(v))>31{
			break
		}
		fmt.Println(math.Pow(2,float64(v)))

		
	}
	cities := map[string]int{
		"New York":    8336697,
		"Los Angeles": 3857799,
		"Chicago":     2714856,
	}
	for key, value := range cities { //for each key-value pair in cities
		fmt.Printf("%s has %d inhabitants\n", key, value)
	}
}