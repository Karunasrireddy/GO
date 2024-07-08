package main

import "fmt"

func main()  {
	x := []string{"James", "Bond", "Shaken, not stirred"}
	y := []string{"Miss", "Moneypenny", "I'm 008."}
	fmt.Println(x)
	fmt.Println(y)
	xp := [][]string{x, y}
	fmt.Println(xp)
	for i, v := range xp {
		fmt.Println("--------------------------------")
		fmt.Printf("index - %v\t value - %v\n", i, v)
		for a, b := range v {
			fmt.Printf("index - %v\t value - %v\n", a, b)
		}
	}
}
