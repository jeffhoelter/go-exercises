package main

import (
	"fmt"
)

// Hands-on exercise #3
// Starting with this code, pull the values off
// the channel using a for range loop

func main() {
	c := gen()
	receive(c)

	fmt.Println("about to exit")
}

func receive(c <-chan int) {
	for elem := range c {
		fmt.Println(elem)
	}
}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}
