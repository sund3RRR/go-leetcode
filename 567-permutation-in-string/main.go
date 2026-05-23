// 567. Permutation in String
//
// https://leetcode.com/problems/permutation-in-string/description/
//
// Given two strings s1 and s2, return true if s2 contains a permutation of s1, or false otherwise.
// In other words, return true if one of s1's permutations is the substring of s2.
package main

import "fmt"

func main() {
	fmt.Println(checkInclusion("ooo", "eidbaooo"))
}

func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}

	need := make(map[byte]int, len(s1))
	for i := range s1 {
		need[s1[i]]++
	}

	diff := len(s1)
	for i := range s1 {
		if need[s2[i]] > 0 {
			diff--
		}
		need[s2[i]]--
	}

	if diff == 0 {
		return true
	}

	for i := len(s1); i < len(s2); i++ {
		if need[s2[i]] > 0 {
			diff--
		}
		need[s2[i]]--

		need[s2[i-len(s1)]]++
		if need[s2[i-len(s1)]] > 0 {
			diff++
		}

		if diff == 0 {
			return true
		}
	}

	return false
}
