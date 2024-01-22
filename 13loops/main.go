package main

import "fmt"

func main() {
	fmt.Println("Hello")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)
	// [Sunday Tuesday Wednesday Friday Saturday]

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	//Sunday
	//Tuesday
	//Wednesday
	//Friday
	//Saturday

	// another better way

	for i := range days { // here i will be the index
		fmt.Println(days[i])
	}

	// another better way

	for i, day := range days {
		fmt.Printf("index is %v and the value present in this index is %v\n", i, day)
	}
	//index is 0 and the value present in this index is Sunday
	//index is 1 and the value present in this index is Tuesday
	//index is 2 and the value present in this index is Wednesday
	//index is 3 and the value present in this index is Friday
	//index is 4 and the value present in this index is Saturday

	//***************************************************************************************************************************

	rogueValue := 1

	for rogueValue < 10 {

		// if rogueValue == 5 {
		// 	break
		// }

		//>Value is  1
		//>Value is  2
		//>Value is  3
		//>Value is  4

		if rogueValue == 5 {
			rogueValue++
			continue
		}

		//>>Value is  1
		//>>Value is  2
		//>>Value is  3
		//>>Value is  4
		//>>Value is  6
		//>>Value is  7
		//>>Value is  8
		//>>Value is  9

		fmt.Println("Value is ", rogueValue)
		rogueValue++
	}

	//Value is  1
	//Value is  2
	//Value is  3
	//Value is  4
	//Value is  5
	//Value is  6
	//Value is  7
	//Value is  8
	//Value is  9

	// using goto

	i := 1
	for i < 10 {
		if i == 2 {
			goto hello //hello is a label name
		}
		fmt.Println(i)
		i++
	}

hello:
	fmt.Println("hello brother")

}

//1
//hello brother
