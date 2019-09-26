package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Hands-on exercise #4
// Fix the race condition you created in the previous exercise by using a mutex
// it makes sense to remove runtime.Gosched()

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			v++
			counter = v
			mu.Unlock()
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
// go run -race level9-4.go
// No race condition detected.
