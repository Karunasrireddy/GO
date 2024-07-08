package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Begin CPUs", runtime.NumCPU())
	fmt.Println("Begin Goroutines", runtime.NumGoroutine())
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("I am from Goroutines1")
		wg.Done()
	}()
	go func() {
		fmt.Println("I am from Goroutines2")
		wg.Done()
	}()

	fmt.Println("Mid CPUs", runtime.NumCPU())
	fmt.Println("Mid Goroutines", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("about to exit.")
	fmt.Println("End CPUs", runtime.NumCPU())
	fmt.Println("End Goroutines", runtime.NumGoroutine())
}
