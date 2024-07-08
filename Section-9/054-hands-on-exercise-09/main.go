package main

import "fmt"

var a = 10
const b = "Hello"

func main()  {
	fmt.Printf("The value of a is %v and the type of a is %T\n" ,a,a)
	fmt.Printf("The value of b is %v and the type of b is %T\n" ,b,b)
	c := 30
	fmt.Printf("The value of c is %v and the type of c is %T\n" ,c,c)
}