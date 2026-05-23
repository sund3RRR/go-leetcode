// 1984. Minimum Difference Between Highest and Lowest of K Scores
//
// https://leetcode.com/problems/minimum-difference-between-highest-and-lowest-of-k-scores/description/
//
// You are given a 0-indexed integer array nums, where nums[i] represents the score of the ith student.
// You are also given an integer k.
// Pick the scores of any k students from the array so that the difference between the highest and the lowest of the k scores is minimized.
// Return the minimum possible difference.
package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(minimumDifference([]int{9, 4, 1, 7}, 5))
}

func minimumDifference(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}

	slices.Sort(nums)

	if len(nums) <= k {
		return nums[len(nums)-1] - nums[0]
	}

	k--
	m := nums[len(nums)-1]
	for i := k; i < len(nums); i++ {
		m = min(m, nums[i]-nums[i-k])
	}

	return m
}
