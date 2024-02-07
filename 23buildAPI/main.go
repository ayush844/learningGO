package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
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
	// return c.CourseId=="" && c.CourseName=="" 	//ensures that both fields don't stay empty
	return c.CourseName=="" 
}


func main()  {
	
}


// controllers - file


//serve home route
func serveHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>WELCOME TO API BY AYUSH SHARMA</h1>"))
}

//get all the courses
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}


// get a particular course
func getOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)
	fmt.Println(params)

	//loop through courses find matching id and return the response

	for _, course := range courses{
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("no course found of the given id")

	return

}


//create a new course
func createOneCourse(w http.ResponseWriter, r *http.Request)  {

	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// what if - body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// what about - {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.isEmpty() {
		json.NewEncoder(w).Encode("no data inside JSON")
		return
	}


	// generate unique id, string
	// append course into courses


	seed := time.Now().Unix() // Or any other seed value
	source := rand.NewSource(seed)
	rng := rand.New(source)

// Generate random numbers using rng
	randomNumber := rng.Intn(100) // Generate a random number using the current Unix timestamp in nanoseconds

	course.CourseId = strconv.Itoa(randomNumber)


	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)

	return



}