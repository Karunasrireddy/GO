package main

import "fmt"

func main()  {
	x := foo()
	fmt.Println("The value of foo",x)

	i, s := bar()
	fmt.Println("The value of bar",i, s)
}

func foo() int {
	return 22
}

func bar() (int, string) {
	return 22, "James"
}