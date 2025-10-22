package main

import "fmt"

func main(){
	age:=17
	if(age>=30 && age>0){
		fmt.Println("Eligible for voting")
	} else if(age>=18){
		fmt.Println("yaada")
	}else{
		fmt.Println("shava shava")
	}

	if mage:=15;mage>18{
		fmt.Println("adult",mage)
	}else{
		fmt.Println("not an adult",mage)
	}
//we can declare variable inside if
// &&-and
// ||-or
//go doesnt have ternary operator

}