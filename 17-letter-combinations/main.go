// 17. Letter Combinations of a Phone Number
//
// https://leetcode.com/problems/letter-combinations-of-a-phone-number/
//
// Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.
// Return the answer in any order.
// A mapping of digits to letters (just like on the telephone buttons) is given below.
// Note that 1 does not map to any letters.
package main

import (
	"fmt"
)

func main() {
	fmt.Println(letterCombinations("8259"))
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	} else if len(digits) == 1 {
		return digitToLetters(digits[0])
	}

	lower := letterCombinations(digits[1:])
	letters := digitToLetters(digits[0])

	result := make([]string, 0)
	for i := range letters {
		for j := range lower {
			result = append(result, letters[i]+lower[j])
		}
	}

	return result
}

func digitToLetters(digit byte) []string {
	switch digit {
	case '2':
		return []string{"a", "b", "c"}
	case '3':
		return []string{"d", "e", "f"}
	case '4':
		return []string{"g", "h", "i"}
	case '5':
		return []string{"j", "k", "l"}
	case '6':
		return []string{"m", "n", "o"}
	case '7':
		return []string{"p", "q", "r", "s"}
	case '8':
		return []string{"t", "u", "v"}
	case '9':
		return []string{"w", "x", "y", "z"}
	default:
		return []string{}
	}
}
