package main

import (
	"fmt"
)

func hello(st string) {
	fmt.Println(st)
	
}
func main() {

	go hello("hello")
	fmt.Println("Hello how are you")

}
