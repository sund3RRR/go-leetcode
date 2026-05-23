// 14. Longest Common Prefix
//
// https://leetcode.com/problems/longest-common-prefix/description/
//
// Write a function to find the longest common prefix string amongst an array of strings.
// If there is no common prefix, return an empty string "".
package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	for c := range prefix {
		for i := range strs {
			if len(strs[i]) == c || strs[i][c] != prefix[c] {
				return prefix[:c]
			}
		}
	}

	return prefix
}
