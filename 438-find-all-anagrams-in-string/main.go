// 438. Find All Anagrams in a String
//
// https://leetcode.com/problems/find-all-anagrams-in-a-string/description/
//
// Given two strings s and p, return an array of all the start indices of p's anagrams in s.
// You may return the answer in any order.
// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.
package main

import "fmt"

func findAnagrams(s string, p string) []int {
	if len(p) > len(s) {
		return nil
	}

	need := make(map[byte]int)
	for i := range p {
		need[p[i]]++
	}

	diff := len(p)

	// инициализация первого окна
	for i := 0; i < len(p); i++ {
		if need[s[i]] > 0 {
			diff--
		}
		need[s[i]]--
	}

	result := make([]int, 0)

	if diff == 0 {
		result = append(result, 0)
	}

	// скользящее окно
	for i := len(p); i < len(s); i++ {
		// добавляем новый символ
		if need[s[i]] > 0 {
			diff--
		}
		need[s[i]]--

		// убираем старый символ
		left := s[i-len(p)]
		need[left]++
		if need[left] > 0 {
			diff++
		}

		if diff == 0 {
			result = append(result, i-len(p)+1)
		}
	}

	return result
}

func main() {
	fmt.Println(findAnagrams("aaabb", "bb"))
}
