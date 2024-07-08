package main

import "fmt"

func addI(a int, b int) int {
	return a + b
}

func addF(a float64, b float64) float64 {
	return a + b
}

// myNumber is an interface that allows any type with an underlying type of int or float64.
type myNumber interface {
	~int | ~float64
}

// Define a new type with int as the underlying type.
// Type alias
type myAlias int

func addT[T myNumber](a, b T) T {
	return a + b
}
func main() {
	var n myAlias = 42
	fmt.Println(addI(1, 2))
	fmt.Println(addF(1.2, 2.2))
	fmt.Println(addT(n, 2))
	fmt.Println(addT(1.2, 2.2))
}
