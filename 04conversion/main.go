package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	// "strings"
)

func main() {
	fmt.Println("welcome to my pizza shop")
	fmt.Println("please rate us between 1 and 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for the rating, ", input)

	// numRating, err := strconv.ParseFloat(input, 64) ====> strconv.ParseFloat: parsing "3\n": invalid syntax

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("added 1 to your rating, ", numRating+1)
	}

	fmt.Printf("type of input is %T\n", input)
	// type of input is string

	fmt.Printf("type of numRating is %T\n", numRating)
	// type of numRating is float64
}
