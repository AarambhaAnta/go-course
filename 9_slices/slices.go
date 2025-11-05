package main

import (
	"fmt"
)

// slice -> dynamic arrays
// most used construct in go
// + useful methods
func main() {
	// uninitialized slice is nil
	// var nums []int
	
	// fmt.Println("nums[]: ", nums)
	// fmt.Println("Is nums[] nil", nums == nil)
	// fmt.Println("nums[] length: ", len(nums))

	// nums := []int{}
	// var nums = make([]int, 2, 5)

	// fmt.Println("nums[]: ", nums)
	// fmt.Println("Is nums[] nil: ", nums == nil)
	// fmt.Println("nums[] capacity: ", cap(nums))

	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 3, 4)

	// fmt.Println("nums[]: ", nums)
	// fmt.Println("nums[] capacity: ", cap(nums))
	// fmt.Println("nums[] length: ", len(nums))

	// copy function

	// nums := []int{}
	// nums = append(nums, 2)
	// var arr = make([]int, len(nums))

	// fmt.Println("nums[]: ", nums)
	// fmt.Println("arr[]: ", arr)

	// copy(arr, nums)

	// fmt.Println("nums[]: ", nums)
	// fmt.Println("arr[]: ", arr)

	// slice operator

	// var nums = []int{1, 2, 3}

	// fmt.Println("num[0],...,nums[1]: ", nums[0:3])
	// fmt.Println("num[0],...,nums[1]: ", nums[:3])
	// fmt.Println("num[0],...,nums[1]: ", nums[0:])

	// slices
	// var nums1 = []int{1, 2}
	// var nums2 = []int{1, 2}

	// fmt.Println("Is nums1[] == nums2[]: ", slices.Equal(nums1, nums2))

	// 2d slices
	var nums = [][]int{{1, 2, 3}, {4, 5}}

	fmt.Println("nums[]: ", nums)
}