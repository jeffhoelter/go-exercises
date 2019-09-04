package main

import "fmt"

// Hands-on exercise #4
// Create and use an anonymous struct.

func main() {
	anonymousStruct := struct {
		NESCarts      []string
		numberOfCarts int
	}{
		[]string{"Super Mario Bros.", "Punch Out!"},
		2,
	}
	fmt.Println(anonymousStruct)
}
