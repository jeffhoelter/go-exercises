package main

import "fmt"

func main() {
	nums := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	// append 52
	nums = append(nums, 52)
	fmt.Println(nums)

	// append 53, 54, 55 in one statement
	nums = append(nums, 53, 54, 55)
	fmt.Println(nums)

	// append to the slice this slice
	newSlice := []int{56, 57, 58, 59, 60}
	nums = append(nums, newSlice...)
	fmt.Println(nums)
}
