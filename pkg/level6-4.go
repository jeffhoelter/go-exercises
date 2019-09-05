package main

import "fmt"

// Hands-on exercise #4
// Create a user defined struct with
// the identifier “person”
// the fields:
// first
// last
// age
// attach a method to type person with
// the identifier “speak”
// the method should have the person say their name and age
// create a value of type person
// call the method from the value of type person

type person struct {
	first string
	last  string
	age   int
}

func main() {
	fmt.Println("hello")
	person1 := person{
		first: "William",
		last:  "Shakespeare",
		age:   400,
	}
	fmt.Println(person1)
	person1.Speak()

}

func (p person) Speak() {
	fmt.Printf("My name is %s %s", p.first, p.last)
	fmt.Printf(" and I am %d years old.\n", p.age)
}
