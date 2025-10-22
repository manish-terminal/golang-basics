package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Println("doing task", id)
}

func main() {
	fmt.Println(time.Now())
	for i := 0; i <= 10; i++ {
		// go task(i)
		
		go func(i int){
			fmt.Println(i)
		}(i)
	}
	fmt.Println(time.Now())
//main function ko rokne k liye
	time.Sleep(time.Second*2)
}
