package main

import "fmt"

func main()  {
	x := 42
	fmt.Println(x)
	fmt.Println(&x)
	y := &x
	fmt.Println(*y)
	fmt.Printf("value - %v\t type - %T\n", *y, y)
	fmt.Printf("value - %v\t type - %T\n", y, y)

	*y = 43
	fmt.Printf("value - %v\t type - %T\n", x, x)
	fmt.Println(*y)

}