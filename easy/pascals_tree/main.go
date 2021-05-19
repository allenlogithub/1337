package main

import (
	"fmt"
)

func generate(numRows int) [][]int {
	res := [][]int{}
	if numRows == 1 {
		res = append(res, []int{1})
	} else if numRows == 2 {
		res = append(res, []int{1})
		res = append(res, []int{1,1})
	} else {
		res = append(res, []int{1})
		res = append(res, []int{1,1})
		downArr := []int{1,1}
		for idx := 1; idx < numRows - 1; idx ++ {	
			var temp []int
			for pos := 0; pos < len(downArr) + 1; pos ++ {
				if pos == 0 || pos == len(downArr) {
					temp = append(temp, 1)
				} else {
					temp = append(temp, downArr[pos - 1] + downArr[pos])
				}
			}
			downArr = temp
			res = append(res, temp)
		}
	}
	return res
}

func main() {
	numRows := 5

	// numRows := 1

	fmt.Println(generate(numRows))
}

// 1
// 1,1
// 1,2,1
// 1,3,3,1
// 1,4,6,4,1
// 1,5,10,10,5,1

// 1,1,1,1,1,1
// 1,2,3,4,5
// 1,3,6,10
// 1,4,10
// 1,5
// 1