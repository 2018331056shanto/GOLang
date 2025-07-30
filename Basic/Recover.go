package main

import (
	"fmt"
	"time"
)

func panicRecoverFunction(){

	a,b:=10,0

	fmt.Println("pre panic")
	defer func (){
	
		if r:=recover(); r!=nil{

			fmt.Println(" recorded",r)
		}
	}()
	if b==0 {

		panic("We can not divide by zero")
	}

	fmt.Println(a)
	fmt.Println("Post Panic")

}
func main(){


	go panicRecoverFunction()

	time.Sleep(5*time.Second)


}
