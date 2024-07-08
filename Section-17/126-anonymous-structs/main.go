package main

import "fmt"

type person struct{
	first string
	last string
	age int
}
func main()  {
	p1 := struct{
		first string
		last string
		age int
	}{
		first: "James",
		last: "Bond",
		age: 32,
	}

	p2 := person{
		first: "Jenny",
		last: "Moneypenny",
		age: 42,
	}
	fmt.Printf("type - %T\n", p1)
	fmt.Printf("type - %T\n", p2)
}