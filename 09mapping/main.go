package main

import (
	"fmt"
)

func main() {

	fmt.Println("hello")

	language := make(map[string]string)
	language["js"] = "javascript"
	language["rb"] = "ruby"
	language["py"] = "python"

	fmt.Println("list of all languages: ", language)
	// list of all languages:  map[js:javascript py:python rb:ruby]

	fmt.Println("JS shorts for: ", language["js"])
	// JS shorts for:  javascript

	delete(language, "rb") // to delete a value from the map

	fmt.Println("list of all languages: ", language)
	// list of all languages:  map[js:javascript py:python]

	// have a look at loops in maps
	for key, value := range language {
		fmt.Printf("for key %v, we have a value of %v\n", key, value)
	}
	//for key py, we have a value of python
	//for key js, we have a value of javascript

	// // here if we don't want key:
	// for _, value := range language {
	// 	fmt.Printf("we have a value of %v\n", value)
	// }
	// //we have a value of javascript
	// //we have a value of python

}
