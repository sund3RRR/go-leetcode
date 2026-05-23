// 1493. Longest Subarray of 1's After Deleting One Element
//
// https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/description/
//
// Given a binary array nums, you should delete one element from it.
// Return the size of the longest non-empty subarray containing only 1's in the resulting array.
// Return 0 if there is no such subarray.
package main

import "fmt"

func main() {
	fmt.Println(longestSubarray([]int{1, 1, 0, 1}))
}

func longestSubarray(nums []int) int {
	var result int
	var left, right int

	var hasZero bool
	for i := range nums {
		if nums[i] == 1 {
			right++
		} else {
			hasZero = true
			result = max(result, left+right)
			left = right
			right = 0
		}
	}

	if !hasZero {
		return max(result, right-1)
	}

	return max(result, left+right)
}
