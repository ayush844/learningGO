package main

import "fmt"

// defer is a mechanism in Go that allows you to delay the execution of a function until the surrounding function completes.
// deferred functions are invoked immediately before the surrounding function returns, in the reverse order they were referred i.e. it is following LIFO (last in first out)

func main() {

	// // 1)

	// 	defer fmt.Println("world")
	// 	fmt.Println("Hello")

	// 	//Hello
	// 	//world

	//**********************************************************************************************************

	// // 2)

	// defer fmt.Println("world")
	// defer fmt.Println("One")
	// defer fmt.Println("Two")
	// fmt.Println("Hello")

	// //Hello
	// //Two
	// //One
	// //world

	//**********************************************************************************************************

	// 3)

	defer fmt.Println("world")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	myDefer()

	//Hello
	//43210Two
	//One
	//world

}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
