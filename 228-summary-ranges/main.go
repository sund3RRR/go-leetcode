// 228. Summary Ranges
//
// https://leetcode.com/problems/summary-ranges/description/
//
// You are given a sorted unique integer array nums.
// A range [a,b] is the set of all integers from a to b (inclusive).
// Return the smallest sorted list of ranges that cover all the numbers in the array exactly.
// That is, each element of nums is covered by exactly one of the ranges,
// and there is no integer x such that x is in one of the ranges but not in nums.
//
// Each range [a,b] in the list should be output as:
//   - "a->b" if a != b
//   - "a" if a == b
package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	var startIdx int
	var result = make([]string, 0)

	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] == 1 {
			continue
		}
		result = append(result, formatRange(nums[startIdx], nums[i-1]))
		startIdx = i
	}

	result = append(result, formatRange(nums[startIdx], nums[len(nums)-1]))

	return result
}

func formatRange(start, end int) string {
	if start == end {
		return strconv.Itoa(start)
	}

	return strconv.Itoa(start) + "->" + strconv.Itoa(end)
}

func main() {
	fmt.Println(summaryRanges([]int{}))
}
