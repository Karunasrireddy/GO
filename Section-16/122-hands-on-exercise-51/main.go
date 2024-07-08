package main

import "fmt"

func main()  {
	a := make(map[string][]string)
		a[`bond_james`] = []string{`shaken, not stirred`, `martinis`, `fast cars`}
		a[`moneypenny_jenny`] = []string{`intelligence`, `literature`, `computer science`}
		a[`no_dr`] = []string{`cats`, `ice cream`, `sunsets`}
		a[`fleming_ian`] = []string{`steaks`, `cigars`, `espionage`}
		for k, v := range a {
			fmt.Println(k)
			for i, v1 := range v {
				fmt.Printf("index - %v\t value - %v\n", i, v1)
			}
		}
		fmt.Println("------------record deleted----------------")
		delete(a, `moneypenny_jenny`)
		for k, v := range a {
			fmt.Println(k)
			for i, v1 := range v {
				fmt.Printf("index - %v\t value - %v\n", i, v1)
			}
		}
 }
