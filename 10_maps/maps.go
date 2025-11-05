package main

import (
	"fmt"
	"maps"
)

// maps -> hash, object, dict
func main() {
	// creating map

	// m := make(map[string]string)

	// // setting an element
	// m["name"] = "golang"
	// m["area"] = "backend"

	// // getting an element
	// fmt.Println("m[\"name\"]: ", m["name"])
	// fmt.Println("m[\"area\"]: ", m["area"])

	// //! IMP: if `key` does not exists in the map then it return default zero value
	// fmt.Println("m[\"phone\"]: ", m["phone"])

	// m := make(map[string]int)

	// m["age"] = 30
	// m["price"] = 50

	// fmt.Println("m[\"age\"]: ", m["age"])
	// fmt.Println("m[\"phone\"]: ", m["phone"])

	// fmt.Println("m[] length: ", len(m))
	// fmt.Println("m[]: ", m)
	
	// delete(m, "price")
	// fmt.Println("m[]: ", m)

	// clear(m)
	// fmt.Println("m[]: ", m)

	// m := map[string]int{"price": 40, "age": 20}
	
	// fmt.Println("m[]: ", m)

	// m := map[string]int{"price": 40, "age": 20}

	// // Good practice to access elements
	// val, ok := m["price"]
	// if ok {
	// 	fmt.Println("all ok")
	// 	fmt.Println("val: ", val)
	// }else {
	// 	fmt.Println("not ok")
	// }

	m1 := map[string]int{"price": 40, "phones": 3}
	m2 := map[string]int{"price": 40, "phones": 2}

	fmt.Println("Is m1 == m2: ", maps.Equal(m1, m2))
}