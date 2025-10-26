package main

import (
	"fmt"
	"net/url"
)

const myURL string = "https://jsonplaceholder.typicode.com/comments?postId=1"

func main()  {
	fmt.Println("Handle urls in go lang")
	fmt.Println(myURL)

	result, err := url.Parse(myURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}
	fmt.Println("Result:", result)
	fmt.Println("Scheme:", result.Scheme)
	fmt.Println("Host:", result.Host)
	fmt.Println("Path:", result.Path)
	fmt.Println("Query:", result.RawQuery)
	fmt.Println("Port:", result.Port())
	queryParams := result.Query()
	fmt.Printf("Type of queryParams is: %T\n", queryParams)
	fmt.Println("Query Params:", queryParams["postId"])
	keys := make([]string, 0, len(queryParams))
	for k := range queryParams {
		keys = append(keys, k)
	}
	fmt.Println("Keys:", keys)
	for _,value:=range queryParams{
		fmt.Println("Value is:",value)
	}//order isnt guaranteed


		parts := &url.URL{
		Scheme:   "https",
		Host:     "jsonplaceholder.typicode.com",
		Path:    "/comments",
		RawQuery: "postId=1",
		}
	fmt.Println("URL is:", parts.String())
	}