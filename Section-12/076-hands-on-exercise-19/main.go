package main

import (
	"fmt"
	"math/rand"
)

func main()  {
	x := rand.Intn(250)
	fmt.Printf("The value of x is %v\n",x)

	if x <= 100 {
		fmt.Println("The value is less than 100")
	} else if x >= 101 && x <= 200 {
		fmt.Println("The value is between 101 and 200")
	} else if x >= 201 && x <= 250 {
		fmt.Println("The value is between 201 and 250")
	} else {
		fmt.Println("The value is more than 250")
	}
	
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))

}