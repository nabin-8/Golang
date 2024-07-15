package main

import "fmt"

func main() {
	// age := 16

	// if age >= 18 {
	// 	fmt.Println("Person is an adult")
	// } else {
	// 	fmt.Println("Person is not an adult")
	// }

	// if age >= 18 {
	// 	fmt.Println("Person is not an adult")

	// } else if age >= 12 {
	// 	fmt.Println("Person is Teneger")
	// } else {
	// 	fmt.Println("Person is kid")
	// }

	// var role = "admin"
	// var hasPremission = true

	// if role == "admin" && hasPremission {
	// 	fmt.Println("yes")
	// }

	// we can declare a variable inside if
	if age := 15; age >= 18 {
		fmt.Println("Person is an adult", age)
	} else if age >= 12 {
		fmt.Println("Persion is teenager", age)
	}

	// go does not have ternary, you will have to use normal if-else
}
