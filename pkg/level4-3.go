package main

import "fmt"

// Using the code from the previous example, use slicing to create the following new slices and print them:
// [42 43 44 45 46]
// [47 48 49 50 51]
// [44 45 46 47 48]
// [43 44 45 46 47]
func main() {
	nums := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	var firstSlice []int
	firstSlice = nums[:5]
	fmt.Println(firstSlice)

	secondSlice := nums[5:]
	fmt.Println(secondSlice)

	thirdSlice := nums[2:7]
	fmt.Println(thirdSlice)

	fourthSlice := nums[1:6]
	fmt.Println(fourthSlice)

}
