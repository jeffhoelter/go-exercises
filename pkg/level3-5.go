package main

import "fmt"

// Write a program that prints out the remainder (modulus) for every number between
// 10 and 100 when it is divided by 4
func main() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("%d\n", i%4)
	}
}
