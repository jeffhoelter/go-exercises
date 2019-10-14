package main

import (
	"fmt"
)
// Hands-on exercise #1
// get this code working:
// using func literal, aka, anonymous self-executing func

func main() {
	c := make(chan int)

	go func() {c <- 42}()

	fmt.Println(<-c)
}
