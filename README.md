![golang](https://go.dev/images/go-logo-white.svg)

# Golang

- Go is a statically typed, compiled high-level programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson. It is syntactically similar to C, but also has memory safety, garbage collection, structural typing, and CSP-style concurrency.

---

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

---

#Day2

### 5 Reasons to choose Golang

1. Build time
1. Fast starup
1. Performance and Efficiency
1. Concurrency Model
1. Static typing and compilation

---

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
