package main

import (
	"fmt"
	"sort"
)

// Hands-on exercise #5
// Starting with this code, sort the []user by
// age
// last
// Also sort each []string “Sayings” for each user
// print everything in a way that is pleasant

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

// override String() function to print each user object in a pleasant way
func (u user) String() string {
	s := fmt.Sprintf("Name: %s %s\n", u.First, u.Last)
	s = s + fmt.Sprintf("Age: %d\n", u.Age)
	for i, v := range u.Sayings {
		s = s + fmt.Sprintf("Saying %d: %s\n", i, v)
	}
	return s
}

// ByAge implements sort.Interface based on the Age field.
type ByAge []user

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// ByLastName implements sort.Interface based on the Last NAme field.
type ByLastName []user

func (a ByLastName) Len() int           { return len(a) }
func (a ByLastName) Less(i, j int) bool { return a[i].Last < a[j].Last }
func (a ByLastName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println("__________________________________________________")
	fmt.Println("Unsorted:")
	fmt.Println(users)

	// your code goes here
	sort.Sort(ByAge(users))

	fmt.Println("__________________________________________________")
	fmt.Println("Sorted by Age:")
	fmt.Println(users)

	sort.Sort(ByLastName(users))

	fmt.Println("__________________________________________________")
	fmt.Println("Sorted by Last Name:")
	fmt.Println(users)

	for _, v := range users {
		sort.Strings(v.Sayings)
	}
	fmt.Println("__________________________________________________")
	fmt.Println("Sorted by Last Name and Sayings:")
	fmt.Println(users)

}
