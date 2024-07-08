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

	m := map[string]person{
		p1.last: p1,
		p2.last: p2, 
	}

	for _, v := range m {
		fmt.Printf("Value - %v\n", v)
		for _, v1 := range v.favoriteIceCreamFlavors {
			fmt.Println(v.first, v.last, v1)
		}		
	}
}