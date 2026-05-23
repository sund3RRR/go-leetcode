// 268. Missing Number
//
// https://leetcode.com/problems/missing-number/description/
//
// Given an array nums containing n distinct numbers in the range [0, n], return the only number in the range that is missing from the array.
package main

func missingNumber(nums []int) int {
	n := len(nums)
	expected := n * (n - 1) / 2
	s := 0
	for _, v := range nums {
		s += v
	}
	return expected - s
}
