package main

import "fmt"

type person struct{
	first string
	last string
	favoriteIceCreamFlavors []string
}

func main()  {
	p1 := person{
		first : "James",
		last : "Bond",
		favoriteIceCreamFlavors : []string{"chocolate", "banana", "passion fruit with mango and guava"},
	}

	p2 := person{
		first : "Jenny",
		last : "Moneypenny",
		favoriteIceCreamFlavors : []string{"Strawberry", "Chocolate"},
	}
	fmt.Println(p1)
	fmt.Println(p2)
	for _, v := range p1.favoriteIceCreamFlavors {
		fmt.Println(p1.first, "favorite is", v)
	}
	for _, v := range p2.favoriteIceCreamFlavors {
		fmt.Println(p2.first, "favorite is", v)
	}
}