package main

import "fmt"

func add(a, b int)int {
	return a + b
}

func getLanguages() (string, string, bool) {
	return "golang", "javascript", true
}

func processIt(fn func(a int) int) {
	fn(1)
}

func processThis() func(a int) int {
	return func(a int) int {
		return 4
	}
}

func main() {
	a := 3
	b := 5

	fmt.Println(a, " + ", b, " = ", add(a, b))

	lang1, lang2, _ := getLanguages()
	fmt.Println(lang1, lang2)
	fmt.Println(getLanguages())

	fn := func(a int) int {
		return 2;
	}
	processIt(fn)

	fn2 := processThis()
	fmt.Println(fn2(3))
}