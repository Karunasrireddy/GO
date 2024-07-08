package main

import (
	"fmt"
	"math/rand"
)

func main()  {
	x := rand.Intn(10)
	y := rand.Intn(10)
	fmt.Printf("The value of x is %v and value of y is %v\n", x, y)

	switch  {
	case x < 4 && y < 4:
		fmt.Println("x and y are both less than 4")
	case x > 6 && y > 6:
		fmt.Println("x and y are both greater than 6")
	case x >= 4 && x <= 6:
		fmt.Println("x is from 4 to 6 inclusive of both numbers")
	case y != 5:
		fmt.Println("y is not 5")
	default:
		fmt.Println("None of the previous cases were met")
	}
	
}