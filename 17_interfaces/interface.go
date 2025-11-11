package main

import "fmt"

type paymenter interface {
	pay(amount float32)
}

type payment struct {
	gateway paymenter
}

// open-close principle
func (p payment) makePayment(amount float32) {
	// razorpayPaymentGw := razorpay{}
	// razorpayPaymentGw.pay(amount)
	// stripPaymentGw := strip{}
	p.gateway.pay(amount)
}

type razorpay struct {
}

func (r razorpay) pay(amount float32) {
	// logic to make payment
	fmt.Println("Making payment using razorpay: ₹", amount)
}

// type strip struct {
// }

// func (s strip) pay(amount float32) {
// 	fmt.Println("Making payment using strip: ₹", amount)
// }

type fakepay struct {
}

func (f fakepay) pay(amount float32) {
	fmt.Println("Making payment using fake gateway for testing purpose.")
}

type paypal struct {
}

func (p paypal) pay(amount float32) {
	fmt.Println("Making payment using paypal: ₹", amount)
}

func main() {
	// stripPaymentGw := strip{}
	// razorpayPaymentGw := razorpay{}
	// fakepayPaymentGw := fakepay{}
	paypalPaymentGw := paypal{}
	newPayment := payment{
		gateway: paypalPaymentGw,
	}
	newPayment.makePayment(100)
}
