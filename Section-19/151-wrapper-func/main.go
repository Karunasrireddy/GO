package main 

import (
	"fmt"
	"os"
	"log"
)
func main()  {
	xb, err := readFile("SNOWY-EVENING.txt")
	if err != nil {
		log.Fatalf("error in main in readFile: %s", err) 
	}
	fmt.Println(xb)
	fmt.Println(string(xb))
}
func readFile(fileName string) ([]byte, error) {
	xb, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("there was an error in readfile: %s", err)
	}
	return xb, nil
}

