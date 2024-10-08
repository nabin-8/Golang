![golang](https://go.dev/images/go-logo-white.svg)

# Golang

- Go is a statically typed, compiled high-level programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson. It is syntactically similar to C, but also has memory safety, garbage collection, structural typing, and CSP-style concurrency.

|Days|Link|Days|Link|Days|Link|Days|Link|
|---|---|---|---|---|---|---|---|
|1|[Golang Features](#golang-features)|2|[Reasons to choose golang](#5-reasons-to-choose-golang)|3|[Installation](#go-installation-and-hello-world)|4|[Simple Variables](#simple-valuesdatatypes)|
|5|[Constants](#constants-in-go)|6|[Loop](#looping-in-golang)|7|[Condition](#conditional-statement-if-else)|8|[switch](#switch-statement-in-go)|
|9|[Array](#arrays-in-go)|10|[slices](#slices)|11|[maps-in-golang](#maps-in-golang)|12|[range](#range-in-golang)|
|13|[Functions](#functions)|14|[variadic variadic-functionsfunctions](#variadic-functions)|

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

#Day6
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

#Day7
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

#Day8
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

#Day9
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

##

#Day10
### Slices
- slices -> dynamic array
- most used cobstruct in go
- +useful methods
```go
	// unitinilize slice is nil
	var nums []int

	var nums = make([]int, 0, 5)
	// capacity -> maximum number of elements can fit

	// slice operator
	var nums = []int{1, 2, 3, 4, 5}
	fmt.Println(nums[0:2])
	fmt.Println(nums[1:])

	fmt.Println(cap(nums))
	var nums = make([]int, 0, 5)
	nums = append(nums, 2)
	var nums2 = make([]int, len(nums))
	// copy function
	copy(nums2, nums)
	fmt.Println(nums, nums2)

	var nums = [][]int{{1, 2, 3}, {4, 5, 6}}
```

##

#Day11

### Maps in Golang
- maps are like: maps -> hash tables, object, dict
- way to declare maps in golang
```go
m := make(map[string]string)
// another way
m1 := map[string]int{"price": 40, "age": 20, "phones": 3}
// (map[key]value)
```
- way to add and get data and check
```go
	m := make(map[string]int)
	// add data
	m["age"] = 30
	m["price"] = 50
	// print,get data
	fmt.Println(m["age"])

	// check data
	k, ok := m["age"]
	fmt.Println(k)

	if ok {
		fmt.Println("all ok")
	} else {
		fmt.Println("not ok")
	}
```
- way to delete data and map
```go
	// to delete
	delete(m, "price")
	// to clear map
	clear(m)
	fmt.Println(len(m))
```

##

#Day12

### Range in Golang
- iterating over data structure
```go
	m := map[string]string{"fname": "john", "lname": "doe"}
		for k, v := range m {
			fmt.Println(k, v)
		}

	// i starting byte of rune	// i starting byte of rune
	// c is unicode
	for i, c := range "Golang" {
		fmt.Println(i, string(c))
	}
	// c is unicode
	for i, c := range "Golang" {
		fmt.Println(i, string(c))
	}
```
##

#Day13

### Functions
- In Go, functions are first-class values. This means that you can do with functions the same things you can do with all other values - assign functions to variables, pass them as arguments to other functions or even return functions from other functions.

```go
func add(a, b int) int {
	return a + b
}

func main() {
	result := add(5, 5)
	fmt.Println(result)
}
```
- return multiple values
```go
func getLanguages() (string, string, bool) {
	return "Golang", "Javascript", true
}
func main(){
	lang1, lang2, _ := getLanguages()
	fmt.Println(lang1, lang2)
}
```
- pass function as an argument
```go
func processIt(fn func(a int) int) {
	fn(1)
}

func main(){
	fn := func(a int) int {
		return 2
	}
	processIt(fn)
}
```
- function return function
```go
func processIt() func(a int) int {
	return func(a int) int {
		return 4
	}
}

func main(){
	fn := processIt()
	fn(6)
}
```

##

#Day13

### Variadic Functions
- The function that is called with the varying number of arguments is known as variadic function.
- like:
```go
fmt.Println(1, 2, 3, 4, 5, "hello")
```
- Let's make custom Variadic Functions
```go
package main

import "fmt"

func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total = total + num
	}
	return total
}

func main() {

	nums := []int{3, 4, 5, 6}
	result := sum(nums...)
	fmt.Println(result)
}
```

##

#Day13

### Closures in Go
- A function can access all variables defined inside the function, like this:
```go
func counter() func() int {
	var count int = 0

	return func() int {
		count += 1
		return count
	}
}


func main() {

	increment := counter()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
}
```