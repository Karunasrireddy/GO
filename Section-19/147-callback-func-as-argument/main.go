package main

import "fmt"

func main()  {
	fmt.Printf("Type of add - %T\n", add)
	fmt.Printf("Type of subtract - %T\n", subtract)	
	fmt.Printf("Type of doMath - %T\n", doMath)

	x := doMath(42, 16, add)
	fmt.Println(x)

	y := doMath(42, 16, subtract)
	fmt.Println(y)
}

func doMath(a int, b int, f func(int, int)int) int {
	return f(a, b)
}

func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}