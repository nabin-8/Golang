package main

import "fmt"

func main() {
	// for -> only construct in go for looping

	// while loop
	// i := 1
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

	// infinite loop
	// for {
	// 	println("1")
	// }

	// classic for loop
	// for i := 0; i <= 3; i++ {
	// 	// break
	// 	if i == 2 {
	// 		fmt.Println("if exexutes continue")
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// 1.22 range
	for i := range 10 {
		fmt.Println(i)
	}
}
