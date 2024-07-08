package main

import "fmt"

func main()  {
	foo()
	func ()  {
		fmt.Println("Anonymous func ran")
	}()
	// with parameter(s)
	func (s string)  {
		fmt.Println("This is an anonymous func showing my name", s)
	}("Todd")
	// Function with a return
	s := func (s string) string {
		fmt.Println("This function has a return")
		return fmt.Sprint("The book is ",s)
	}("yoga for everyone")
	fmt.Println(s)
}
func foo()  {
	fmt.Println("foo ran")
}

// a named function with an identifier
// func (r receive) identifier(p parameter(s)) (r return(s)) { code }

// an anonymous function
// func(p parameter(s)) (r return(s)) { code }
