package main

import "fmt"

// Using a compose literal, create an array which holds 5 values of type int - assign values
// to each index position.
// Range over the array and print the values out
// Using format printing, print out the type of the array
func main() {
	nums := [5]int{2, 3, 4, 5, 6}
	for _, num := range nums {
		fmt.Println(num)
	}
	fmt.Printf("%T\n", nums)
}
