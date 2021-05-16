package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
    if x < 0 {
		return false
	} else if x < 10 {
		return true
	}

	var arr []int
	for x >= 10 {
		digit := x % 10
		x = (x - digit) / 10

		arr = append(arr, digit)
		if x < 10 {
			arr = append(arr , x)
		}
	}

	arrLen := len(arr)
	halfArrLen := 0
	if arrLen % 2 == 1 {
		halfArrLen = (arrLen - 1) / 2
	} else {
		halfArrLen = arrLen / 2
	}

	var res bool
	for idx := 0; idx < halfArrLen; idx++ {
		if arr[idx] != arr[arrLen - idx - 1] {
			res = false
			return res
		} else {
			res = true
		}		
	}	
	return res
}

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
	fmt.Println(isPalindrome(-101))
	fmt.Println(isPalindrome(42312))
	fmt.Println(isPalindrome(43234))
	fmt.Println(isPalindrome(1000021))
}
