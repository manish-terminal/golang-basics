package main

import "fmt"
var age int=24

func main()  {
	const name="Manish"//cant be changed. these can be declared outside. you need to assign while declaring
	age=17
	//constant grouping
	const(
		port=5000
		host="local host"
	)


	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(port)
	fmt.Println(host)
}
