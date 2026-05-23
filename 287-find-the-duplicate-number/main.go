// 287. Find the Duplicate Number
//
// https://leetcode.com/problems/find-the-duplicate-number/description/
//
// Given an array of integers nums containing n + 1 integers where each integer is in the range [1, n] inclusive.
// There is only one repeated number in nums, return this repeated number.
// You must solve the problem without modifying the array nums and using only constant extra space.
package main

import "fmt"

func findDuplicateNumber(arr []int) int {
	tortoise := arr[0]
	hare := arr[0]

	for {
		tortoise = arr[tortoise]
		hare = arr[arr[hare]]
		if tortoise == hare {
			return tortoise
		}
	}
}

func main() {
	fmt.Println(findDuplicateNumber([]int{1, 3, 5, 4, 2, 6, 5}))
}
