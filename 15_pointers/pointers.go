package main

import "fmt"

func changeNum(num int){
	num=5
	fmt.Println("In changeNum",num)
}

func changeNumByAddress(num *int){
	*num=5
	fmt.Println("In changeNum",num)//dono me same address aa rha
	fmt.Println("In changeNum",*num)
}


func main(){
	num:=1
	changeNum(num)
	fmt.Println(num)//number change nhi hua sirf copy pass hota hai copy by value not copy by reference
	
	//agr hamko actual change krna hai toh we use concept of pointers reference memory location
	
	fmt.Println("Memory address",&num)
	changeNumByAddress(&num)
		fmt.Println(num)


}