package main

import (
    "fmt"
)

func productExceptSelf(nums []int) []int {
	// var res []int
	// val := 1
    // for _, num := range(nums) {
    //     res = append(res, val)		
	// 	val *= num
    // }

	// i, val := len(nums) - 1, 1
	// for i > -1 {		
	// 	res[i] = res[i] * val
	// 	val *= nums[i]
	// 	i --
	// }

	// return res

	var fArr []int
	val := 1
	for _, num := range(nums) {
		fArr = append(fArr, val)
		val *= num
	}
	// fmt.Println("fArr: ", fArr)

	var bArr []int
	val = 1
	i := len(nums) - 1
	for i > -1 {
		bArr = append(bArr, val)
		val *= nums[i]
		i --
	}
	// fmt.Println("bArr: ", bArr)

	i = 0
	var res []int
	for i < len(nums) {
		res = append(res, fArr[i] * bArr[len(nums) - 1 - i])
		i ++
	}
	// fmt.Println("res: ", res)
    return res
}

func main() {
	nums := []int{1,2,3,4}

	fmt.Println(productExceptSelf(nums))
}
