package main

// import "fmt"

// func add(a, b int) int {
// 	return a + b
// }

// func getLanguages() (string, string, bool) {
// 	return "Golang", "Javascript", true
// }

// func processIt(fn func(a int) int) {
// 	fn(1)
// }

func processIt() func(a int) int {
	return func(a int) int {
		return 4
	}
}

func main() {
	// result := add(5, 5)
	// fmt.Println(result)

	// lang1, lang2, _ := getLanguages()
	// fmt.Println(lang1, lang2)

	// fn := func(a int) int {
	// 	return 2
	// }
	// processIt(fn)

	fn := processIt()
	fn(6)
}
