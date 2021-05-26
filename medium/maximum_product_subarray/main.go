package main

import (
	"fmt"
)

func maxProduct(nums []int) int {
	curMax, curMin, res := 1, 1, max(nums)
	for _, num := range(nums) {
		if num == 0 {
			curMax, curMin = 1, 1
			continue
		}
		temp := curMax * num
		arr := []int{temp, curMin * num, num}
		curMax = max(arr)
		curMin = min(arr)
		arr = []int{curMax, res}
		res = max(arr)
	}
	return res
}

func max(nums []int) int {
	res := nums[0]
	for _, num := range(nums) {		
		if num > res {
			res = num
		} 
	}
	return res
}

func min(nums []int) int {
	res := nums[0]
	for _, num := range(nums) {				
		if num < res {
			res = num
		} 
	}
	return res
}

func main() {
	// nums := []int{2,3,-2,4}

	// nums := []int{-2}

	nums := []int{-3,-1,-1}

	// nums := []int{-2,0,-1}

	fmt.Println(maxProduct(nums))
}

