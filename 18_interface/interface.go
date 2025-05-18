package main

import "fmt"


type paymenter interface{
	pay(amount float32) //method acting as a contract 
	//here any gateway needs to have a pay method

	//any struct having pay method would work
}

type payment struct {
	gateway paymenter
}

func (p payment) makepayment(amount float32) {
	p.gateway.pay(amount)
}

type razorpay struct {
}

func (r razorpay) pay(amount float32) {
	//logic to make payment
	fmt.Println("making payment using razorpay", amount)
}

type stripe struct {
}

func (s stripe) pay(amount float32) {
	fmt.Println("making payment using stripe", amount)
}

type paypal struct{}
func (p paypal) pay(amount float32){
	fmt.Println("making payment using paypal", amount)
}

type fakePayment struct{}
func (f fakePayment) pay(amount float32){
	fmt.Println("making payment using fake gateway", amount)
}

func main() {
	// stripeGateway := stripe{}
	// razorpayGw := razorpay{}
	fakeGw := fakePayment{}
	newPayment := payment{
		gateway: fakeGw,
	}
	newPayment.makepayment(100)
}
