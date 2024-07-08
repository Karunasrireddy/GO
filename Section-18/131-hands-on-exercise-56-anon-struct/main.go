package main

import "fmt"

func main()  {
	a := struct {
		first string
		friends map[string]int
		favDrinks []string
	}{
		first : "James",
		friends : map[string]int{
			"Jenny": 27,
			"Q":     87,
			"Ian":   47,
		},
		favDrinks: []string{
			"Martini",
			"Water",
		},
	}
	for k, v := range a.friends {
		fmt.Println(a.first, "- friends - ", k, v)
	}

	for _, v := range a.favDrinks {
		fmt.Println(a.first, "- drinks - ", v)
	}
}