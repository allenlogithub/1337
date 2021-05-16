package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	count := 0
	n := len(nums)
	for idx := 1; idx < n; idx ++ {
		if nums[idx] == nums[idx - 1] {
			count ++
		} else {
			nums[idx - count] = nums[idx]
		}
	}
	return n - count
}

func main() {
	// arr := []int{1, 1, 2, 2, 3}
	arr := []int{0,0,1,1,1,2,2,3,3,4}

	fmt.Println(removeDuplicates(arr))
}
