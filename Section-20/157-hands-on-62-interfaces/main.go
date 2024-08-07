package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
	width float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.width
}

func (c circle) area() float64{
	return math.Pi * math.Pow(c.radius,2)
}

type shape interface {
	area() float64
}

func info(sh shape) float64 {
	return sh.area() 
}
func main() {
	s1 := square{
		length : 2,
		width : 4,
	}

	c1 := circle{
		radius : 4,
	}
	
	info(s1)
	fmt.Println("Area of square is",info(s1))
	info(c1)
	fmt.Println("Area of cicle is",info(c1))
	
}
