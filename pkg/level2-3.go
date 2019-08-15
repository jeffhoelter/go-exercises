package main

import "fmt"

// Write a program that creates types and untyped constants and prints the values
func main() {
	const typedConstant int = 4
	const untypedConstant = "hey"

	fmt.Printf("%T\n", typedConstant)
	fmt.Println(typedConstant)

	fmt.Printf("%T\n", untypedConstant)
	fmt.Println(untypedConstant)
}
