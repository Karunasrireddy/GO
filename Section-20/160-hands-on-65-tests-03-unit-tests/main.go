package main

import "fmt"

func main() {
	fmt.Println(paradise("Hyd"))
}

func paradise(loc string) string {
	return fmt.Sprint("My idea of paradise is ",loc)
}
