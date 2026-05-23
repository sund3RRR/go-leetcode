// 35. Search Insert Position
//
// https://leetcode.com/problems/search-insert-position/
//
// Given a sorted array of distinct integers and a target value, return the index if the target is found.
// If not, return the index where it would be if it were inserted in order.
// You must write an algorithm with O(log n) runtime complexity.
package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 5, 5, 5, 7}, 5))
}

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}
