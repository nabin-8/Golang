package main

import "fmt"

func main() {
	// numbered sequence of specific length
	// zero values
	// int->0, string->"", bool->false
	// var nums [4]int

	// array length
	// fmt.Println(len(nums))
	// nums[0] = 1
	// fmt.Println(nums)

	// var vals [4]bool
	// vals[2] = true
	// fmt.Println(vals)

	// var name [3]string
	// name[0] = "true"
	// fmt.Println(name)

	// to declare in single line
	// nums := [3]int{1, 2, 3}
	// fmt.Println(nums)

	// 2d array
	nums := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(nums)

	// fixed size, that is predictable
	// Memory optimization
	// Contant time access
	// slices
}
