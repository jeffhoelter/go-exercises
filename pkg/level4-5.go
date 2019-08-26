package main

import "fmt"

// Deleting from a slice
// Start with this slice []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
// use append and slicing to make the slice [42, 43, 44, 48, 49, 50, 51]
func main() {
	nums := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	newNums := nums[:3]
	newNums = append(newNums, nums[6:10]...)
	fmt.Println(newNums)

}
