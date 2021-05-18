package main

import (
	"fmt"
)

func climbStairs(n int) int {
    pre, cur := 0, 1
	if n == 1 {
		return 1
	}
	for n > 1 {
		temp := cur
		cur = pre + cur
		pre = temp
		n --
	}
	return pre + cur
}

func main() {
	// n := 1
	// 1

	// n := 2 
	// 11, 2

	// n := 3
	// 111, 12, 21

	// n := 4
	// 1111, 112, 121, 211, 22

	n := 5

	fmt.Println(climbStairs(n))
}
