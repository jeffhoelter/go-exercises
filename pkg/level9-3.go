package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Hands-on exercise #3
// Using goroutines, create an incrementer program
// 		have a variable to hold the incrementer value
// 		launch a bunch of goroutines
// 			each goroutine should
// 			read the incrementer value
// 				store it in a new variable
// 			yield the processor with runtime.Gosched()
// 			increment the new variable
// 			write the value in the new variable back to the incrementer variable
// use waitgroups to wait for all of your goroutines to finish
// the above will create a race condition.
// Prove that it is a race condition by using the -race flag
// if you need help, here is a hint: https://play.golang.org/p/FYGoflKQej

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
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
// WARNING: DATA RACE
// Write at 0x00c0000c8008 by goroutine 14:
//   main.main.func1()
//       /Users/jhoelter/dev/go/src/github.com/jeffhoelter/go-exercises/pkg/level9-3.go:42 +0x69

// Previous write at 0x00c0000c8008 by goroutine 10:
//   main.main.func1()
//       /Users/jhoelter/dev/go/src/github.com/jeffhoelter/go-exercises/pkg/level9-3.go:42 +0x69

// Goroutine 14 (running) created at:
//   main.main()
//       /Users/jhoelter/dev/go/src/github.com/jeffhoelter/go-exercises/pkg/level9-3.go:36 +0x23c

// Goroutine 10 (finished) created at:
//   main.main()
//       /Users/jhoelter/dev/go/src/github.com/jeffhoelter/go-exercises/pkg/level9-3.go:36 +0x23c
// ==================
// Goroutines: 2
// count: 1
// Found 4 data race(s)
// exit status 66
