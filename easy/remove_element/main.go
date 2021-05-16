package main

import (
	"fmt"
)

func removeElement(nums []int, val int) int {
	count := 0
	n := len(nums)
    for idx := 0; idx < n; idx ++ {
		if nums[idx] == val {
			count ++
		} else {
			nums[idx - count] = nums[idx]
		}
	}
	return n - count
}

func main() {
	// val := 3
	// arr := []int{3, 2, 2, 3}

	val := 2
	arr := []int{0,1,2,2,3,0,4,2}

	fmt.Println(removeElement(arr, val))
}
