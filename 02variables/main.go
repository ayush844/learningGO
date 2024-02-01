package main

import "fmt"

const LoginToken string = "abcd"

// here name "LoginToken" have first letter capital  to show that it's a public variable

func main() {
	var username string = "ayush"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n", username)
	//ayush
	//variable is of type: string

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type: %T \n", isLoggedIn)
	//true
	//variable is of type: bool

	var smallValue uint = 255
	fmt.Println(smallValue)
	fmt.Printf("variable is of type: %T \n", smallValue)
	//255
	//variable is of type: uint

	var smallFloat float32 = 255.435367855775434577789764
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type: %T \n", smallFloat)
	//255.43536
	//variable is of type: float32

	//default variable and some aliases

	var variable int
	fmt.Println(variable)
	fmt.Printf("variable is of type: %T \n", variable)
	//default int is 0

	var variable02 string
	fmt.Println(variable02)
	fmt.Printf("variable is of type: %T \n", variable02)
	//default string is ""

	//impilicit way of declaring variables
	var website = "learncodeonline.in"
	fmt.Println(website)
	fmt.Printf("variable is of type: %T \n", website)
	//here lexer make this variuable string and it will stay as string unlike in js here we can not change it to int

	// no var style <>walrus operator<>(it is only allowed to use inside methods, should not be global)
	numberOfUser := 300000
	fmt.Println(numberOfUser)
	fmt.Printf("variable is of type: %T \n", numberOfUser)
	//300000
	//variable is of type: int

	fmt.Println(LoginToken)
	fmt.Printf("variable is of type: %T \n", LoginToken)
	//abcd
	//variable is of type: string

}
