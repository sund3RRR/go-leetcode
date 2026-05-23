// 153. Find Minimum in Rotated Sorted Array
//
// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/description/
//
// Suppose an array of length n sorted in ascending order is rotated between 1 and n times.
// Given the sorted rotated array nums of unique elements, return the minimum element of this array.
// You must write an algorithm that runs in O(log n) time.
package main

import "fmt"

func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return nums[left]
}

func main() {
	fmt.Println(findMin2([]int{1, 2, 3, 4, 5}))
}

func findMin2(nums []int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (right + left) / 2

		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return nums[left]
}
