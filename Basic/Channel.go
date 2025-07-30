package main

import "fmt"

// named return valued function
func sliceToChannel(nums []int) <-chan int {

	out := make(chan int)

	go func() {

		for n := range nums {
			out <- n
		}
		close(out)
	}()

	return out
}

func convertToSquare(in <-chan int) <-chan int {

	do := make(chan int)

	go func() {

		for n := range in {

			do <- n * n
		}
		close(do)

	}()

	return do

}

func main() {

	// channel := make(chan int)

	// go func() {
	// 	channel <- 100
	// 	close(channel)
	// }()

	// // channel <- 100
	// // close(channel)

	// msg := <-channel

	// fmt.Println(msg)

	num := []int{1, 2, 3, 4, 5, 6, 7, 8}

	myChannel := sliceToChannel(num)

	anotherChannel := convertToSquare(myChannel)
	for n := range anotherChannel {
		fmt.Println(n)
	}

	// select{

	// case myChannel:
	// 	for n:=range myChannel{
	// 		fmt.Println(n)
	// 	}
	// }

}
