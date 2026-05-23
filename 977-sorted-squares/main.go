// 977. Squares of a Sorted Array
//
// https://leetcode.com/problems/squares-of-a-sorted-array/description/
//
// Given an integer array nums sorted in non-decreasing order, return an array of the squares of each number sorted in non-decreasing order.
package main

import "fmt"

func sortedSquares(nums []int) []int {
	n := len(nums)
	res := make([]int, n)

	left, right := 0, n-1
	pos := n - 1

	for left <= right {
		l, r := nums[left]*nums[left], nums[right]*nums[right]
		if l > r {
			res[pos] = l
			left++
		} else {
			res[pos] = r
			right--
		}
		pos--
	}

	return res
}

func main() {
	fmt.Println(sortedSquares([]int{-3, -2, 0, 1, 2, 3, 4}))
}
