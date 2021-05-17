package main

import (
	"fmt"
)

func searchInsert(nums []int, target int) int {
    var pos int
	numsLen := len(nums)
	for idx := 0; idx < numsLen; idx ++ {
		if nums[idx] >= target {
			pos = idx
			break
		} else {
			pos = idx + 1
		}
	}
	return pos
}

func main() {
	// arr := []int{1, 3, 5, 6}
	// target := 5

	// arr := []int{1, 3, 5, 6}
	// target := 2
	
	// arr := []int{1, 3, 5, 6}
	// target := 7
	
	// arr := []int{1, 3, 5, 6}
	// target := 0

	arr := []int{1}
	target := 0

	fmt.Println(searchInsert(arr, target))
}
