// 18. 4Sum
//
// https://leetcode.com/problems/4sum/description/
//
// Given an array nums of n integers, return an array of all the unique quadruplets [nums[a], nums[b], nums[c], nums[d]] such that:
//   - 0 <= a, b, c, d < n
//   - a, b, c, and d are distinct.
//   - nums[a] + nums[b] + nums[c] + nums[d] == target
//
// Your solution must not contain duplicate quadruplets.
package main

import (
	"fmt"
	"slices"
)

func fourSum(nums []int, target int) [][]int {
	slices.Sort(nums)
	result := [][]int{}

	n := len(nums)

	for i := 0; i < n-3; i++ {
		// скипаем дубликаты для i
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < n-2; j++ {
			// скипаем дубликаты для j
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			left, right := j+1, n-1

			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]

				if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})

					// пропускаем дубликаты для left
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					// пропускаем дубликаты для right
					for left < right && nums[right] == nums[right-1] {
						right--
					}

					left++
					right--
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}

	return result
}

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
}
