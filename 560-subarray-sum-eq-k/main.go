// 560. Subarray Sum Equals K
//
// https://leetcode.com/problems/subarray-sum-equals-k/description/
//
// Given an array of integers nums and an integer k, return the total number of subarrays whose sum equals to k.
// A subarray is a contiguous non-empty sequence of elements within an array.
package main

import "fmt"

func subarraySum(nums []int, k int) int {
	pfxSum := make(map[int]int, len(nums))
	s := 0
	count := 0
	pfxSum[0] = 1
	for _, num := range nums {
		s += num
		if pfxCount, ok := pfxSum[s-k]; ok {
			count += pfxCount
		}
		pfxSum[s]++
	}
	return count
}

func main() {
	fmt.Println(subarraySum2([]int{1, 2, 3}, 3))
}

func subarraySum2(nums []int, k int) int {
	pfxSum := make(map[int]int, len(nums))
	var s, count int
	pfxSum[0] = 1
	for i := range nums {
		s += nums[i]
		if c, ok := pfxSum[s-k]; ok {
			count += c
		}
		pfxSum[s]++
	}

	fmt.Println(pfxSum)
	return count
}
