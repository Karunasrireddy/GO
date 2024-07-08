package main

import "fmt"

func main()  {
	i := 5
	fmt.Println("Decrement")
	for i > 0{
		fmt.Printf("The value of i is %v\n", i)
		i--
	}

	x := 0
	fmt.Println("\nIncrement")
	for x <= 5{
		fmt.Printf("The value of x is %v\n", x)
		x++
	}
	
}