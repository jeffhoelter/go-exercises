package main

import (
	"fmt"
)
// Hands-on exercise #1
// get this code working:
// using a buffered channel


func main() {
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}
