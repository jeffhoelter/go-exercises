package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// Hands-on exercise #6
// Create a program that prints out your OS and ARCH.
// Use the following commands to run it:
// go run
// go build
// go install

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	var counter int64

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}

// Output:
// ==================
// ==================
// go run -race level9-5.go
// No race condition detected.
