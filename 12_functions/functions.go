package main

import "fmt"

func main()  {
	c:=add(3,4)
	lang1,lang2,check:=getLanguages()
	fmt.Println(c)
	fmt.Println(lang1,lang2,check)
	fn:=func (a int)int  {
		return 2*a
	}
	processIt(fn)
}

func add(a ,b int )int{
	return a+b
}

func getLanguages()(string,string,bool){
	return "golang","javascript",false
}
//to supress compiler use underscore.

func processIt(fn func(a int)int){
	fn(1)
}