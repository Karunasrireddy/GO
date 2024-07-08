package main

import (
	"fmt"
	"math/rand"
)

func main()  {
	// x := rand.Intn(230)
	// x := rand.Intn(330)
	// x := rand.Intn(430)
	x := rand.Intn(250)
	fmt.Printf("The value of x is %v\n",x)

	switch {
	case x <= 100:
		fmt.Println("The value is less than 100")
	case x >= 101 && x <= 200:
		fmt.Println("The value is between 101 and 200")
	case x >= 201 && x <= 250:
		fmt.Println("The value is between 201 and 250")
	default:
		fmt.Println("The value is more than 250")
	}
}