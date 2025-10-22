package main

import (
	"fmt"
	"time"
)

// order
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time //nanosecond precision
}

func main() {
	order1 := order{id: "1", amount: 101, status: "confirmed"}
	order2 := order{id: "2", amount: 100, status: "delivered"}
	fmt.Println(order1)
	age:=24
	fmt.Printf("age is %d years\n",age)
	var age2 int=30
	fmt.Printf("age2 is %d years\n",age2)
	order1.createdAt=time.Now()
	order2.createdAt=time.Now()
	fmt.Println(order1.status)
	changeStatus("cancelled",&order2)
	fmt.Println(order2)
	order1.changeStatusofOrder("hola")
	fmt.Println(order1.getAmount())
	my:=newOrder("a",50,"pending")
	fmt.Println(my.amount)

}
func changeStatus(status string,order *order){
	(*order).status=status
}

//receiver type function
func (o *order) changeStatusofOrder(status string){
	o.status=status
}
func (o *order) getAmount()float32{
	return o.amount
}

func newOrder(id string,amount float32,status string) *order{
	order3 := order{id: id, amount: amount, status: status}
	return &order3

}