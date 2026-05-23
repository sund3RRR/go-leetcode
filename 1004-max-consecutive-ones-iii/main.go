// 1004. Max Consecutive Ones III
//
// https://leetcode.com/problems/max-consecutive-ones-iii/description/
//
// Given a binary array nums and an integer k, return the maximum number of consecutive 1's in the array if you can flip at most k 0's.
package main

import "fmt"

// func MaxConsecutiveOnesIII(nums []int, k int) int {
// 	left := 0
// 	zeros := 0
// 	maxLen := 0

// 	for right := 0; right < len(nums); right++ {
// 		if nums[right] == 0 {
// 			zeros++
// 		}

// 		for zeros > k {
// 			if nums[left] == 0 {
// 				zeros--
// 			}
// 			left++
// 		}

// 		maxLen = max(maxLen, right-left+1)
// 	}

// 	return maxLen
// }

func main() {
	fmt.Println(MaxConsecutiveOnesIII([]int{0, 0, 1, 1, 0, 1, 1}, 1))
}

func MaxConsecutiveOnesIII(nums []int, k int) int {
	var left, result, zeroCount int
	for right := range nums {
		if nums[right] == 0 {
			zeroCount++
		}

		for zeroCount > k {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}

		result = max(result, right-left+1)

	}
	return result
}
