package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	defer wg.Done()
	go Greeter("Hello")
	Greeter("World")
	getStatusCode("https://pkg.go.dev/context")
	time.Sleep(1 * time.Second)
	websiteList := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.linkedin.com",
		"https://www.linkedin.com",
		"https://www.linkedin.com",
		"https://www.linkedin.com",
		"https://www.linkedin.com",
		"https://www.linkedin.com",
		"https://www.linkedin.com",
		"https://www.linkedin.com",
		"https://www.linkedin.com",
		"https://www.linkedin.com",
		"https://www.linkedin.com",
		"https://www.github.com",
	}
	for _, website := range websiteList {
		go getStatusCode(website)
		wg.Add(1)
	}
	time.Sleep(2 * time.Second) //for waiting all go routines to complete
	

}

func Greeter(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}

func getStatusCode(endpoint string) int {
	resp, err := http.Get(endpoint)
	if err != nil {
		// fmt.Println("Status Code:", resp.StatusCode)
		fmt.Println("Error:", err)
		return 0
	}
	defer resp.Body.Close()
	fmt.Println("Status Code:", resp.StatusCode)
	return resp.StatusCode
}
