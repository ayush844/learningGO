package main

import "fmt"

// sometimes when we pass variables to a function instead of our input a copy of our input is passed to avoid this we can use pointers, pointers help us to pass the actual memory address or the reference to the variables to functions

func main() {
	fmt.Println("hello")

	// var ptr *int
	// fmt.Println("Default value of pointer is: ", ptr)
	// // Default value of pointer is:  <nil>

	//************************************************************************************************

	myNumber := 69
	var myPtr = &myNumber // & -> for referencing

	fmt.Println("actual memory address of myNumber is", myPtr)
	// actual memory address of myNumber is 0xc000012188

	fmt.Println("value present at the address that is stored inside our pointer:", *myPtr)
	// value present at the address that is stored inside our pointer: 69

	*myPtr = *myPtr * 2
	fmt.Println("our new value is", myNumber)
	//our new value is 138

}
