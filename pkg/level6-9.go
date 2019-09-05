package main

import (
	"fmt"
	"strconv"
)

// Hands-on exercise #Hands-on exercise #9
// A “callback” is when we pass a func into a func as an argument.
// For this exercise, pass a func into a func as an argument

func main() {
	funcForFun(5, callbackFunction)
}

func funcForFun(timesToCall int, callback func(string)) {
	for i := 0; i < timesToCall; i++ {
		callback("hello " + strconv.Itoa(i))
	}
}

func callbackFunction(s string) {
	fmt.Println("I got called with string", s)
}
