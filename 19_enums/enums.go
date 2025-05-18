package main

import "fmt"

//enumerated types
type OrderStatus string

const (
	Recieved OrderStatus = "recieved"
	Confirmed = "confirmed"
	Prepared = "prepared"
	Delivered = "delivered"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("changing status to", status)
}

func main() {
	changeOrderStatus(Confirmed)

}
