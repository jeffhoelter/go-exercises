package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Hands-on exercise #3
// Starting with this code, encode to JSON the []user sending the
// results to Stdout. Hint: you will need to use
// json.NewEncoder(os.Stdout).encode(v interface{})

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

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

	fmt.Println("String: _________________________")
	users := []user{u1, u2, u3}

	fmt.Println(users)

	// your code goes here
	fmt.Println("")
	fmt.Println("Encoded: _________________________")
	json.NewEncoder(os.Stdout).Encode(users)

}
