package main

import "fmt"

// Using a compose literal, create a slice that contains 10 values of type int - assign values
// to each index position.
// Range over the array and print the values out
// Using format printing, print out the type of the slice
func main() {
	nums := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for _, num := range nums {
		fmt.Println(num)
	}
	fmt.Printf("%T\n", nums)
}
