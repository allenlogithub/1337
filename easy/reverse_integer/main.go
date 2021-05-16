package main

import (
	"fmt"
)

func reverse(x int) int {	
	var arr []int
	status := 1
	if x < 0 {
		status = -1
		x = x * status
	}

	if x < 10 {
		return x
	}

	for x >= 10 {
		digit := x % 10
		x = (x - digit) / 10
		
		arr = append(arr, digit)
		if x < 10 {
			arr = append(arr, x)
		}
	}
		
	arrLen := len(arr)
	res := 0
	for idx, val := range(arr) {
		res = res + val * power(10, (arrLen - idx -1))
	}
	if (res < power(-2, 31)) {
		return 0
	} else if (res > (power(2, 31) -1)) {
		return 0
	} else {
		return res * status
	}
}

func power(x int, n int) int {
	res := 1
	for n > 0 {
		res = res * x
		n = n - 1
	}
	return res
}

func main() {
	// fmt.Println(power(10, 3))
	// fmt.Println(power(2, 4))
	// fmt.Println(power(-3, 3))

	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(120))
	fmt.Println(reverse(0))
	fmt.Println(reverse(1))
	fmt.Println(reverse(10))
	fmt.Println(reverse(1534236469))
}
