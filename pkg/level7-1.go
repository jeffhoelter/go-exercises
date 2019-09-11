package main

import "fmt"

// Hands-on exercise #1
// Create a variable and assign a value to it
// Print out the address of that variable

func main() {
	var x int
	x = 6

	fmt.Println(&x)
}
