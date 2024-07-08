package main

import (
	"testing"
	"log"
)

func TestAdd(t * testing.T)  {
	got := add(2,5)
	want := 7
	if got != want {
		log.Fatalf("error - want: %v and got: %v", got, want)
		// t.Errorf("Sum was incorrect, got: %d, want: %d.", got, want)
	}
}

func TestSubtract(t * testing.T)  {
	got := subtract(5,2)
	want := 3
	if got != want {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", got, want)
	}
}

func TestDomath(t * testing.T)  {
	got := doMath(2,5,add)
	want := 7
	if got != want {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", got, want)
	}
	
}