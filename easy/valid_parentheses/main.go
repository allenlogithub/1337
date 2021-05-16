package main

import (
	"fmt"
)

func isValid(s string) bool {
	res := true
	var arr []string
	for _, symb := range(s) {
		if string(symb) == "(" {
			arr = append(arr, ")")
		} else if string(symb) == "[" {
			arr = append(arr, "]")
		} else if string(symb) == "{" {
			arr = append(arr, "}")
		} else if string(symb) == ")" {
			if len(arr) == 0 {
				res = false
				break
			}
			if arr[len(arr) - 1] != ")" {
				res = false
				break
			} else {
				arr = arr[:len(arr) - 1]				
			}
		} else if string(symb) == "]" {
			if len(arr) == 0 {
				res = false
				break
			}
			if arr[len(arr) - 1] != "]" {
				res = false
				break
			} else {
				arr = arr[:len(arr) - 1]
			}
		} else if string(symb) == "}" {
			if len(arr) == 0 {
				res = false
				break
			}
			if arr[len(arr) - 1] != "}" {
				res = false
				break
			} else {
				arr = arr[:len(arr) - 1]
			}
		}		
	}

	if len(arr) != 0 {
		res = false
	}
	
	return res
}

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("{[]}"))

	fmt.Println(isValid("("))
	fmt.Println(isValid("(("))

	fmt.Println(isValid(")"))
}
