package main

import (
	"fmt"
	"runtime"
)

// Hands-on exercise #6
// Create a program that prints out your OS and ARCH.
// Use the following commands to run it:
// go run
// go build
// go install

func main() {
	fmt.Println("Architecture:", runtime.GOARCH)
	fmt.Println("OS:", runtime.GOOS)
}
