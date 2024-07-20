package main

import (
	"fmt"
	"maps"
)

func main() {
	// maps -> hash tables, object, dict
	// creating map

	// m := make(map[string]string)

	// setting an element
	// m["name"] = "golang"
	// m["area"] = "backend"

	// get an element
	// fmt.Println(m["name"], m["area"])
	// IMP: if key does not exists in the map then it returns zero value

	// m := make(map[string]int)
	// m["age"] = 30
	// m["price"] = 50
	// fmt.Println(m["age"])

	// to delete
	// delete(m, "price")
	// to clear map
	// clear(m)
	// fmt.Println(len(m))

	// create map with make func
	// m := map[string]int{"price": 40, "age": 20, "phones": 3}
	// fmt.Println(m)

	// k, ok := m["age"]
	// fmt.Println(k)

	// if ok {
	// 	fmt.Println("all ok")
	// } else {
	// 	fmt.Println("not ok")
	// }

	m1 := map[string]int{"price": 40, "age": 20, "phones": 3}
	m2 := map[string]int{"price": 40, "age": 20, "phones": 4}

	fmt.Println(maps.Equal(m1, m2))

}
