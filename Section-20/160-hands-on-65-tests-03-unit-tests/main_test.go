package main

import (
	"testing"
	"log"
)

func TestParadise(t * testing.T)  {
	got := paradise("Hyd")
	want := "My idea of paradise is Hyd"
	if got != want {
		log.Fatalf("error - want %s and got %s", want, got)
	}
}
