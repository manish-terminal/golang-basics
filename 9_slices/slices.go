package main

import (
	"fmt"
	"slices"
)

func main()  {
	//uninitialized slice is nil
	var nums[]int

	fmt.Println(nums==nil)
	fmt.Println(len(nums))

	var nums2=make([]int,0,7)
	fmt.Println(cap(nums2))//maximum number of elements can fit
	fmt.Println(nums2)
	nums2 = append(nums2, 1)
	nums2 = append(nums2, 1)
	nums2 = append(nums2, 1)
	nums2 = append(nums2, 1)
	nums2 = append(nums2, 1)
	nums2 = append(nums2, 2)
	nums2 = append(nums2, 3)
	nums2 = append(nums2, 4)
	nums2 = append(nums2, 4)
	nums2 = append(nums2, 4)
	nums2 = append(nums2, 4)
	nums2 = append(nums2, 4)
	nums2 = append(nums2, 4)
	nums2 = append(nums2, 4)
	//capacity is resized automatically by double
	fmt.Println(nums2)
	fmt.Println(cap(nums2))
	nums3:=[]int{}
	fmt.Println(nums3)
	
	//copy function
	var nums4=make([]int,len(nums2))
	copy(nums4,nums2)
	fmt.Println(nums4)

	//slice operator
	var nums5=[]int{1,2,3}
	fmt.Println(nums5[0:])//end is not included,koi ek gayab kr do toh shuru se ya end se tk..

	//slices package has some inbuilt functions
	fmt.Println(slices.Equal(nums2,nums2))

	var nums6=[][]int{{1,2},{3,4}}
	fmt.Println(nums6)



}