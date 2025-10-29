package main

import (
	"fmt"
	"time"
)

// orders
type order struct {
	id string
	amount float32
	status string 
	createdAt time.Time   // nanosecond precision
}

func newOrder (id string, amount float32, status string) *order {
	myOrder := order{
		id: id,
		amount: amount,
		status: status,
	}
	return &myOrder
}

// (o order) means function only for order structs

func (o *order) changeStatus(msg string) {     // pointers just to change the value. otherwise, go creates a copy of the msg
	o.status = msg                             // to make changes, pointers are necessary
}   // no dereferencing needed here, go does the work here..

func (o *order) updateAmount(amount float32) {
	o.amount = amount
}

func (o order) getAmount() float32 {   // no struct values are being altered, so no pointers
	return o.amount
}



func main() {
	

	order1 := order{
		id: "1",
		amount: 50.92,
		status: "pending",
	}

	order2 := order{
		id: "2",
		amount: 200,
		status: "received",
		createdAt: time.Now(),
	}

	order1.createdAt = time.Now()
	fmt.Println(order1)

	fmt.Println(order1.status)
	order2.changeStatus("done done done")    // values altered
	order2.updateAmount(10)
	fmt.Println(order2)

	x := order2.getAmount()
	fmt.Println("amount", x)

}
