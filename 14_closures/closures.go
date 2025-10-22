package main

import "fmt"

func main(){

increment:=counter()
fmt.Println(increment())
fmt.Println(increment())
fmt.Println(increment())
}

func counter()func()int{

	var count int=0

	return  func() int {
		count++
		return count
	}
}