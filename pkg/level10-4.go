package main

import (
	"fmt"
)

// Hands-on exercise #4
// Starting with this code, pull the values off
// the channel using select

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func receive(c <-chan int, q <-chan int) {
	for {
		select {
		case i := <-c:
			fmt.Println("Received value: ", i)
		case i := <-q:
			fmt.Println("Quitting: ", i)
			return
		}
	}
}
func gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
		q <- 1
	}()

	return c
}
