package main

import (
	"fmt"
)

// Hands-on exercise #4
// Show the comma ok idiom starting with this code.

func main() {
	c := make(chan int)

	go func() { c <- 1 }()

	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}
