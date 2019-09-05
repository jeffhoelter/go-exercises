package main

import (
	"bytes"
	"fmt"
)

// Hands-on exercise #10
// Closure is when we have “enclosed” the scope of a variable
// in some code block. For this hands-on exercise, create a
// func which “encloses” the scope of a variable:

func main() {
	s1 := stringAppender()
	fmt.Println(s1("woo"))

	s2 := stringAppender()
	fmt.Println(s2("Baltimore "))

}

func stringAppender() func(baseString string) string {
	stringToAppend := " I'm at the end!"
	return func(baseString string) string {
		buf := bytes.Buffer{}
		buf.WriteString(baseString)
		buf.WriteString(stringToAppend)
		return buf.String()
	}
}
