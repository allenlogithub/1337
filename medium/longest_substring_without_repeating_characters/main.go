package main

import (
	"fmt"
)

func inSlice(element byte, list []byte) (bool, int) {
	for idx, v := range(list) {
		if v == element {
			return true, idx
		}
	}
	return false, 0
}

func lengthOfLongestSubstring(s string) int {
	var localSlice []byte
	var globalSlice []byte
	for i, _ := range(s) {
		// fmt.Println(s[i])
		isIn, idx := inSlice(s[i], localSlice); if isIn {			
			// localSlice = localSlice[idx + 1:]			
			localSlice = append(localSlice[idx + 1:], s[i])
		} else {
			localSlice = append(localSlice, s[i])
		}

		// fmt.Println(localSlice)
		if len(localSlice) > len(globalSlice) {
			globalSlice = localSlice
		}
	}
	return len(globalSlice)
}

func main() {
	// a := 3
	// arr := []int{2,1,3}
	// fmt.Println(inSlice(a, arr))

	s := "abcabcbb"

	// s := "bbbbb"

	// s := "pwwkew"

	// s := ""

	// s := " "

	// s := "aab"

	// s := "dvdf"

	// s := "abba"

	// s := "aabaab!bb"

	fmt.Println(lengthOfLongestSubstring(s))
}