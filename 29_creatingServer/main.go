package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// We'll be adding code in the next steps
	fmt.Println("Creating a basic web server in Go")
	// PerformGetRequest("https://jsonplaceholder.typicode.com/comments?postId=1")
	// PerformPostJSONRequest()
	PerformPostFormRequest()
}

func PerformGetRequest(url string) {

	response, error := http.Get(url)
	if error != nil {
		fmt.Println("Error occurred while making GET request:", error)
		return
	}
	defer response.Body.Close()
	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Println("Byte count is:", byteCount)
	fmt.Println("Response string is:", responseString.String())
	body, _ := io.ReadAll(response.Body)
	fmt.Println("Response body is:", string(body))
	fmt.Println("Response status is:", response.Status)

}

func PerformPostJSONRequest() {
	url := "https://jsonplaceholder.typicode.com/posts"
	requestBody := strings.NewReader(`{
	"name":"Manish Gupta",
	"age":20,
	"platform":"LearnCodeOnline"
	}`)
	response, err := http.Post(url, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Print(string(content))
}

func PerformPostFormRequest() {
	myUrl := "https://jsonplaceholder.typicode.com/posts"
	data := url.Values{}
	data.Add("firstname", "Manish")
	data.Add("lastname", "Gupta")
	response,err:=http.PostForm(myUrl,data)
	if err!=nil{
		panic(err)
	}
	content,_:=io.ReadAll(response.Body)
	fmt.Println(string(content))
}
