package main

import "fmt"

func main()  {
	//iterating over data structures
	nums:=[]int{6,7,8}
	for i:=0;i<len(nums);i++{
		fmt.Println(nums[i])
	}
	sum:=0
	for index,num:=range nums{
		sum+=num
		fmt.Println(index,":",num)
	}
	fmt.Println(sum)
	m:=map[string]string{"fname":"john","lname":"doe"}

	for k:=range m{
		fmt.Println(k)//only key
	}
	for k,v:=range m{
		fmt.Println(k)//only key
		fmt.Println(v)//only value
	}
	for i,c:=range "manish"{
		fmt.Println(i,string(c))
	}

}