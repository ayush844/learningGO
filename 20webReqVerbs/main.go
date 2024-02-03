package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)


func main()  {
	fmt.Println("welcome to web verb video")

	// performGetRequest()

	// performPostJsonRequests()

	performPostFormRequest()
}


// making a get requests
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


// making a post requests
func performPostJsonRequests()  {
	const myUrl = "http://localhost:8000/post"

	//fake json payload
	requestBody := strings.NewReader(`
		{
			"coursename": "Lets go with golang",
			"price": 0,
			"platform": "lco.in"
		}
	`)

	response, err :=http.Post(myUrl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))

	// {"coursename":"Lets go with golang","price":0,"platform":"lco.in"}

}


func performPostFormRequest()  {
	const myUrl = "http://localhost:8000/postform"

	//formdata

	data := url.Values{}
	data.Add("firstName", "Ayush")
	data.Add("lastname","Sharma")
	data.Add("email", "ayush.xyz1625@gmail.com")

	response, err := http.PostForm(myUrl, data)
	
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
	// {"email":"ayush.xyz1625@gmail.com","firstName":"Ayush","lastname":"Sharma"}



}