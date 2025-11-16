package main

import "fmt"

// func printSlice[T comparable](items []T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

func printSlice[T int, V string](items []T, name V) {
	for _, item := range items {
		fmt.Println(item, name)
	}
}

// func printSlice[T any](items []T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// func printSlice[T interface{}](items []T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// func stringSlice(items []string) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// // LIFO
// type stack[T any] struct {
// 	elements []T
// }

func main() {

	// myStack := stack[string]{
	// 	elements: []string{"golang","typescript"},
	// }

	// fmt.Println(myStack)

	nums := []int{1, 2, 3}
	// printSlice(nums)

	names := []string{"golang", "typescript"}
	// stringSlice(names)
	// printSlice(names)

	// vals := []bool{true, false, true, false, false}
	// printSlice(vals)

	printSlice(nums, names[0])
}
