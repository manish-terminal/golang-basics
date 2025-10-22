package main

import (
	"fmt"
	"time"
)

func main() {
	// messageChan:=make(chan string)
	// messageChan <-"ping"// blocking

	// msg:=<-messageChan
	// fmt.Println(msg)

	numChan := make(chan int)
	go processNum(numChan)
	numChan <- 5

	time.Sleep(time.Second*2)
}
func processNum(numChan chan int) {
	fmt.Println("number", <-numChan)
}
