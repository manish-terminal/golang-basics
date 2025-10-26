package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://jsonplaceholder.typicode.com/posts"

func main() {
	fmt.Println("Web request to", url)
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	// fmt.Println("Response is of type : %T\n",response)
	fmt.Println(response.Body)
	defer response.Body.Close() //close the connection

	databytes,err:=io.ReadAll(response.Body)
	if err!=nil{
		panic(err)
	}
	content:=string(databytes)
	fmt.Println(content)

}
