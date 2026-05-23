// 238. Product of Array Except Self
//
// https://leetcode.com/problems/product-of-array-except-self/description/
//
// Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
// You must write an algorithm that runs in O(n) time and without using the division operation.
package main

import "fmt"

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}

func productExceptSelf(nums []int) []int {
	n := len(nums)
	pfx := make([]int, n)
	pfx[0] = 1
	sfx := make([]int, n)
	sfx[n-1] = 1

	for i := 1; i < n; i++ {
		pfx[i] = pfx[i-1] * nums[i-1]
	}
	for i := n - 2; i >= 0; i-- {
		sfx[i] = sfx[i+1] * nums[i+1]
	}

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = pfx[i] * sfx[i]
	}

	return ans
}
