package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int{
	m := make(map[int]int)
	for i, val := range nums {
		sub := target - val
		j, ok := m[sub]
		if ok && i != j {
			return []int{i, j}
		}
		m[val] = i
	}
	return []int{}
}

func main() {
	// arr := []int{2, 7, 15, 11}
	// tar := 22

	// arr := []int{3, 2, 4}
	// tar := 6

	arr := []int{3, 3}
	tar := 6

	fmt.Println(twoSum(arr, tar))
}
