package main

import "fmt"

func main(){
	fmt.Println(1,2,3,4,5,"hello")//any number of parameters
	res:=sum(3,4,5,6)
	fmt.Println(res)
}

func sum(nums ...int)int{
	total:=0
	for _,num:= range nums{
		total+=num
	}
	return  total

}