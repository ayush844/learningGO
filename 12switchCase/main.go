package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("hello")

	diceNumber := rand.Intn(6) + 1
	fmt.Println("value on dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("dice value is 1 and you can open")
	case 2:
		fmt.Println("you can move 2 spots")
	case 3:
		fmt.Println("you can move 3 spots")
	case 4:
		fmt.Println("you can move 4 spots")
		fallthrough // if we will hit case 4 then we'll also show the result of cse 5
	case 5:
		fmt.Println("you can move 5 spots")
	case 6:
		fmt.Println("you can move 6 spots and roll dice again")
	default:
		fmt.Println("What the hell was that!!!")

	}

	//for case 4:
	////>hello
	////>value on dice is  4
	////>you can move 4 spots
	////>you can move 5 spots

}
