package main

import "fmt"

func main()  {
	m := map[string]int{
		"James": 42,
		"Moneypenny": 32,
		}
		
	for k, v := range m {
		fmt.Printf("key %v\t value %v\n", k, v)
	}	
	
	age1 := m["James"]
	fmt.Println("The age of bond ", age1)
	if v, ok := m["James"]; ok{
		fmt.Println("There is bond lookup entry, and bond's age is ", v)
	}

	age2 := m["Q"]
	fmt.Println("The age of Q ", age2)

	if v, ok := m["Q"]; !ok{
		fmt.Println("There is no Q, and here is the zero value of an init", v)
	}
}

// for range statement syntax
// for _, v := range v {	
	// }