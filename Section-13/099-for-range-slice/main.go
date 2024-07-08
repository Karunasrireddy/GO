package main

import "fmt"

func main()  {
	xs := []string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)"}
	fmt.Println(xs)
	fmt.Printf("%T\n",xs)
	fmt.Println(len(xs))
	fmt.Println("------------------------")
	// blank identifier to not use a returned value
	for _, v := range xs {
		fmt.Printf("%v\n",v)
	}
	// Access the slice values by index position
	fmt.Println("------------------------")
	fmt.Println(xs[0])
	fmt.Println(xs[1])
	fmt.Println(xs[2])
	// doesn't work
	// fmt.Println(xs[3])
	fmt.Println("------------------------")
	for i := 0; i < 3; i++ {
		fmt.Println(xs[i])
	}
	fmt.Println("------------------------")
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(len(x))
	fmt.Println(x)
	fmt.Println("------------------------")
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x[4])
	fmt.Println("------------------------")
	for i, v := range x {
		fmt.Println(i, v)	
	} 

}