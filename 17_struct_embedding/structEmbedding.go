package main

import (
	"fmt"
	"time"
)


type customer struct{
	name string 
	phone string
}

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
	customer
}

func main() {
	newCustomer := customer{
		name: "hong",
		phone: "123813",
	}

	newOrder := order{
		id: "1",
		amount: 10,
		status: "recieved",
		customer: newCustomer,
		
	}
	fmt.Println(newOrder.customer)

}
