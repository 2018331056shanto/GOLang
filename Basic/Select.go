package main

import "fmt"

/*
	Select lets a go routine wait on multiple communication operations
	In this case it's going to let our main function wait on messages from
	multiple channels
	A select statment is going to block all of the channels until one of it's cases run
	If Select is able to receive message from multiple channels at the same time
	in that case it randomly choose one channel

*/

func main() {

	myChannel := make(chan int)
	anotherChannel := make(chan string)

	go func() {
		anotherChannel <- "shanto"
	}()

	go func() {
		myChannel <- 10000
	}()

	// msg:=<- myChannel

	select {

	case msgFromMyChannel := <-myChannel:
		fmt.Println(msgFromMyChannel)

	case msgFromAnotherChannel := <-anotherChannel:
		fmt.Println(msgFromAnotherChannel)

	// default:
	// 	fmt.Println("goooooooooooo")
	}

}
