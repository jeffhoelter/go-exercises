package main

import "fmt"

// Write a program that uses the following operations,
// write exressions and assign their values to variables,
// then print each of the variables.
// ==, <=, >= !=, <, >
func main() {

	equals := "hey" == "hey"
	lessThanEquals := "123" <= "123"
	greaterThanEquals := 5 >= 3
	doesntEqual := 3 != 4
	lessThan := 5 < 6
	greaterThan := 3 > 2
	fmt.Println(equals)
	fmt.Println(lessThanEquals)
	fmt.Println(greaterThanEquals)
	fmt.Println(doesntEqual)
	fmt.Println(lessThan)
	fmt.Println(greaterThan)
}
