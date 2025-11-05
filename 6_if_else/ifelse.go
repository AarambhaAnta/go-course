package main

import "fmt"

func main() {
	// age := 16

	// if age >= 18 {
	// 	fmt.Println("Person is an adult")
	// } else {
	// 	fmt.Println("Person is not an adult")
	// }

	// age := 18

	// if age >= 18 {
	// 	fmt.Println("Person is an adult")
	// }else if age >= 12 {
	// 	fmt.Println("Person is teenager")
	// }else {
	// 	fmt.Println("Person is a kid")
	// }

	// var role = "admin"
	// var hasPermissions = false

	// if role == "admin" && hasPermissions {
	// 	fmt.Println("Yes")
	// }

	// we can declare a variable inside `if construct`
	if age := 20; age >= 18 {
		fmt.Println("Person is an adult", age)
	}else if age >= 12 {
		fmt.Println("Person is teenager", age)
	}

	// go does not have ternary operator
}