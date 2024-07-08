package main

import "fmt"

func main() {
	xi := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("xi - %#v\n", xi)
	fmt.Println("-------------")

	// [inclusive:exclusive]
	fmt.Printf("xi - %#v\n", xi[0:4])
	fmt.Println("-------------")

	// [:exclusive]
	fmt.Printf("xi - %#v\n", xi[:7])
	fmt.Println("-------------")

	// [inclusive:]
	fmt.Printf("xi - %#v\n", xi[4:])
	fmt.Println("-------------")

	// [:]
	fmt.Printf("xi - %#v\n", xi[:])
	fmt.Println("-------------")

	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x[1])
	fmt.Println(x)
	fmt.Println(x[1:])
	fmt.Println(x[1:3])
	fmt.Println("-------------")
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Println("-------------")
	for i := 0; i <= 4; i++ {
		fmt.Println(i, x[i])
	}
}
