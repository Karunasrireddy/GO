package main

import "fmt"

func main() {
	a := []string{"hello", "world"}
	fmt.Println(a)

	// COMPOSITE LITERAL
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
}

// a SLICE allows you to group together VALUES of the same TYPE
/*
	slice is a data structure with three elements:
	-- len
	-- cap
	-- pointer to an underlying array

	type slice struct {
		array unsafe.Pointer
		len   int
		cap   int
	}
*/
// src/runtime/slice.go
// google: "golang pkg runtime" then click on the "slice.go" file at the bottom of the page

