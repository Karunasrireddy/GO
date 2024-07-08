package main

import "fmt"

func main()  {
	a := [5]int{}
	a[0] = 1
	a[1] = 2
	a[2] = 3
	a[3] = 4
	a[4] = 5

	for i, v := range a {
		fmt.Printf("index - %v\t value - %v\t type - %T\n", i, v, v)
	}

	/*
	// create an array
	x := [5]int{}
	// assign values to each index position
	for i := 0; i < 5; i++ {
		x[i] = i
	}
	// print out
	for i, v := range x {
		fmt.Printf("%v - %T - %v\n", v, v, i)
	}
	*/

}