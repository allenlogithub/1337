package main

import (
	"fmt"
)

func sort(nums []int) []int {
	idx, numsLen := 1, len(nums)
	for idx < numsLen {
		subIdx := idx		
		for nums[subIdx] < nums[subIdx - 1] && subIdx > 0 {
			fmt.Println(subIdx, nums)
			temp := nums[subIdx - 1]
			nums[subIdx - 1] = nums[subIdx]
			nums[subIdx] = temp
			if subIdx > 1 {
				subIdx --
			}			
		}
		idx ++
	}
	return nums
}

func main() {
	// nums := []int{1,3,2,4,1}

	// nums := []int{3,7,1,4,6,2,5}

	nums := []int{4,3,2,1}

	fmt.Println(sort(nums))
}