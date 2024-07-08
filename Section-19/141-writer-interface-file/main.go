package main

import (
	"log"
	"os"
)
// type person struct{
// 	first string
// }
// func (p person) writeOut(w io.writer) error  {
// 	_, err := w.write([]byte(p.first))
// 	return err
// }
func main()  {
	// creating a file
	f, err := os.Create("Output.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}
	// closing a file
	defer f.Close()
	s := []byte("Hello Go!")
	// writing to a file
	_, err = f.Write(s)
	if err != nil {
		log.Fatalf("error %s", err)
	}
}

