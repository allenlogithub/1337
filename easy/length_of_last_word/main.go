package main

import (
	"fmt"
)

func lengthOfLastWord(s string) int {
	l := 0
    for idx := len(s) - 1; idx >= 0; idx -- {
		if string(s[idx]) != " " {
			l ++
		} else if l != 0{
			return l
		}
	}
	return l
}

func main() {
	// s := "  Hello World"

	// s := "Hello World"

	// s := " "

	// s := "a"

	// s := "Today is a nice day"

	// s := "a "

	// s := "b   a    "

	fmt.Println(lengthOfLastWord(s))
}
