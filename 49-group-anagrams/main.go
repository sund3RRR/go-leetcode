// 49. Group Anagrams
//
// https://leetcode.com/problems/group-anagrams/description/
//
// Given an array of strings strs, group the anagrams together. You can return the answer in any order.
package main

import (
	"fmt"
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)

	for _, s := range strs {
		keyRune := []rune(s)
		slices.Sort(keyRune)
		key := string(keyRune)
		m[key] = append(m[key], s)
	}

	result := make([][]string, 0, len(m))

	for _, slice := range m {
		result = append(result, slice)
	}

	return result
}

func groupAnagramsFreqVector(strs []string) [][]string {
	m := make(map[[26]int][]string)

	for _, word := range strs {
		var count [26]int
		for _, ch := range word {
			count[ch-'a']++
		}
		m[count] = append(m[count], word)
	}

	res := make([][]string, 0, len(m))
	for _, group := range m {
		res = append(res, group)
	}
	return res
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
