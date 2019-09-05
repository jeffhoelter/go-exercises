package main

import (
	"fmt"
)

// Hands-on exercise #7
// Assign a func to a variable, then call that func

func main() {
	anonymous := func(s string) {
		fmt.Println("function assigned to variable with string:", s)
	}
	anonymous("private!")
}
