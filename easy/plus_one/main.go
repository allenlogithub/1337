package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	digitsLen, carry := len(digits), false
	var arr []int

	for idx := 0; idx < digitsLen; idx ++ {
		revIdx := digitsLen - idx - 1

		if idx == 0 {
			if digits[revIdx] == 9 {
				carry = true
				arr = append(arr, 0)
			} else {
				arr = append(arr, digits[revIdx] + 1)
			}
		} else if idx > 0 {
			if carry == true {
				if digits[revIdx] == 9 {
					arr = append(arr, 0)		
				} else {
					arr = append(arr, digits[revIdx] + 1)
					carry = false
				}
			} else {
				arr = append(arr, digits[revIdx])
			}
		}		
	}
	if carry == true {
		arr = append(arr, 1)
	}
	for i, j := 0, len(arr) - 1; i < j; i, j = i + 1, j - 1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func main() {
	// digits := []int{1,2,3}

	// digits := []int{1,2,9}

	// digits := []int{1,2,9,3,4,5,5,2,1,4}

	digits := []int{7,2,8,5,0,9,1,2,9,5,3,6,6,7,3,2,8,4,3,7,9,5,7,7,4,7,4,9,4,7,0,1,1,1,7,4,0,0,6}

	// digits := []int{7,2,8,5,0,9,1,2,9,5,3,6,6,7,3,2,8,4}

	// digits := []int{9}

	// digits := []int{9,9}

	// digits := []int{0}

	fmt.Println(plusOne(digits))
}
