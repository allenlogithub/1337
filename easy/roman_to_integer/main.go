package main

import (
	"fmt"
)

func romanToInt(s string) int {
    rMap := make(map[string]int)

	rMap ["I"] = 1
	rMap ["V"] = 5
	rMap ["X"] = 10
	rMap ["L"] = 50
	rMap ["C"] = 100
	rMap ["D"] = 500
	rMap ["M"] = 1000

	
	rMap ["IV"] = 4
	rMap ["IX"] = 9
	rMap ["XL"] = 40
	rMap ["XC"] = 90
	rMap ["CD"] = 400
	rMap ["CM"] = 900

	idx := 0
	sum := 0
	for idx < len(s) {
		if idx < len(s) -1 {
			i, ok := rMap[s[idx: idx + 2]]
			if ok {
				sum = sum + i
				idx = idx + 2
			} else {
				i, ok := rMap[s[idx: idx + 1]]
				if ok {
					sum = sum + i
					idx = idx + 1
				}			
			}
		} else {
			i, ok := rMap[s[idx: idx + 1]]
			if ok {
				sum = sum + i
				idx = idx + 1
			}
		}
				
	}
	return sum
}

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("IV"))
	fmt.Println(romanToInt("IX"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}