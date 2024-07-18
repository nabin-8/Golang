package main

import (
	"fmt"
	// "slices"
)

// slices -> dynamic array
// most used cobstruct in go
// + useful methods

func main() {
	// unitinilize slice is nil
	// var nums []int

	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))

	// var nums = make([]int, 0, 5)
	// capacity -> maximum number of elements can fit
	// fmt.Println(nums == nil)
	// nums := []int{}
	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums[0]=3
	// fmt.Println(nums)
	// fmt.Println(cap(nums))

	// var nums = make([]int, 0, 5)
	// nums = append(nums, 2)
	// var nums2 = make([]int, len(nums))
	// copy function
	// copy(nums2, nums)
	// fmt.Println(nums, nums2)

	// slice operator
	// var nums = []int{1, 2, 3, 4, 5}

	// fmt.Println(nums[0:2])
	// fmt.Println(nums[1:])

	// slice
	// var nums1 = []int{1, 2, 4}
	// var nums2 = []int{1, 2, 3}

	// fmt.Println(slices.Equal(nums1, nums2))

	var nums = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(nums)
}
