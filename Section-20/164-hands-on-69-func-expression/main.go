package main
import "fmt"
func main()  {
	x := func ()  {
		fmt.Println("This is function expression")
		for i := 0; i < 5; i++ {
			fmt.Println("The value of i is ",i)
		}	
	}
	x()
	y()
}
var y = func ()  {
	fmt.Println("This is function expression")
	for i := 5; i < 10; i++ {
		fmt.Println("The value of i is ",i)
	}	
}