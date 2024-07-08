package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//SEQUENCE
	fmt.Println("this is the first statement to run")
	fmt.Println("this is the second statement to run")
	x := 40 // this is the third statement to run
	y := 5  // this is the fourth statement to run
	fmt.Printf(" x=%v \n y=%v\n", x, y)

	// CONDITIONAL
	// if statements
	// switch statements

	/*
		The expression [evaluated in an if statement] may be preceded by a simple statement,
		which executes before the expression is evaluated.
	*/
	// https://go.dev/ref/spec#If_statements

	if z := 2 * rand.Intn(x); z >= x {
		fmt.Printf("z is %v and that is GREATER THAN OR EQUAL x which is %v\n", z, x)
	} else {
		fmt.Printf("z is %v and that is LESS THAN x which is %v\n", z, x)
	}

	/*
		The comma ok idiom is also like this
	*/
	// https://go.dev/play/p/OXGzjxVkag0
	
	m := make(map[string]int)
	m["James"] = 42
	m["Moneypenny"] = 32
	fmt.Println("James is ", m["James"])
	fmt.Println("Miss Moneypenny is ", m["Moneypenny"])
	fmt.Println("Q is ", m["Q"])

	// comma ok idiom
	if v, ok := m["James"]; ok {
		fmt.Println("James is ", v)
	}

	if v, ok := m["Moneypenny"]; ok {
		fmt.Println("Miss Moneypenny is ", v)
	}

	// will not print
	if v, ok := m["Q"]; ok {
		fmt.Println("Q is ", v)
	}
}
