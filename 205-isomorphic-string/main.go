// 205. Isomorphic Strings
//
// https://leetcode.com/problems/isomorphic-strings/description/
//
// Given two strings s and t, determine if they are isomorphic.
// Two strings s and t are isomorphic if the characters in s can be replaced to get t.
// All occurrences of a character must be replaced with another character while preserving the order of characters.
// No two characters may map to the same character, but a character may map to itself.
package main

import "fmt"

func main() {
	fmt.Println(isIsomorphic("foo", "bar"))
}

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m1, m2 := make(map[byte]byte, len(s)), make(map[byte]byte, len(s))
	for i := range s {
		if val, ok := m1[t[i]]; ok && val != s[i] {
			return false
		} else {
			m1[t[i]] = s[i]
		}

		if val, ok := m2[s[i]]; ok && val != t[i] {
			return false
		} else {
			m2[s[i]] = t[i]
		}

	}

	return true
}
