package main

import (
	"fmt"
	// "time"
)

func main() {
	// 	// simple switch
	// 	i := 5

	// 	switch i {
	// 	case 1:
	// 		fmt.Println("one")
	// 	case 2:
	// 		fmt.Println("two")
	// 	case 3:
	// 		fmt.Println("three")
	// 	default:
	// 		fmt.Println("other")
	// 	}

	// In go we can use switch without break like other languages

	// multiple condition switch

	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("it's weekend")

	// default:
	// 	fmt.Println("its workday")
	// }

	// type switch
	whoAMI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("its an integer")
		case string:
			fmt.Println("its an string")
		case bool:
			fmt.Println("its an bool")
		default:
			fmt.Println("other", t)

		}
	}

	whoAMI(11)
}
