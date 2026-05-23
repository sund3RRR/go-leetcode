// 283. Move Zeroes
//
// https://leetcode.com/problems/move-zeroes/description/
//
// Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
// Note that you must do this in-place without making a copy of the array.
package main

func moveZeroes(nums []int) {
	buff := make([]int, 0)

	for i := range nums {
		if len(buff) > 0 && nums[i] != 0 {
			idx := buff[0]
			nums[idx], nums[i] = nums[i], nums[idx]
			buff = buff[1:]
		}
		if nums[i] == 0 {
			buff = append(buff, i)
		}
	}
}

func main() {
	moveZeroes([]int{0, 1, 0, 3, 0, 0, 12, 0, 0})
}
