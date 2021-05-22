package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	i, sum := 1, 0
	for i < len(prices) {
		if prices[i] > prices[i - 1] {
			sum = sum + (prices[i] - prices[i - 1])
		}
		i ++
	}
	return sum
	
	// i, buy, sell, sum, N := 0, 0, 0, 0, len(prices) - 1
	// if len(prices) == 1 {
	// 	sum = 0
	// }
	// for i < N {
	// 	fmt.Println(i)
	// 	for i < N && prices[i + 1] <= prices[i] {
	// 		buy = i + 1
	// 		i ++			
	// 	} 
	// 	for i < N && prices[i + 1] > prices[i] {
	// 		sell = i + 1
	// 		i ++					
	// 	}
	// 	sum = sum + (prices[sell] - prices[buy])
	// 	fmt.Println(sell, buy, sum)
	// }	
}

func main() {
	prices := []int{7,1,5,3,6,4}

	// prices := []int{6,1,3,2,4,7}

	// prices := []int{1}

	// prices := []int{1,2,3,4,5}

	// prices := []int{7,6,4,3,1}

	fmt.Println(maxProfit(prices))
}