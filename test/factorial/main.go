package main

import (
	"fmt"
)

func factorial(num int) int {
	i, res := 1, 1
	for i < num {
		res *= i + 1
		i ++
	}
	return res
}

func main() {
	// num := 4

	// num := 0

	// num := 1

	num := 19

	fmt.Println(factorial(num))
}