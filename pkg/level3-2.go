package main

import "fmt"

// Write a program that prints every rune code point in the alphabet three times
func main() {
	const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for _, letter := range alphabet {
		fmt.Printf("U+00%x %+q\n", letter, letter)
	}
}
