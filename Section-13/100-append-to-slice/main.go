package main

import "fmt"

func main()  {
	xi := []int{42, 43, 44}
	fmt.Println(xi)
	fmt.Println("---------------------")
	// variadic parameter
	xi = append(xi, 45, 46, 47, 48, 99, 777)
	fmt.Println(xi)
	fmt.Println("***********************")
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
	fmt.Println("---------------------")
	x = append(x, 77, 88, 99, 1014)
	fmt.Println(x)
	fmt.Println("***********************")
	y := []int{234, 456, 678, 987}
	// variadic parameter
	x = append(x, y...)
	fmt.Println(x)
}