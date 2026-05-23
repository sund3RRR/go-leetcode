// 300. Longest Increasing Subsequence
//
// https://leetcode.com/problems/longest-increasing-subsequence/description/
//
// Given an integer array nums, return the length of the longest strictly increasing subsequence.
package main

import "fmt"

func lengthOfLIS(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	count := 1
	prev := count

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			count = max(prev, count+1)
			prev = count
		}
	}

	return count
}

func main() {
	fmt.Println(lengthOfLIS([]int{1, 1}))
}
