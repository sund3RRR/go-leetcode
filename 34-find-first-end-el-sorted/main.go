// 34. Find First and Last Position of Element in Sorted Array
//
// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/
//
// Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target number.
// If target is not found in the array, return [-1, -1].
// You must write an algorithm with O(log n) runtime complexity.
package main

import "fmt"

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 9}, 6))
}

func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	start := left

	if left >= len(nums) || nums[left] != target {
		return []int{-1, -1}
	}

	right = len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	end := left - 1

	return []int{start, end}
}
