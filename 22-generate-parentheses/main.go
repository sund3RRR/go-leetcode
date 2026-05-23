// 22. Generate Parentheses
//
// https://leetcode.com/problems/generate-parentheses/description/
//
// Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
package main

import "fmt"

func generateParenthesis(n int) []string {
	var result []string

	var backtrack func(current string, opened, closed int)

	backtrack = func(current string, opened, closed int) {
		if len(current) == 2*n {
			result = append(result, current)
			return
		}

		if opened < n {
			backtrack(current+"(", opened+1, closed)
		}
		if closed < opened {
			backtrack(current+")", opened, closed+1)
		}
	}

	backtrack("", 0, 0)
	return result
}

func main() {
	fmt.Println(generateParenthesis(4))
}
