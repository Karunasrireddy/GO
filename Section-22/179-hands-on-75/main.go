package main

import "fmt"

var (
	a, b, c *string
	d       *int
)

func init() {
	p := "Drop by drop, the bucket gets filled."
	q := "Persistently, patiently, you are bound to succeed."
	r := "The meaning of life is ..."
	n := 42
	a = &p
	b = &q
	c = &r
	d = &n
}
func main() {
	fmt.Printf(" Value(address) of a - %v\n Value(address) of b - %v\n", a, b)
	fmt.Printf(" Value(address) of c - %v\n Value(address) of d - %v\n", c, d)
	fmt.Printf(" Type of a is - %T\n Type of b is - %T\n Type of c is - %T\n Type of d is - %T\n", a, b, c, d)
	fmt.Println("Data stored at memory locations -", *a)
	fmt.Println("Data stored at memory locations -", *b)
	fmt.Println("Data stored at memory locations -", *c)
	fmt.Println("Data stored at memory locations -", *d)
}
