![golang](https://go.dev/images/go-logo-white.svg)

# Golang

- Go is a statically typed, compiled high-level programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson. It is syntactically similar to C, but also has memory safety, garbage collection, structural typing, and CSP-style concurrency.

|Days|Link|Days|Link|Days|Link|Days|Link|
|---|---|---|---|---|---|---|---|
|1|[Golang Features](#golang-features)|2|[Reasons to choose golang](#5-reasons-to-choose-golang)|3|[Installation](#go-installation-and-hello-world)|4|[Simple Variables](#simple-valuesdatatypes)|
|5|[Constants](#constants-in-go)|

##

#Day1

### Golang Features

- Multi-Threading
- Mutexes
- Buffers
- Networks
- Standard Applications

##### Using Golang yo can make

- Powerful Backend
- Cloud Applications

##### Golang Applications

- Docker Build in Go
- Kubernetes

##

#Day2

### 5 Reasons to choose Golang

1. Build time
1. Fast starup
1. Performance and Efficiency
1. Concurrency Model
1. Static typing and compilation

##

#Day3

### Go Installation and Hello World

- [Download Go](https://go.dev/dl/)
- Current Version of Go is `go1.22.5`
- check version of Go `go version`
- If successfully installed go in your machine it will return
- go version go1.22.5 windows/amd64

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}
```

- To run go: `go build filename`
- It will build .exe `executable file` then you can run it.
- There is another way to direct run it.
- `go run filename`

#### Code explanation

- **package main:** Every go program have main package
- **func:** keyword is used to make function in go.
- **main():** main function of go program/file
- **fmt:** is Standard library package used to formatting it has one method called `Println`
- **Println:** id used to print in console.


##

#Day4
### Simple Values/Datatypes
- int
- string
- bool
- float

```go
package main

import (
	"fmt"
)
func main() {
	// Simple values
	// int
	fmt.Println(1 + 1)
	// string
	fmt.Println("hello Golang")
	// bool
	fmt.Println(false)
	fmt.Println(true)
	// float
	fmt.Println(10.5)
	fmt.Println(7.0 / 3.0)
}

```
### Variables
#### Way to Declare Variable in Go
1. var name string ="Golang"
1. var name="Golang"
1. name := "Golang"

```go
package main

import "fmt"

func main() {
	// var name string = "Golang"

	// automatically type infer
	// var name = "Golang"
	// var isAdult=true

	// var age int = 22
	// there is different int16, 32, 64
	// if you use int only golang automatically manage according to its assign value internally

	// short hand syntax
	// name := "Golang"
	// in golang := is used when we declare value and same time assigned into it

	// then why we use upper two ways to declare variables
	// var name string
	// if we have declare variable but not assigned value
	// name = "Golang"

	// var price float32 = 50.5
	// var price =50.5
	price := 50.5

	fmt.Println(price)
}
```

##

#Day5
### Constants in Go
- Constant are the fixed value which does not change over time like other variables created using `var` keyword
- Constant are declared using const keyword in Go.
- int

```go
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
```

##

#Day5
### Looping in Golang
- In Golang there is only way to looping is `for` loop.
- we can achieve `infinite` loop and there is one new feature comes in golang 1.22 which is `range` given example below.

```go
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
```

##

#Day6
### Conditional statement if-else
- go does not have `ternary` in version 1.22
```go

	age := 16

	if age >= 18 {
		fmt.Println("Person is an adult")
	} else {
		fmt.Println("Person is not an adult")
	}

	if age >= 18 {
		fmt.Println("Person is not an adult")

	} else if age >= 12 {
		fmt.Println("Person is Teneger")
	} else {
		fmt.Println("Person is kid")
	}

	// we can declare a variable inside if
	if age := 15; age >= 18 {
		fmt.Println("Person is an adult", age)
	} else if age >= 12 {
		fmt.Println("Persion is teenager", age)
	}

	// go does not have ternary, you will have to use normal if-else
```


##

#Day7
### Switch statement in go
- switch is used to check multiple conditions
```go
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
```
##

#Day7
### Arrays in go
- Arrays are numbered sequence of specific length
- when do we use array in golang
	- fixed size, that is predictable
	- Memory optimization
	- Contant time access
```go
var nums [4]int

	// array length
	nums[0] = 1
	fmt.Println(len(nums))
	
	// to declare in single line
	nums := [3]int{1, 2, 3}
	fmt.Println(nums)

	// 2d array
	nums := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(nums)
```