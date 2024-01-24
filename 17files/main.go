package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("hello")
	content := "this needs to go inside a file"

	file, err := os.Create("./myfile.txt")

	if err != nil {
		panic(err) // panic will shut down execution of program and show us the error
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println("length is ", length)

	defer file.Close()

	readFile("/home/ayush/Desktop/goWithGolang/17files/myfile.txt")

}

// to read file

func readFile(fileName string) {

	dataBytes, err := os.ReadFile(fileName)
	// here the data which we will read will not be read in string, it will be read in bytes

	checkNilErr(err)

	fmt.Println("Text data inside the file is : \n", dataBytes)

	//Text data inside the file is :
	// [116 104 105 115 32 110 101 101 100 115 32 116 111 32 103 111 32 105 110 115 105 100 101 32 97 32 102 105 108 101]

	fmt.Println("Text data inside the file is : \n", string(dataBytes))
	//Text data inside the file is :
	// this needs to go inside a file

}

//if we don't want to check for err every time:

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
