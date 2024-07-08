package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	xs := []string{"James", "Q", "M", "Moneypenny", "Dr. No"}
	fmt.Println("Unsorted int -", xi)
	sort.Ints(xi)
	fmt.Println("Sorted int -", xi)
	fmt.Println("*************************")
	fmt.Println("Unsorted string -", xs)
	sort.Strings(xs)
	fmt.Println("Sorted string -", xs)
}
