package main

import "fmt"

func main() {
	xi := make([]int, 0, 10)
	fmt.Println(xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))
	fmt.Println("-------------------------")
	xi = append(xi, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))
	fmt.Println("-------------------------")
	xi = append(xi, 10, 11, 12, 13)
	fmt.Println(xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))
	fmt.Println("-------------------------")
	x := make([]int, 10, 12)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x[0] = 42
	x[9] = 999
	fmt.Println("-------------------------")
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Println("-------------------------")
	x = append(x, 3423)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Println("-------------------------")
	x = append(x, 3424)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Println("-------------------------")
	x = append(x, 3425)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}

/*
	xi := make([]int, 0, 10)
	fmt.Println(xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))
	xi = append(xi, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Printf("xi - %#v\n", xi)
	fmt.Println("-------------")
*/
