package main

import (
	"fmt"
	"sync"
)

// Hands-on exercise #1
// in addition to the main goroutine, launch two additional goroutines
// each additional goroutine should print something out
// use waitgroups to make sure each goroutine finishes before your program exists

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("First goroutine")
		wg.Done()
	}()

	go func() {
		fmt.Println("Second goroutine")
		wg.Done()
	}()

	wg.Wait()
}
