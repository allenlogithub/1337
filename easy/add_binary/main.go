package main

import (
	"fmt"
)

func addBinary(a string, b string) string {
	aLen, bLen, diffLen := len(a), len(b), 0
    if aLen < bLen {
		temp := a
		a = b		
		b = temp
		diffLen = bLen - aLen
	} else {
		diffLen = aLen - bLen
	}	
	for diffLen > 0 {
		b = "0" + b
		diffLen --
	}
	carry, sum := 0, 0
	m := make(map[byte]int)
	m[48] = 0
	m[49] = 1
	var s string
	for idx := len(a) - 1; idx >= 0; idx -- {
		sum = m[a[idx]] + m[b[idx]] + carry
		if sum == 3 {
			carry = 1
			s = "1" + s
		} else if sum == 2 {
			carry = 1
			s = "0" + s
		} else if sum == 1 {
			carry = 0
			s = "1" + s
		} else {
			carry = 0
			s = "0" + s
		}
	}
	if carry == 1 {
		s = "1" + s
	}
	return s
}

func main() {
	// a := "11"
	// b := "1"

	// a := "1010"
	// b := "1011"

	a := "111"
	b := "1"

	// a := "1"
	// b := "111"

	fmt.Println(addBinary(a, b))
}
