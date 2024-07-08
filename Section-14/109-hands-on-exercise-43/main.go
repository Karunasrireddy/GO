package main

import "fmt"

func main()  {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for i, v := range x {
		fmt.Printf("value - %v\t type - %T\t index - %v\n", v, v, i)	
	}
	fmt.Printf("type - %T\t value - %#v\n", x, x)
	fmt.Printf("type - %T\t value - %v\n", x, x)
}