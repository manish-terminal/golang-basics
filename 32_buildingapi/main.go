package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseID    string `json:"courseid"`
	CourseName  string `json:"coursename"`
	CoursePrice int    `json:"courseprice"`
	Author      *Author
}
type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var course []Course

//middleware,helper-file

func isEmpty(c *Course) bool {
	return c.CourseName == ""
}

//controllers-file

// serve home router
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>You have created first api </h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(course)
	return
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")
	//grab id from request
	params := mux.Vars(r)

	fmt.Println(params)
	for _, c := range course {
		if c.CourseID == params["id"] {
			json.NewEncoder(w).Encode(c)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
	return

}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}
	// what about -{}
	var c Course
	_ = json.NewDecoder(r.Body).Decode(&c)
	if isEmpty(&c) {
		w.Write([]byte("<h1>Wrong request</h1>"))
		return
	}
	randomNumber := rand.IntN(100)
	c.CourseID = strconv.Itoa(randomNumber)
	course = append(course, c)
	json.NewEncoder(w).Encode(c)
	return

}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete a course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, c := range course {
		if c.CourseID == params["id"] {
			course = append(course[:i], course[i+1:]...)
			var c Course
			_ = json.NewDecoder(r.Body).Decode(&c)
			c.CourseID = params["id"]
			json.NewEncoder(w).Encode(c)
			return
		}
	}
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r) //
	// loop,id ,remove, add with my id

	for i, c := range course {
		if c.CourseID == params["id"] {
			course = append(course[:i], course[i+1:]...)
			var c Course
			_ = json.NewDecoder(r.Body).Decode(&c)
			c.CourseID = params["id"]
			course = append(course, c)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
}

//challanges we have is generating unique id,string

//append course into courses

func main() {
	fmt.Println("Hello from main function ")
	r := mux.NewRouter()
	course = append(course, Course{CourseID: "2", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{FullName: "Manish Gupta", Website: "manish-terminal"}})
	course = append(course, Course{CourseID: "4", CourseName: "JS", CoursePrice: 299, Author: &Author{FullName: "Manish Gupta", Website: "manish-terminal"}})
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/courses", createOneCourse).Methods("POST")
	r.HandleFunc("/courses/{id}", deleteCourse).Methods("DELETE")
	r.HandleFunc("/courses/{id}", updateOneCourse).Methods("PUT")
	// fmt.Print(rand.IntN(100))
	log.Fatal(http.ListenAndServe(":4000", r))

}
//auto reload like nodemon for nodejs is there any ? yes airgo install -v github.com/cosmtrek/air@latest and air init and air run

