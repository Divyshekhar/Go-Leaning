package main

import (
	"fmt"
	"time"
)

// ecommerce struct
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time //nanosecond precision
}

// constructer
func newOrder(id string, amount float32, status string) *order {
	new := order{
		id:     id,
		amount: amount,
		status: status,
	}
	return &new
}

// reciever type
func (o *order) changeStatus(status string) {
	o.status = status
}

//when getting the value no need to pass by reference

func (o order) getAmount() float32 {
	return o.amount
}

// if you dont set any of the field then it takes the zero value of the type
func main() { 
	//inline struct
	language := struct {
		name   string
		isGood bool
	}{"golang", true}

	fmt.Println(language)

	orderInstance := order{
		id:     "1",
		amount: 500.0,
		status: "recieved",
	}
	orderInstance.changeStatus("confirmed lmao")

	orderInstance.createdAt = time.Now()

	fmt.Println("amount: ", orderInstance.getAmount())

	fmt.Println("order", orderInstance)

	myOrder := order{
		id:        "2",
		amount:    20.4,
		status:    "delivered",
		createdAt: time.Now(),
	}
	fmt.Println(myOrder)

}
