package main

import "fmt"

// func cnangeNum(num int)  {
// 	num = 5
// 	fmt.Println("In changeNum ", num)
// }

// by reference
func changeNum(num *int)  {
	*num = 5
	fmt.Println("In ChangeNum ", *num)
}
func main() {
	num := 1
	// cnangeNum(num)
	// fmt.Println("After changeNum in main ", num)

	fmt.Println("Memory address: ", &num)
	changeNum(&num)
	fmt.Println("After changeNum in main ", num)
}