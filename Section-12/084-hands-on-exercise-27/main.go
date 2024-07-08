package main

import "fmt"

func main()  {
	i := 0
	fmt.Println("Increment")
	for {
		fmt.Printf("The value of i is %v\n", i)
		if i >= 5{
			break
		}
		i++
	}

	x := 10
	fmt.Println("\nDecrement")
	for {
		fmt.Printf("The value of x is %v\n", x)
		if x < 5{
			break
		}
		x--
	}
}