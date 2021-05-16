package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	strCount := len(strs)

	shortest_str := strs[0]
	for _, str := range(strs) {
		if len(str) < len(shortest_str) {
			shortest_str = str
		}
	} 

	end := 0
	stop := false
	for idx, char := range(shortest_str) {
		count := 0
		for _, str := range(strs) {
			if string(char) == str[idx: idx + 1] {
				count = count + 1				
			} else {
				stop = true
				break
			}

			if count == strCount {
				end = end + 1
			}
		}
		if stop == true {
			break
		}
	}

	return shortest_str[:end]
}

func main() {
	// arr := []string{"flower","flow","flight"}

	// arr := []string{"flower","flow","fight"}

	arr := []string{"ab", "a"}

	// arr := []string{"flower"}

	// arr := []string{""}

	// arr := []string{"dog","racecar","car"}

	fmt.Println(longestCommonPrefix(arr))
}
