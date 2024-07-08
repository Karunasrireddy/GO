package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Martini", "Chocolate"}
	jm := []string{"Jenny", "Moneypenny", "Guiness", "Wolverine Tracks"}
	fmt.Println(jb)
	fmt.Println(jm)
	xxs := [][]string{jb, jm}
	fmt.Println(xxs)
	fmt.Println("-----------------------------------------------")
	ja := []string{"James", "Bond", "Chocolate", "martini"}
	fmt.Println(ja)
	mp := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut"}
	fmt.Println(mp)
	xp := [][]string{ja, mp}
	fmt.Println(xp)
}

