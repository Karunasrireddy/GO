package main

import (
	"fmt"
	"math/rand"
)

// func init()  {
// 	fmt.Println("This is where initialization for my program occurs")
// }
func main()  {
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

func init()  {
	fmt.Println("This is where initialization for my program occurs")
}