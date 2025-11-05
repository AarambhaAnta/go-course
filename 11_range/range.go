package main

import "fmt"

// iterating over data-structures
func main() {
	// nums := []int {6, 7, 8}

	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	// using for with range
	// sum := 0
	// for i, num := range nums {
	// 	fmt.Println(i," : ", num)
	// 	sum = sum + num
	// }

	// fmt.Println("sum: ", sum)

	//  range with maps

	// m := map[string]string {"fname": "john", "lname": "doe"}

	// for key, val := range m {
	// 	fmt.Println(key, " : ", val)
	// }

	// for key := range m {
	// 	fmt.Println(key)
	// }

	// range with strings

	// unicode code point rune
	// starting byte of rune
	for i, c := range "golang" {
		fmt.Println(i, " : ", c)
	}
}