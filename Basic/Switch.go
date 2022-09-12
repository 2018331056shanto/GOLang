package main

import (
	"fmt"
	"time"
)

// You can only compare values of the same type.

// You can set an optional default statement to be executed if all the others fail.

// You can use an expression in the case statement, for instance you can calculate a value to use in the case:

func main(){
	now:=time.Now().Unix()
	mins:=now%2
	

	switch mins{
	case 0:
		fmt.Println("Even")
	case 1:
		fmt.Println("Odd")
	}
	// You can have multiple values in a case statement:
		score := 7
	switch score {
	case 0, 1, 3:
		fmt.Println("Terrible")
	case 4, 5:
		fmt.Println("Mediocre")
	case 6, 7:
		fmt.Println("Not bad")
	case 8, 9:
		fmt.Println("Almost perfect")
	case 10:
		fmt.Println("hmm did you cheat?")
	default: //to be executed if no other case matches
		fmt.Println(score, " off the chart")
	}

	// You can execute all the following statements after a match using the fallthrough statement
	n:=3
	switch n{
	case 1:
		fmt.Println("One")
		fallthrough
	case 3:
		fmt.Println("Three")
		fallthrough
	case 4:
		fmt.Println("Four")
	default:
		fmt.Println("Not found")
	}

}