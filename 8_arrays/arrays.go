package main

import "fmt"

// numbered sequence of specific length
func main() {
	// integer array
	// default value: 0
	var nums [4]int

	nums[0] = 1

	fmt.Println("nums[0]: ", nums[0])
	fmt.Println("nums[]: ", nums)
	// array length
	fmt.Println("nums length:", len(nums))

	// boolean array
	// default value: false
	var vals [4]bool

	vals[2] = true

	fmt.Println("vals[]: ", vals)

	// string array
	// default value: ""
	var names [3]string

	names[0] = "golang"
	fmt.Println("names[]: ", names)

	// single line declaration
	arr := [3]int {1, 2, 3}

	fmt.Println("arr[]: ", arr)

	// 2d arrays
	mat := [2][2]int {{1,2},{3,4}}

	fmt.Println("mat[][]: ", mat)
}