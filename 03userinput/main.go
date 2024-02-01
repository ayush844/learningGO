package main

import (
	"bufio"
	"fmt"
	"os"
)

// Comma ok syntax and packages in golang
func main() {
	message := "welcome brother"
	fmt.Println(message)

	//bufio and os syntax
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the rating for our pizza: ")

	//------------////------------// comma ok syntax || comma err syntax //------------////------------//

	input, _ := reader.ReadString('\n')

	// input, err := reader.ReadString('\n')

	// _, err := reader.ReadString('\n')

	fmt.Println("Thanks for rating, ", input) // Thanks for rating,  3

	//every input is taken as strings

	fmt.Printf("type of rating is: %T \n", input) // type of rating is: string

}
