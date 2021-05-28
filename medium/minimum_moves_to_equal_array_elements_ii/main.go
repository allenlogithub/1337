package main

import (
	"fmt"
	"sort"
)

func minMoves2(nums []int) int {
    sum := 0
	for _, num := range(nums) {
		sum += num
	} 

	// avg := 0
	// mod := sum % len(nums)
	// if mod != 0 {
	// 	sum -= mod
	// 	avg = sum / len(nums)
	// } else {
	// 	avg = sum / len(nums)
	// }

	// avg doesn't work, change to median
	sort.Sort(sort.IntSlice(nums))
	pos := len(nums) / 2
	med := nums[pos]
	moves := 0
	for _, num := range(nums) {
		val := num - med
		if val > 0 {
			moves += val
		} else {
			moves -= val
		}
	}
	return moves
}

func main() {
	// nums := []int{1,10,2,9}

	// nums := []int{1,2,3}

	nums := []int{1,0,0,8,6}

	fmt.Println(minMoves2(nums))
}

