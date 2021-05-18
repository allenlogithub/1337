package main

import (
	"fmt"
)

func mySqrt(x int) int {
	// binary search
	if x == 0 || x == 1 {
		return x
	}
    low, high := 0, x
	for low < high {
		mid := int(low + (high - low) / 2)
		if mid > x /mid {
			high = mid - 1
		} else {
			low = mid + 1
			if low > x / low {
				return mid
			}
		}
	}
	return low
	
	// xx := float64(x)
	// ans := xx
	// for ans * ans - xx > 0.0000001 {
	// 	ans = (ans * ans + xx) / (ans * 2)
	// }
	// return int(ans)
}

func main() {
	// x := 5

	// x := 4

	// x := 1

	// x := 16

	// x := 0

	// x := 3

	x := 8

	// x := 99

	fmt.Println(mySqrt(x))
}
