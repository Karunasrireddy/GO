package main

import (
	"io"
	"log"
	"os"
	"bytes"
	"fmt"
)
type person struct{
	first string
}
func (p person) writeOut(w io.Writer)  {
	w.Write([]byte(p.first))
	
}
func main()  {
	p := person{
		first : "Jenny",
	}
	// creating a file
	f, err := os.Create("Output.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}
	// closing a file
	defer f.Close()
	var b bytes.Buffer
	p.writeOut(f)
	p.writeOut(&b)
	fmt.Println(b.String())
}

