package main

import (
	"fmt"
)

func strStr(haystack string, needle string) int {
	num := -1
	needleLen := len(needle)
	haystackLen := len(haystack)

	if needleLen == 0 {
		num = 0
	} else if haystackLen == 0 {
		num = -1
	} else {
		for idx := 0; idx < haystackLen - needleLen + 1; idx ++ {
			if needle == haystack[idx: idx + needleLen] {
				num = idx
				break
			} else {
				if idx == needleLen - 1 {
					num = - 1
				}						
			}
		}
	}
	return num
}

func main() {
	// haystack := "hello"
	// needle := "ll"

	// haystack := "hello"
	// needle := "lo"

	// haystack := "hello"
	// needle := "he"

	// haystack := "aaaaa"
	// needle := "bba"

	// haystack := ""
	// needle := ""

	haystack := ""
	needle := "a"
	
	fmt.Println(strStr(haystack, needle))
}
