package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"platform"`
	Password string   `json:"-"` //this field wont be exported in json
	Tags     []string `json:"tags"`
}

func main() {
	fmt.Println("Welcome to JSON video")
	// EncodeJSON()
	DecodeJSON()

}

func EncodeJSON() {

	lcoCourses := []course{
		{
			Name:     "ReactJS BootCamp",
			Price:    299,
			Platform: "LearnCodeOnline",
			Password: "abc123",
			Tags:     []string{"web-dev", "js"},
		},
		{
			Name:     "MERN BootCamp",
			Price:    299,
			Platform: "LearnCodeOnline",
			Password: "a23",
			// Tags:     []string{"google", "js"},
			Tags: nil,
		},
	}
	finalJson, err := json.Marshal(lcoCourses)
	fJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	fmt.Println(string(finalJson))
	fmt.Println(string(fJson))
}

func DecodeJSON() {
	jsonDataFromWeb := []byte(`
{
                "coursename": "ReactJS BootCamp",
                "price": 299,
                "platform": "LearnCodeOnline",
                "tags": [
                        "web-dev",
                        "js"
                ]
        }`)
	var lcoCourse course
	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("JSON is valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf( "%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON is not valid")
	}


}
