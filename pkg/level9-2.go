package main

import (
	"fmt"
)

// Hands-on exercise #2
// This exercise will reinforce our understanding of method sets:
// create a type person struct
// attach a method speak to type person using a pointer receiver
// 		*person
// create a type human interface
// 		to implicitly implement the interface, a human must have the speak method
// create func “saySomething”
// 		have it take in a human as a parameter
// 		have it call the speak method
// show the following in your code
// 		you CAN pass a value of type *person into saySomething
// 		you CANNOT pass a value of type person into saySomething
// here is a hint if you need some help
// https://play.golang.org/p/FAwcQbNtMG

// Receivers       Values
// -----------------------------------------------
// (t T)           T and *T
// (t *T)          *T

type person struct {
	firstName string
	lastName  string
	age       int
}

func (p *person) speak() {
	fmt.Println("Hello, my name is", p.firstName)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{
		firstName: "Sally",
		lastName:  "Jones",
		age:       12,
	}

	// saySomething(p)
	// this doesn't work, error is:
	// ./level9-2.go:54:14: cannot use p (type person) as type human in argument to saySomething:
	// person does not implement human (speak method has pointer receiver)

	saySomething(&p)

}
