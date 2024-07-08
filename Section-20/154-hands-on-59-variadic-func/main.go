package main

import "fmt"

func main()  {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Function foo",foo(xi...))
	fmt.Println("Function bar",bar(xi))

}

func foo(i ...int) int {
	t := 0
	for _, v := range i {
		t += v
	}
	return t
	
}

func bar(i []int) int {
	t := 0
	for _, v := range i {
		t += v
	}
	return t
	
}