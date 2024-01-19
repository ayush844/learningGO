package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("hello")

	//we don't have to tell the size here and we need to initialize it in starting
	var fruitList = []string{}
	fmt.Printf("type of fruitlist is %T\n", fruitList)
	//type of fruitlist is []string

	var newFruitList = []string{"apple", "tomato", "peach"}
	fmt.Printf("type of fruitlist is %T\n", newFruitList)

	//-------------------------------we can also add new values in the slice-------------------------------

	newFruitList = append(newFruitList, "mango", "banana")
	fmt.Println(newFruitList)
	// [apple tomato peach mango banana]

	// newFruitList = append(newFruitList[1:])
	// fmt.Println(newFruitList)
	// // [tomato peach mango banana]

	// newFruitList = append(newFruitList[:4])
	// fmt.Println(newFruitList)
	// // [apple tomato peach mango]

	//include 1rst element and go till before the last element
	// newFruitList = append(newFruitList[1:4])
	// fmt.Println(newFruitList)
	// // [tomato peach mango]

	highScore := make([]int, 4)
	highScore[0] = 193
	highScore[1] = 934
	highScore[2] = 385
	highScore[3] = 276

	// highScore[4] = 361 ==> /error/ panic: runtime error: index out of range [4] with length 4

	fmt.Println(highScore)
	// [193 934 385 276]

	highScore = append(highScore, 111, 555, 999) //as we append more values the entire reallocation of memory take place
	fmt.Println(highScore)
	// [193 934 385 276 111 555 999]

	//-------------------------------sorting the slice and other operations-------------------------------

	sort.Ints(highScore) // sorts the entire slice in increasing order
	fmt.Println(highScore)
	// [111 193 276 385 555 934 999]

	fmt.Println(sort.IntsAreSorted(highScore)) //it tells if our slice is sorted or not
	// true

	// How to remove a value from slice based on index in golang
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	// [reactjs javascript swift python ruby]

	var index = 2 // we have to remove the index 2 from courses

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
	// [reactjs javascript python ruby]

}
