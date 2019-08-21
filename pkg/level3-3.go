package main

import "fmt"

// Write a program that uses a for {} and prints out the years you have been alive
func main() {
	for i := 1977; i <= 2019; i++ {
		fmt.Printf("%d\n", i)
	}
}
