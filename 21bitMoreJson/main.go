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

	// EncodeJson()

	DecodeJson()
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

func DecodeJson()  {
	// whenever any data comes from web it is in byte format
	jsonDataFromWeb := []byte(`
		{
		"coursename": "Reactjs Bootcamp",
		"Price": 299,
		"website": "lco.in",
		"tags": ["web-dev","js"]
		}
	`)

	var lcoCourse course 

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse) //here #v will be used
	}else{
		fmt.Println("JSON was not valid")
	}

	// JSON was valid
	// main.course{Name:"Reactjs Bootcamp", Price:299, Platform:"lco.in", Password:"", Tage:[]string{"web-dev", "js"}}


//----------------------------------------------------------------------------------------------------------------------

	// some cases where we just want to get data as key value pair without using predefined struct

	// since we are not sure about type of data(value in key value pair) we recieve therefore we'll use interface here
	//an interface type is defined as a set of method signatures

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	//map[string]interface {}{"Price":299, "coursename":"Reactjs Bootcamp", "tags":[]interface {}{"web-dev", "js"}, "website":"lco.in"}

	for k, v := range myOnlineData{
		fmt.Printf("the key is: %v, the value is: %v and type of value is: %T\n", k, v, v)
	}

	//the key is: coursename, the value is: Reactjs Bootcamp and type of value is: string
	//the key is: Price, the value is: 299 and type of value is: float64
	//the key is: website, the value is: lco.in and type of value is: string
	//the key is: tags, the value is: [web-dev js] and type of value is: []interface {}

	

}