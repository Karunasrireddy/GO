package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func addI(a int, b int) int {
	return a + b
}

func addF(a float64, b float64) float64 {
	return a + b
}

type myNumber interface {
	constraints.Integer | constraints.Float
}

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
