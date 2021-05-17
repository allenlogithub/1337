package main

import (
	"fmt"
)

// sb suggests kadane, and after some research:
func maxSubArray(nums []int) int {
	localMax, globalMax, numsLen := nums[0], nums[0], len(nums)
	for idx := 1; idx < numsLen; idx ++ {
		extLocalMax := nums[idx] + localMax
		if extLocalMax > nums[idx] {
			localMax = extLocalMax
		} else {
			localMax = nums[idx]
		}
		if localMax > globalMax {
			globalMax = localMax
		}
	}
	return globalMax
}

func main() {
	// nums := []int{-2,1,-3,4,-1,2,1,-5,4}

	// nums := []int{1}

	nums := []int{5,4,-1,7,8}

	fmt.Println(maxSubArray(nums))
}
