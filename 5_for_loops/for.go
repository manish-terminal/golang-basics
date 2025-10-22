package main

import "fmt"

func main()  {
	//while go doesnt have
	i:=1
	for i<=5{//brackets not needed
		fmt.Println(i)
		i++
	}
	//infinite loop
	// for{
	// 	fmt.Println(1)
	// }

	for i:=0;i<=3;i++{
		if(i==2){
			continue
		}
		fmt.Println(i)
	}

	for i:=range 11{
		fmt.Println(i)
	}

}