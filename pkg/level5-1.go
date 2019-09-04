package main

import "fmt"

// Hands-on exercise #1
// Create your own type “person” which will have an underlying type of “struct” so that
// it can store the following data:
// first name
// last name
// favorite ice cream flavors
// Create two VALUES of TYPE person. Print out the values, ranging over the elements
// in the slice which stores the favorite flavors.

func main() {

	type person struct {
		firstName               string
		lastName                string
		favoriteIceCreamFlavors []string
	}

	person1 := person{
		firstName:               "William",
		lastName:                "Shakespeare",
		favoriteIceCreamFlavors: []string{"Hazelnut Chunk", "Mint Chocolate Chip"},
	}
	fmt.Println(person1)

	person2 := person{
		firstName:               "Lin",
		lastName:                "Manuel-Miranda",
		favoriteIceCreamFlavors: []string{"Cookie Dough", "Vanilla"},
	}
	fmt.Println(person2)
}
