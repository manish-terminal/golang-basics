package main

import "fmt"

func main()  {
	var nums[4]int//all values are initialized with 0
	//zeroed values by default
	//int->0,string->"",bool->false,float->0.0
	fmt.Println(nums)
	fmt.Println(len(nums))//array length
	var tf[4]bool
	var name[4]string

	nums[0]=3
	fmt.Println(nums)
	tf[2]=true
	fmt.Println(tf)
	fmt.Println(name)

	var nums2=[4]int{1,2,3,4}
	fmt.Println(nums2)
	//to declare and initialize in single line
	nums3:=[5]int{1,2,3,4,5}
	fmt.Println(nums3)

	nums4:=[2][4]int{{1,2,3,4},{1,3,5,6}}
	fmt.Println(nums4)
	//use when u know size-memory optimization,constant time access
	


}