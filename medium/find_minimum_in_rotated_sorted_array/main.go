package main

import (
	"fmt"
)

func findMin(nums []int) int {
    numsLen := len(nums)
	left, right := 0, numsLen - 1

	for left < right {
		mid := (left + right) / 2	
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}

func main() {
	nums := []int{3,4,5,1,2}

	fmt.Println(findMin(nums))
}