package main

import (
	"encoding/json"
	"fmt"
)

type course struct{
	Name string	`json:"coursename"`
	Price int
	Platform string `json:"website"`
	Password string `json:"-"` // "-" it tells that we don't want to show this value in json
	Tage []string `json:"tags,omitempty"` //omitempty tells that if any field is nil then don't include it
}

func main()  {
	fmt.Println("more about json")

	EncodeJson()
}

//encoding json [converting data to json]

func EncodeJson()  {
	// making slices
	lcoCourses := []course{
		{"Reactjs Bootcamp", 299, "lco.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "lco.in", "def123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "lco.in", "ghi123", nil},
	}

	// packing this data as JSON data

	//----------------------------------------------------------------------------------------------------------------

	//way1:

	// finalJson, err := json.Marshal(lcoCourses)

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("%s\n", finalJson)

	//[{"Name":"Reactjs Bootcamp","Price":299,"Platform":"lco.in","Password":"abc123","Tage":["web-dev","js"]},{"Name":"MERN Bootcamp","Price":199,"Platform":"lco.in","Password":"def123","Tage":["full-stack","js"]},{"Name":"Angular Bootcamp","Price":299,"Platform":"lco.in","Password":"ghi123","Tage":null}]

	//----------------------------------------------------------------------------------------------------------------

	//way2:

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")	//"" is the prefix && "\t" tells that tab is the place to 	seperate

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
	//[
    //    {
	//		"Name": "Reactjs Bootcamp",
	//		"Price": 299,
	//		"Platform": "lco.in",
	//		"Password": "abc123",
	//		"Tage": [
	//				"web-dev",
	//				"js"
	//		]
	//	},
	//	{
	//		"Name": "MERN Bootcamp",
	//		"Price": 199,
	//		"Platform": "lco.in",
	//		"Password": "def123",
	//		"Tage": [
	//				"full-stack",
	//				"js"
	//		]
	//	},
	//	{
	//		"Name": "Angular Bootcamp",
	//		"Price": 299,
	//		"Platform": "lco.in",
	//		"Password": "ghi123",
	//		"Tage": null
	//	}
	//]

	//----------------------------------------------------------------------------------------------------------------
		// after updating the struct:
		// type course struct{
			//Name string	`json:"coursename"`
			//Price int
			//Platform string `json:"website"`
			//Password string `json:"-"`
			//Tage []string `json:"tags,omitempty"`
		// }

	//[
    //    {
	//		"coursename": "Reactjs Bootcamp",
	//		"Price": 299,
	//		"website": "lco.in",
	//		"tags": [
	//				"web-dev",
	//				"js"
	//		]
	//	},
	//	{
	//		"coursename": "MERN Bootcamp",
	//		"Price": 199,
	//		"website": "lco.in",
	//		"tags": [
	//				"full-stack",
	//				"js"
	//		]
	//	},
	//	{
	//		"coursename": "Angular Bootcamp",
	//		"Price": 299,
	//		"website": "lco.in"
	//	}
	//]
	




}