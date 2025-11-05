package main

import "fmt"

func sum(nums ... int) int {
	total := 0
	for _, num := range nums {
		total = total + num
	}

	return total
}

func anytype(vals ... interface{}) {
	for _, val := range vals {
		fmt.Println(val)
	}
}

func main() {
	// variadic function
	fmt.Println(1, 2, 3, 4, "hello")

	fmt.Println("sum: ", sum(1,2,3,4,5))

	nums := []int{1,2,3,4,5}
	fmt.Println("sum: ", sum(nums...))
	anytype(3, 3.5, "hello", false)
}