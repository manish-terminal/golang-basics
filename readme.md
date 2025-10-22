5 reasons to choose Golang
1. Build Time
2. Fast startup-Code compiles to machine code.
3. Performance and Efficiency-Memory usage is less.
4. Concurrency Model-Multithreaded 
5. Static typing-type should be defined. Errors could be caught during compile time.

Commands
go build ___
go run ___

In go we usually use slices instead of arrays
nums:=[]int{}
var nums2=make([]int,0,7)
0 number of 0s
capacity is 7
nums2 = append(nums2, 1)

//map
	m2:=make(map[string]int)
    //without make function
	m3:=map[string]int{"price":180,"phone":3}

//learn about returning function or accepting function


//receiver type function
func (o *order) changeStatusofOrder(status string){
	o.status=status
}
func (o *order) getAmount()float32{
	return o.amount
}