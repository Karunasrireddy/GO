package main

import "fmt"

func addI(a int, b int) int {
	return a + b
}

func addF(a float64, b float64) float64 { return a + b }

// Type Set
type myNumber interface {
	int | float64
}

func addT[T myNumber](a, b T) T {
	return a + b
}
func main() {
	fmt.Println(addI(1, 2))
	fmt.Println(addF(1.2, 2.2))
	fmt.Println(addT(1, 2))
	fmt.Println(addT(1.2, 2.2))
}
