package main

import (
	"bytes"
	"fmt"
)

// Hands-on exercise #8
// Create a func which returns a func
// assign the returned func to a variable
// call the returned func

func main() {
	s1 := stringAppender("hoo")
	fmt.Println(s1("woo"))

	s2 := stringAppender("Orioles")
	fmt.Println(s2("Baltimore "))

}

func stringAppender(stringToAppend string) func(baseString string) string {
	return func(baseString string) string {
		buf := bytes.Buffer{}
		buf.WriteString(baseString)
		buf.WriteString(stringToAppend)
		return buf.String()
	}
}
