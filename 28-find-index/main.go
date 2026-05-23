// 28. Find the Index of the First Occurrence in a String
//
// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/description/
//
// Given two strings needle and haystack, return the index of the first occurrence
// of needle in haystack,or -1 if needle is not part of haystack.
package main

func strStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}

	result := -1
loop:
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		for j := range needle {
			if haystack[i+j] != needle[j] {
				continue loop
			}
		}
		return i
	}
	return result
}
