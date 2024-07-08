package main
import "fmt"
func main()  {
	func ()  {
		fmt.Println("This is an anonymous function")
		for i := 0; i < 10; i++ {
			fmt.Println("The value of i is ",i)
		}	
	}()
}