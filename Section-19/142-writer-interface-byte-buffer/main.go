package main

import (
	// "log"
	// "os"
	"bytes"
	"fmt"
)
// type person struct{
// 	first string
// }
// func (p person) writeOut(w io.writer) error  {
// 	_, err := w.write([]byte(p.first))
// 	return err
// }
func main()  {
	// // creating a file
	// f, err := os.Create("Output.txt")
	// if err != nil {
	// 	log.Fatalf("error %s", err)
	// }
	// // closing a file
	// defer f.Close()
	// s := []byte("Hello Go!")
	// // writing to a file
	// _, err = f.Write(s)
	// if err != nil {
	// 	log.Fatalf("error %s", err)
	// }
	b := bytes.NewBufferString("Hello ")
	fmt.Println(b.String())
	b.WriteString("Gophers!")
	fmt.Println(b.String())
	b.Reset()
	b.WriteString("It's good! ")
	fmt.Println(b.String())
	b.Write([]byte("Happy Happy"))
	fmt.Println(b.String())
}

