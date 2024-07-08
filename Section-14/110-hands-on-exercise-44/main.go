package main

import "fmt"

func main()  {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	// fmt.Printf("x - %v\n",x[:5])
	//[inclusive:exclusive]
	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])

}
