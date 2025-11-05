package main

import "fmt"

const age = 30
var name string = "golang"

func main()  {
	// const name = "golang"
	// const age = 30

	fmt.Println(age)
	fmt.Println(name)

	// constants grouping
	const(
		port = 5000
		host = "localhost"
	)

	fmt.Println(port, host)
}