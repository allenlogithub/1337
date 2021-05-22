package main

import (
	"fmt"
)

var hash = map[byte]int {
	48: 0,
	49: 1,
	50: 2,
	51: 3,
	52: 4,
	53: 5,
	54: 6,
	55: 7,
	56: 8,
	57: 9,
	65: 97, // A: a
	66: 98, // B: b
	67: 99, // C: c
	68: 100,
	69: 101,
	70: 102,
	71: 103,
	72: 104,
	73: 105,
	74: 106,
	75: 107,
	76: 108,
	77: 109,
	78: 110,
	79: 111,
	80: 112,
	81: 113,
	82: 114,
	83: 115,
	84: 116,
	85: 117,
	86: 118,
	87: 119,
	88: 120,
	89: 121,
	90: 122,
	97: 97, // a: a
	98: 98, // b: b
	99: 99, // c: c
	100: 100,
	101: 101,
	102: 102,
	103: 103,
	104: 104,
	105: 105,
	106: 106,
	107: 107,
	108: 108,
	109: 109,
	110: 110,
	111: 111,
	112: 112,
	113: 113,
	114: 114,
	115: 115,
	116: 116,
	117: 117,
	118: 118,
	119: 119,
	120: 120,
	121: 121,
	122: 122,
}


func isPalindrome(s string) bool {
    // i, j := 0, len(s) - 1
	// res := true
	// for i < len(s) {
	// 	if val, ok := hash[s[i]]; ok {
	// 		for j > 0 {
	// 			if val2, ok := hash[s[j]]; ok {
	// 				if val2 != val {
	// 					res = false
	// 				}
	// 				j --
	// 				break
	// 			} else {
	// 				j --
	// 			}
	// 		}
	// 	}
	// 	i ++		
	// }	
	// return res

	i := 0
	var arr []int
	for i < len(s) {
		if val, ok := hash[s[i]]; ok {
			arr = append(arr, val)
		}
		i ++
	}	
	i, arrLen, res := 0, len(arr), true
	halfArrLen := (arrLen - (len(arr) % 2)) / 2
	for i < halfArrLen {
		if arr[i] != arr[arrLen - 1 - i] {
			res = false
			break
		}
		i ++
	}
	return res
}

func main() {
	s := "A man, a plan, a canal: Panama"

	// s := "z"
	
	// s := "0P"

	// s := "0"

	// s := ":"

	// s := "race a car"

	fmt.Println(isPalindrome(s))
}
