package main

import (
	"fmt"
)

// Hands-on exercise #6
// Build and use an anonymous func

func main() {
	func(s string) {
		fmt.Println("Anonymous function with string:", s)
	}("Private!")
}
