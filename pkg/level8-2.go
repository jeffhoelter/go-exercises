package main

import (
	"encoding/json"
	"fmt"
)

// Hands-on exercise #2
// Starting with this code, unmarshal the JSON into a Go data structure.

type generatedPerson []struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32222222222222222222222222,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println("String: _________________________")
	fmt.Println(s)

	var people generatedPerson

	if err := json.Unmarshal([]byte(s), &people); err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
	}

	fmt.Println("")
	fmt.Println("Struct: _________________________")
	for _, v := range people {
		fmt.Println(v)
	}

}
