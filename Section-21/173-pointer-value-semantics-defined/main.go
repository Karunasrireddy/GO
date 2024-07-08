package main

import "fmt"

// value semantics
func addOne(v int) int {
	return v + 1
}

// pointer semantics
func addOneP(v *int) {
	*v += 1
}

func main() {
	// value semantics
	a := 1
	fmt.Println("value semantics")
	fmt.Println(a)
	fmt.Println(addOne(a))
	fmt.Println(a)
	// pointer semantics
	b := 1
	fmt.Println("pointer semantics")
	fmt.Println(b)
	addOneP(&b)
	fmt.Println(b)

}
