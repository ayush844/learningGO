package main

import "fmt"

func main() {
	fmt.Println("hello")

	//there is no concept of inheritance or super or parent in golang

	ayush := User{"ayush", "ayush@gmail.com", true, 20}
	fmt.Println("the person is as foolows:", ayush)
	// the person is as foolows: {ayush ayush@gmail.com true 20}

	// // for structure we should use %+v instead of %v
	fmt.Printf("the person is as follows: %+v\n", ayush)
	// the person is as follows: {Name:ayush Email:ayush@gmail.com Status:true Age:20}

	fmt.Printf("the name of the person is %v and his age is %v\n", ayush.Name, ayush.Age)
	// the name of the person is ayush and his age is 20
}

//*********************************** MAKING THE STRUCTURE ***********************************

// here the name of our structure and also the name of the elements inside it are starting with the capital letter because they can can be accessed from anywhere in this file so they are public
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
