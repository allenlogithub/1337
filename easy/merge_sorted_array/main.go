package main

import (
	"fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	// fmt.Println(nums1, m, n)
	for m > 0 && n > 0 {
		if nums1[m - 1] >= nums2[n - 1] {
			nums1[m + n - 1] = nums1[m - 1]
			m --
		} else { // last of nums2 > last of nums1
			nums1[m + n - 1] = nums2[n - 1]
			n --
		}
		// fmt.Println(nums1, m, n)		
	}
	if n > 0 {
		// nums1 = append(nums2[:n], nums1[n:]...)
		// fmt.Println(nums1[:n], nums2[:n])
		for i := 0; i < len(nums2[:n]); i++ {
			nums1[i] = nums2[i]
		}
		// nums1 = copy(nums1[:n], nums2[:n])
	}
	return nums1
}

func main() {
	// nums1 := []int{1,2,3,0,0,0}
	// m := 3
	// nums2 := []int{2,5,6}
	// n := 3

	nums1 := []int{2,3}
	m := 2
	nums2 := []int{0}
	n := 0

	// nums1 := []int{1}
	// m := 1
	// nums2 := []int{0}
	// n := 0

	// nums1 := []int{2,3,6,7,0,0}
	// m := 4
	// nums2 := []int{1,1}
	// n := 2

	// nums1 := []int{0}
	// m := 0
	// nums2 := []int{1}
	// n := 1

	fmt.Println(merge(nums1, m, nums2, n))
}
