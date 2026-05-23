// 763. Partition Labels
//
// https://leetcode.com/problems/partition-labels/description/
//
// You are given a string s. We want to partition the string into as many parts as possible so that each letter appears in at most one part.
// Note that the partition is done so that after concatenating all the parts in order, the resultant string should be s.
// Return a list of integers representing the size of these parts.
package main

import "fmt"

const ALPHABET_SIZE = 26

func partitionLabels(s string) []int {
	m := make(map[rune]int, ALPHABET_SIZE)

	for i, r := range s {
		m[r] = i
	}

	pos := make([]int, 0, len(s))
	for _, r := range s {
		pos = append(pos, m[r])
	}

	result := make([]int, 0)
	l := -1
	localMax := 0
	for i, p := range pos {
		localMax = max(localMax, p)

		if localMax == p && p == i {
			result = append(result, i-l)
			l = i
		}
	}

	return result
}

func main() {
	fmt.Println(partitionLabels("eccbbbbdec"))
}
