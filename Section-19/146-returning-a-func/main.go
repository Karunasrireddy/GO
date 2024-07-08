package main

import "fmt"

func main()  {
	x := foo()
	fmt.Println(x)

	y := bar()
	fmt.Println(y())
	fmt.Printf("Type of foo - %T\n", foo)
	fmt.Printf("Type of x - %T\n", x)
	fmt.Printf("Type of bar - %T\n", bar)
	fmt.Printf("Type of y - %T\n", y)	
}
func foo() int {
	return 42
}
func bar() func() int {
	return func() int{
		return 43
	}
}

