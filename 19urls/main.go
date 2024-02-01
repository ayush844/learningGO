package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main()  {
	fmt.Println("Welcome to handling URLs in golang")

	fmt.Println(myurl)
	// https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb

	// PARSING URL
	result, _ := url.Parse(myurl)

	fmt.Println(result)
	// https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb

	fmt.Println(result.Scheme) // The URL scheme is the part of the URL that indicates the protocol or type of communication to be used.
	// https

	fmt.Println(result.Host)

	//lco.dev:3000

	fmt.Println(result.Path)
	// /learn

	fmt.Println(result.Port())
	//3000

	fmt.Println(result.RawQuery)
	// coursename=reactjs&paymentid=ghbj456ghb


	qparams := result.Query()
	fmt.Println(qparams)
	// map[coursename:[reactjs] paymentid:[ghbj456ghb]]
	fmt.Printf("the type of query params are: %T\n", qparams)
	// the type of query params are: url.Values

	fmt.Println(qparams["coursename"])
	// [reactjs]

	for _, value := range qparams{
		fmt.Println("Param is: ", value)
	}
	//Param is:  [reactjs]
	//Param is:  [ghbj456ghb]

//************************************************************************************************************************

	// constructing the url

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawQuery: "user=ayush",
	}

	anotherUrl := partsOfUrl.String()

	fmt.Println(anotherUrl)
	// https://lco.dev/tutcss?user=ayush


}