package main

import "fmt"

func main() {
	fmt.Println("hello")

	var fruitList [4]string
	fruitList[0] = "apple"
	fruitList[1] = "tomato"
	fruitList[3] = "peach"

	fmt.Println("our fruitlist is ", fruitList)
	//our fruitlisty is  [apple tomato  peach] (extra space between tomato and peach showing that index 2 is not filled)

	fmt.Println("length of our fruitlisty is ", len(fruitList))
	//length of our fruitlist is  4 (it tell us the size of array not the total nuber of element stored inside it)

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("vegylist is ", vegList)
	//vegylist is  [potato beans mushroom]

}
