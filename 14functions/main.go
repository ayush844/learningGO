package main

import "fmt"

// main act as an entry point for our program
// we can not declare a function inside another function

func main() {
	fmt.Println("welcome to functions in golang")
	greeter()
	greeter2()

	result := adder(3, 5)
	fmt.Println("the result is ", result)
	// the result is  8

	proResult := proAdder(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13)
	fmt.Println("the pro result is ", proResult)
	// the pro result is  91

	ultraProResult, message := ultraProAdder(123, 234, 345, 456)
	fmt.Println("the result is ", ultraProResult)
	fmt.Println("the message is ", message)
	//the result is  1158
	//the message is  calculation done perfectly

}

func greeter() {
	fmt.Println("Namaste from Golang")
}

func greeter2() {
	fmt.Println("hello from Golang")
}

func adder(a int, b int) int {
	return a + b
}

//if we are not sure about the number of arguments

func proAdder(values ...int) int {
	total := 0

	for _, value := range values {
		total += value
	}

	return total
}

//if we want to return more than one value

func ultraProAdder(values ...int) (int, string) {
	total := 0

	for _, value := range values {
		total += value
	}

	return total, "calculation done perfectly"

}
