package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//model for course - file
type Course struct{
	CourseId string `json:"courseid"`
	CourseName string `json:"coursename"`
	CoursePrice int `json:"price"`
	Author *Author `json:"author"`
}

type Author struct{
	Fullname string `json:"fullname"`
	Website string `json:"website"`
}


//fakeDB
var courses []Course

//middleware or helper methods

func (c *Course) isEmpty() bool {
	return c.CourseId=="" && c.CourseName=="" 	//ensures that both fields don't stay empty
}


func main()  {
	
}


// controllers - file


//serve home route
func serveHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>WELCOME TO API BY AYUSH SHARMA</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}