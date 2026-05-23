// 15. 3Sum
//
// https://leetcode.com/problems/3sum/description/
//
// Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that
// i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
//
// Notice that the solution set must not contain duplicate triplets.
package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(threeSum([]int{-100, -70, -60, 110, 120, 130, 160}))
}

func threeSum(nums []int) [][]int {
	slices.Sort(nums)

	var result [][]int

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum < 0 {
				left++
			} else if sum > 0 {
				right--
			} else {
				result = append(result, []int{nums[i], nums[left], nums[right]})

				left++
				right--

				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			}
		}
	}

	return result
}
