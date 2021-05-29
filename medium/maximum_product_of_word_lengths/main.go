package main

import (
	"fmt"
)

func calMask(word string) int {
	mask := 0
	for _, c := range(word) {
		offset := c - 97
		// fmt.Println("1 << offset: ", 1 << offset)
		mask |= (1 << offset)
	}
	return mask
}

func maxProduct(words []string) int {
	var mask []int

	for _, word := range(words) {
		mask = append(mask, calMask(word))
	}
	
	max := 0
	i := 0
	for i < len(words) {
		j := i + 1
		for j < len(words) {
			if mask[i] & mask[j] == 0 {
				temp := len(words[i]) * len(words[j])
				if temp > max {
					max = temp
				}				
			}
			j ++
		}
		i ++
	}

	return max


	// globalMax := 0
	// hMap := make(map[string]map[rune]int)
	// for _, word := range(words[1:]) {
	// 	hMap[word] = make(map[rune]int)
	// 	for _, char := range(word) {
	// 		_, ok := hMap[word][char]; if !ok {
	// 			hMap[word][char] = 1
	// 		}
	// 	}
	// }
	// // fmt.Println(hMap)

	// for i, _ := range(words) {
	// 	t1 := words[i]
	// 	for _, word := range(words[i + 1:]) {			
	// 		for i, char := range(t1) {
	// 			_, ok := hMap[word][char]; if ok {
	// 				break
	// 			} else {
	// 				if i == len(t1) - 1 {
	// 					localMax := len(t1) * len(word)
	// 					if localMax > globalMax {
	// 						globalMax = localMax
	// 					}
	// 				}					
	// 			}
	// 		}
	// 	}
	// }
	// return globalMax


	// globalMax := 0
	// for i, _ := range(words) {
	// 	t1 := words[i]
	// 	for _, word := range(words[i + 1:]) {			
	// 		t2 := word
	// 		// fmt.Println(t1, t2)

	// 		hMap := make(map[rune]int)
	// 		for _, char := range(t2) {
	// 			_, ok := hMap[char]; if !ok {
	// 				hMap[char] = 1
	// 			}
	// 		}
			
	// 		for i, char := range(t1) {
	// 			_, ok := hMap[char]; if ok {
	// 				break
	// 			} else {
	// 				if i == len(t1) - 1 {
	// 					localMax := len(t1) * len(t2)
	// 					if localMax > globalMax {
	// 						globalMax = localMax
	// 					}
	// 				}					
	// 			}
	// 		}
	// 	}
	// }
	// return globalMax
}

func main() {
	words := []string{"abcw","baz","foo","bar","xtfn","abcdef"}

	// words := []string{"a","ab","abc","d","cd","bcd","abcd"}

	// words := []string{"a","aa","aaa","aaaa"}
	
	fmt.Println(maxProduct(words))


	// fmt.Println(calMask("abc")) // 7 -> 111
}
