package main
import "fmt"
func main()  {
	x := foo()
	fmt.Println("The value of x is",x())
}

func foo() func() int {
	return func() int {
		return 42
	}
}