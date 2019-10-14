package main

import (
	"fmt"
)

// Hands-on exercise #7
// write a program that
// launches 10 goroutines
// each goroutine adds 10 numbers to a channel
// pull the numbers off the channel and print them

func main() {
	c := make(chan int)

	for i := 0; i < 10; i++ {
		go func() {
			for v := 0; v < 10; v++ {
				c <- v
			}
		}()
	}

	for x := 0; x < 100; x++ {
		fmt.Println(x, <-c)
	}

	fmt.Println("main finished")

}
