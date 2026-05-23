// 1546. Maximum Number of Non-Overlapping Subarrays With Sum Equals Target
//
// https://leetcode.com/problems/maximum-number-of-non-overlapping-subarrays-with-sum-equals-target/description/
//
// Given an array nums, return the maximum number of non-empty subarrays you can pick such that each subarray has a sum equal to target.
// One subarray is considered non-overlapping with another if they do not share any common elements.
package main

import "fmt"

func main() {
	fmt.Println(maxNonOverlapping([]int{1, 1, 1, 1, 1, 1, 1, 1}, 2))
}

func maxNonOverlapping(nums []int, target int) int {
	pfx := make(map[int]int, len(nums))
	pfx[0] = 0
	s := 0
	lastIdx := -1
	count := 0
	for i := range nums {
		s += nums[i]

		if startIdx, ok := pfx[s-target]; ok && startIdx >= lastIdx {
			count++
			lastIdx = i
		}
		pfx[s] = i
	}

	return count
}
