package main

import (
	"fmt"
)

func getRow(rowIndex int) []int {
	arr := make([]int, rowIndex + 1)
	arr[0] = 1
    for idx := 1; idx < rowIndex + 1; idx ++ {
		// fmt.Println("idx: ", idx, "arr: ", arr)
		for mx := idx; mx > 0; mx -- {
			// fmt.Println("mx:", mx)
			arr[mx] = arr[mx - 1] + arr[mx]
		} 
	}
	return arr
}

func main() {
	// rowIndex := 3

	// rowIndex := 4

	rowIndex := 0

	fmt.Println(getRow(rowIndex))
}