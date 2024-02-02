package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)


func main()  {
	fmt.Println("welcome to web verb video")

	performGetRequest()
}

func performGetRequest()  {
	const myUrl = "http://localhost:8000/get"
	response, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("StatusCode: ", response.StatusCode)
	// StatusCode:  200

	fmt.Println("Content length is: ", response.ContentLength)
	// Content length is:  43


	// way 1:
	// content, _ := io.ReadAll(response.Body)
	// fmt.Println(string(content))
	// // {"message":"Hello from learnCodeonline.in"}


	// way 2:
	var responceString strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responceString.Write(content)

	fmt.Println("ByteCount is: ", byteCount)
	// ByteCount is:  43

	fmt.Println(responceString.String())
	// {"message":"Hello from learnCodeonline.in"}


}