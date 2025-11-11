package main

import "fmt"

// // order struct
// type order struct {
// 	id        string
// 	amount    float32
// 	status    string
// 	createdAt time.Time // nanosecond precision
// }

// // hack for constructor
// func newOrder(id string, amount float32, status string) *order {
// 	// inital setup goes here...
// 	myOrder := order{
// 		id:     id,
// 		amount: amount,
// 		status: status,
// 	}
// 	return &myOrder
// }

// // receiver type
// func (o *order) changeStatus(status string) {
// 	o.status = status
// }
// func (o order) getAmount() float32 {
// 	return o.amount
// }

func main() {
	// myOrder := order{
	// 	id:     "1",
	// 	amount: 50.99,
	// 	status: "received",
	// }
	// myOrder.createdAt = time.Now()

	// fmt.Println("Order struct: ", myOrder)
	// fmt.Println(myOrder.status)

	// myOrder2 := order{
	// 	id: "2",
	// 	// amount: 100.00,
	// 	status:    "delivered",
	// 	createdAt: time.Now(),
	// }

	// myOrder2.changeStatus("paid")
	// fmt.Println("myOrder2 ammount: ", myOrder2.getAmount())
	// fmt.Println("Order struct: ", myOrder2)

	// myOrder := newOrder("1", 30.50, "received")
	// fmt.Println(myOrder)

	language := struct {
		name   string
		isGood bool
	} {"golang", true}

	fmt.Println(language)
}
