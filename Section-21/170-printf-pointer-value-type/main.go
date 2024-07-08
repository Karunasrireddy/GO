package main

import "fmt"

func main()  {
	x := 42
	// y := &x
	// fmt.Println(*y)
	// fmt.Printf("value - %v\t type - %T\n", y, y)
	fmt.Println(x)
	fmt.Println(&x)
	fmt.Printf("value - %v\t type - %T\n", &x, &x)

	s := "James"
	fmt.Printf("value - %v\t type - %T\n", &s, &s)

}