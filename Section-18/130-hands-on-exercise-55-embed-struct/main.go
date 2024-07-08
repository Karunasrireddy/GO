package main

import "fmt"

type engine struct{
	electric bool
}
type vehicle struct{
	engine 
	make string
	model string
	doors int
	color string

}
func main()  {
	v1 := vehicle{
		engine: engine{
			electric: true,
		},
		make:  "Ford",
		model: "Mustang",
		doors: 4,
		color: "Blue",
	}
	v2 := vehicle{
		engine: engine{
			electric: false,
		},
		make:  "Toyota",
		model: "Tundra",
		doors: 2,
		color: "White",
	}
	fmt.Printf("Values of v1 - %v\n", v1)
	fmt.Printf("Values of v2 - %v\n", v2)
	fmt.Println(v1.electric, v1.make, v1.model)
	fmt.Println(v2.electric, v2.make, v2.model)
	fmt.Println(v1.engine.electric, v1.make, v1.model)
	fmt.Println(v2.engine.electric, v2.make, v2.model)

}