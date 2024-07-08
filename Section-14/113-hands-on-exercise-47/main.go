package main

import "fmt"

func main()  {
	x := make([]string, 0, 50)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Println("---------------------------")
	x = append(x, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`,` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	for i := 0; i < len(x); i++ {
		fmt.Printf("index - %v\t value - %v\n", i, x[i])
	}
	/*
	x1 := make([]int, 50)
	fmt.Println(x1)
	fmt.Println(len(x1))
	fmt.Println(cap(x1))

	x2 := make([]int, 0, 50)
	fmt.Println(x2)
	fmt.Println(len(x2))
	fmt.Println(cap(x2))

	fmt.Println("---------------------------")
	//x1 = append(x1, 98)
	x1[0] = 98
	x2 = append(x2, 99)
	fmt.Println(x1)
	fmt.Println(len(x1))
	fmt.Println(cap(x1))

	fmt.Println(x2)
	fmt.Println(len(x2))
	fmt.Println(cap(x2))
	*/

}
