package main

import "fmt"

func main() {
	fmt.Println("hello")

	loginCount := 25

	var message string

	if loginCount < 10 {
		message = "regular user"
	} else if loginCount > 10 {
		message = " watch out"
	} else {
		message = "login count reached to 10"
	}

	fmt.Println("message:", message)
	// message:  watch out

	//-------we don't need to create the variable in advance to use in the if else statement we can also create it inside it but this variable can not be used outside if else statemnt

	if num := 3; num < 10 {
		fmt.Println("number is less than 10")
	} else {
		fmt.Println("number is not less than 10")
	}

	// if err != nil {
	// }

}
