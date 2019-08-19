package main

import "fmt"

// Using iota, create four constants for the last four years and print the values
func main() {
	const (
		fourYearsAgo = iota + 2015
		threeYearsAgo
		twoYearsAgo
		lastYear
	)
	fmt.Println(fourYearsAgo)
	fmt.Println(threeYearsAgo)
	fmt.Println(twoYearsAgo)
	fmt.Println(lastYear)
}
