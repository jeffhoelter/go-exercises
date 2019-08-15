package main

import "fmt"

type jefftype int

var x jefftype

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
