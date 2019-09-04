package main

import "fmt"

// Hands-on exercise #2
// Take the code from the previous exercise, then store
// the values of type person in a map with the key of last
// name. Access each value in the map. Print out the values,
// ranging over the slice.

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

	var m map[string]person
	m = make(map[string]person)
	m[person1.lastName] = person1
	m[person2.lastName] = person2

	for _, value := range m {
		fmt.Println("Value:", value)
	}

}
