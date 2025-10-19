// 1. Two Sum
//
// Difficulty: Easy
// Link:	   https://leetcode.com/problems/two-sum
//
// Given an array of integers 'nums' and an integer 'target', return indices of the two numbers such that they add up to 'target'.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.
package two_sum

func TwoSum(nums []int, target int) []int {
	return twoSum(nums, target)
}

func twoSum(nums []int, target int) []int {
	var m = make(map[int]int, len(nums))
	for i, v := range nums {
		reverted := target - v
		if index, ok := m[reverted]; ok {
			return []int{index, i}
		}
		m[v] = i
	}
	return []int{}
}
