package main

import "fmt"

type person struct{
	first string
	last string
	age int
}
func main()  {
	p1 := person{
		first: "James",
		last: "Bond",
		age: 32,
	}

	p2 := person{
		first: "Jenny",
		last: "Moneypenny",
		age: 42,
	}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Printf("type - %T\t value - %#v\n", p1, p1)
	fmt.Println(p1.first ,p1.last ,p1.age)
}