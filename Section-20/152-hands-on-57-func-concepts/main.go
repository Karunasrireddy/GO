package main

import "fmt"

func main()  {
	xi := sum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println(xi)
}
func sum(x []int) int {
		total := 0
		for _, v := range x {
			total += v
		}
		return total
	}

// // named return
// func sum(x []int) (total int) {
// 	total = 0
// 	for _, v := range x {
// 		total += v
// 	}
// 	return
// }
