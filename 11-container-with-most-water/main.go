// 11. Container With Most Water
//
// https://leetcode.com/problems/container-with-most-water/description/
//
// You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
// Find two lines that together with the x-axis form a container, such that the container contains the most water.
// Return the maximum amount of water a container can store.
// Notice that you may not slant the container.

package main

func main() {

}

func maxArea(height []int) int {
	var maxSquare int
	left, right := 0, len(height)-1
	for left < right {
		s := (right - left) * min(height[left], height[right])
		maxSquare = max(maxSquare, s)

		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}

	return maxSquare
}
