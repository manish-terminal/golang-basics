package main

import (
	"fmt"
	"os"
)

func main(){

	f,err:=os.Open("./example.txt")
	if(err!=nil){
		//log error
		panic(err)
	}

	fileInfo,err:=f.Stat()
		if(err!=nil){
		//log error
		panic(err)
	}
	fmt.Println("file name",fileInfo.Name())
	fmt.Println(fileInfo.ModTime())


	//read file
	f,err=os.Open("./example.txt")
	if err!=nil{
		panic(err)
	}
	buf:=make([]byte,13)
	d,err:=f.Read(buf)
	if err!=nil{
		panic(err)
	}

	defer f.Close()
	print("size: ",d," content: ",string(buf))

	data,err:=os.ReadFile("./example.txt")
	if err!=nil{
		panic(err)
	}
	println("content via readfile:",string(data))

	dir,err:=os.Open("../")
	if(err!=nil){
		panic(err)
	}
	defer dir.Close()

	inf,err:=dir.ReadDir(-1)
	for _,f:=range inf{
		fmt.Println(f,f.IsDir())
	}
	bytes:=[]byte("hello world")

	//create a file

	f1,err:=os.Create("./example2.txt")
	if(err!=nil){
		panic(err)
	}
	defer f1.Close()
	f1.WriteString("Hi Manish Gupta\n")
	f1.WriteString("Nice Language\n")
	f1.Write(bytes)
	os.Remove("./example2.txt")
}