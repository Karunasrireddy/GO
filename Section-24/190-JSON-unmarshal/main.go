package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":23},{"First":"Miss","Last":"Moneypenny","Age":34}]`
	bs := []byte(s)
	fmt.Printf("Type of s is - %T\n", s)
	fmt.Printf("Type of bs is - %T\n", bs)

	// people := []person{}
	var people []person

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\nAll of the data -", people)

	for i, v := range people {
		fmt.Println("\nPERSON NUMBER", i)
		fmt.Println(v.First, v.Last, v.Age)
		fmt.Println(v)
	}
}
