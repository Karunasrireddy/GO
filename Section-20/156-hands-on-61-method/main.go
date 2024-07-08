package main

import "fmt"

type person struct{
	first string
	age int
}

func (p person) speak() {
	fmt.Println("I am", p.first, "and my age is", p.age)
}
func main() {
	p1:= person{
		first : "James",
		age : 24,
	}
	p1.speak()
	
}
