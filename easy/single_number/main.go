package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i ++ {
		res ^= nums[i]
	}
	return res
}

func main() {
	// nums := []int{2,2,1}

	nums := []int{4,1,2,1,2}

	// nums := []int{1}

	fmt.Println(singleNumber(nums))
}
