package main

import "fmt"

func main() {
	fmt.Println("hello")

	ayush := User{"ayush", "ayush@gmail.com", true, 20}
	fmt.Println("the person is as foolows:", ayush)
	//the person is as foolows: {ayush ayush@gmail.com true 20}

	fmt.Printf("the person is as follows: %+v\n", ayush)
	//the person is as follows: {Name:ayush Email:ayush@gmail.com Status:true Age:20}


	fmt.Printf("the name of the person is %v and his age is %v\n", ayush.Name, ayush.Age)
	//the name of the person is ayush and his age is 20

	ayush.GetStatus()
	// Is user active:  true

	ayush.NewMail()
	// Email of this user is  test@go.dev

	fmt.Println(ayush) // =======>>>>> this is only changing the copy and don't have any effect on original value
	// {ayush ayush@gmail.com true 20}

	// if we want to chnage the original value we should pass the pointer to the object in the method as an arguement

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

//methods are generally associated with a specific type (here it is User)

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is ", u.Email)
}

// functions are standalone blocks of code, while methods are functions associated with a specific type, often used to define behaviors related to that type.
