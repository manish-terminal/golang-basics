package main

import (
	"fmt"
	"time"
)

func main() {
	// simple switch

	i := 5
	switch i {
	case 1:
		fmt.Println("One")
		break
	case 2:
		fmt.Println("Two")
		break
	case 3:
		fmt.Println("Three")
		break
	case 4:
		fmt.Println("Four")
		break
	case 5:
		fmt.Println("Five")
		break
	default:
		fmt.Println("ANY")
		//break is optional in go
	}

	//multiple condition switch

	switch time.Now().Weekday() {
	case time.Saturday , time.Sunday:fmt.Println("Weekend")
	default:fmt.Println("Weekday")
	}

	//type switch
	whoAmI:=func (i interface{})  {
		switch i.(type){
		case int:fmt.Println("integer")
		case string:fmt.Println("string")
		case bool:fmt.Println("boolean")
		default:
			fmt.Println("other")
		}
	}
	whoAmI(5)
	whoAmI("5")
	whoAmI(false)
	whoAmI(14.3)

}
