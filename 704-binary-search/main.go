// 704. Binary Search
//
// https://leetcode.com/problems/binary-search/description/
//
// Given an array of integers nums which is sorted in ascending order, and an integer target, write a function to search target in nums.
// If target exists, then return its index. Otherwise, return -1.
// You must write an algorithm with O(log n) runtime complexity.
package main

import "fmt"

func search(nums []int, target int) int {
	var left = 0
	var right = len(nums) - 1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}

func main() {
	fmt.Println(search([]int{0, 1, 2, 3, 4, 5, 6}, 10))
}
