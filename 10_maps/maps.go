package main

import (
	"fmt"
	"maps"
)

// maps->hash,object,dict
func main() {
	m := make(map[string]string)
	//setting an element
	m["name"] = "golang"
	m["area"] = "backend"

	//get element
	fmt.Println(m["name"])
	fmt.Println(m["area"])

	//say if key doesnt exist
	fmt.Println(m["phone"]) //empty value say just falsy values
	m2 := make(map[string]int)
	m2["age"] = 18
	m2["price"] = 180
	fmt.Println(m2["phone"]) //0
	fmt.Println(len(m2))
	delete(m2, "age")
	fmt.Println(len(m2))

	//without make function
	m3 := map[string]int{"price": 180, "phone": 3}
	fmt.Println(m3)

	v, ok := m3["price"]
	fmt.Println(v)
	if ok {
		fmt.Println("all okay")
	} else {
		fmt.Println("not okay")
	}

	m4 := map[string]int{"price": 180, "phone": 3}
	fmt.Println(maps.Equal(m3, m4))
}
