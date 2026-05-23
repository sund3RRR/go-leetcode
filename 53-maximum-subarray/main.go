// 53. Maximum Subarray
//
// https://leetcode.com/problems/maximum-subarray/
//
// Given an integer array nums, find the subarray with the largest sum, and return its sum.
package main

import "fmt"

func main() {
	fmt.Println(maxSubArray([]int{10, 2, -1, -2, 2, 1, -2, 1, 4, -5, 4}))
}
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	current := nums[0]
	best := nums[0]

	for i := 1; i < len(nums); i++ {
		current = max(nums[i], current+nums[i])
		best = max(best, current)
	}

	return best
}
