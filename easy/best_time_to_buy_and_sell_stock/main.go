package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	sVal, diffVal := prices[0], 0
    for i := 1; i < len(prices); i ++ {
		if prices[i] > sVal {
			tempDiffVal := prices[i] - sVal
			if tempDiffVal > diffVal {
				diffVal = tempDiffVal
			}
		} else {
			sVal = prices[i]
		}
	}
	return diffVal
}

func main() {
	// prices := []int{7,1,5,3,6,4}

	// prices := []int{7,6,4,3,1}

	prices := []int{2,4,1}

	fmt.Println(maxProfit(prices))
}
