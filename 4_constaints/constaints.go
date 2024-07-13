package main

import "fmt"

// const age=30
// name:="Golang is not allowed outside function/ global variable"
func main() {
	// const name string = "Golang"

	// fmt.Println(name)

	const (
		port = 555
		host = "localhost"
	)

	fmt.Println(port, host)

}
