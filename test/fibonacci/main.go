package main

import (
	"fmt"
)

func fibonacci(num int) int {
	if num == 0 || num == 1 {
		return num
	} else {
		idx, first, second := 1, 0, 1
		for idx < num {
			temp := first
			first = second
			second += temp
			idx ++
		}
		return second
	}

}

func main() {
	// num := 8

	// num := 3

	// num := 21

	num := 40

	fmt.Println(fibonacci(num))
}