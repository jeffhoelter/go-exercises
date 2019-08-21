package main

import "fmt"

// Write a program that uses a for {} and prints out the years you have been alive
func main() {
	i := 1977
	for {
		if i > 2019 {
			break
		}
		fmt.Printf("%d\n", i)
		i++
	}
}
