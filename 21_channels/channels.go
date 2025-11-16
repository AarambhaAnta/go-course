package main

import (
	"fmt"
	"time"
)

// func processNum(numChan chan int) {
// 	for num := range numChan {
// 		fmt.Println("Processing number... ", num)
// 		time.Sleep(time.Second)
// 	}
// }

// func sum(result chan int, num1 int, num2 int) {
// 	numResult := num1 + num2
// 	result <- numResult
// }

// // goroutine synchronizer
// func task(done chan bool) {
// 	defer func() { done <- true }()
// 	fmt.Println("processing...")
// }

// func emailSender(emailChan chan string, done chan bool) {
// 	defer func() { done <- true }()
// 	for email := range emailChan {
// 		fmt.Println("sending email to: ", email)
// 		time.Sleep(time.Second)
// 	}
// }

func emailSender(emailChan <-chan string, done chan<- bool) {
	defer func() { done <- true }()
	for email := range emailChan {
		fmt.Println("sending email to: ", email)
		time.Sleep(time.Second)
	}
}

func main() {
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "pong"
	}()

	for range 2 {
		select {
		case chan1Val := <-chan1:
			fmt.Println("received data from chan1: ", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("received data from chan2: ", chan2Val)
		}
	}

	// emailChan := make(chan string, 100)
	// done := make(chan bool)

	// go emailSender(emailChan, done)

	// for i := range 10 {
	// 	emailChan <- fmt.Sprintf("%d@gmail.com", i)
	// }

	// fmt.Println("done sending...")

	// close(emailChan)

	// // emailChan <- "1@example.com"
	// // emailChan <- "2@example.com"

	// // fmt.Println(<-emailChan)
	// // fmt.Println(<-emailChan)

	// <-done

	// done := make(chan bool)

	// go task(done)

	// <-done // blocks the program

	// result := make(chan int)

	// go sum(result, 4, 5)

	// res := <-result

	// fmt.Println(res)

	// numChan := make(chan int)

	// go processNum(numChan)

	// for {
	// 	numChan <- rand.IntN(100)
	// }

	// messageChan := make(chan string)

	// // sending data into channel
	// messageChan <- "ping"

	// // !channels are blocking -> arises deadlock
	// // receiving data out of channel
	// msg := <-messageChan

	// fmt.Println(msg)
}
