package main

import "fmt"

// by value
// func changeNum(num int) {
// 	num = 5
// 	fmt.Println("In changeNum", num)
// }

// by reference
// func changeNum(num *int) {
// 	*num = 5
// 	fmt.Println("In changeNum", *num)
// }

func changeSlice(nums *[]int) {
	*nums = append(*nums, 4)
}

func main() {
	// num := 1
	// changeNum(&num)
	// fmt.Println("Memory address", &num)

	// fmt.Println("After changeNum in main", num)

	var nums = []int{1, 2, 3}
	changeSlice(&nums)
	fmt.Println("After changing", nums)
}
