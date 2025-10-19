// 48. Rotate Image
//
// Difficulty: Medium
// Link: 	   https://leetcode.com/problems/rotate-image
//
// You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).
// You have to rotate the image in-place, which means you have to modify the input 2D matrix directly.
// DO NOT allocate another 2D matrix and do the rotation.
package rotate_image

func Rotate(matrix [][]int) {
	rotate(matrix)
}

func rotate(matrix [][]int) {
	for i := range matrix {
		for j := i; j < len(matrix[i]); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := range matrix {
		for j, k := 0, len(matrix[i])-1; j < k; j, k = j+1, k-1 {
			matrix[i][j], matrix[i][k] = matrix[i][k], matrix[i][j]
		}
	}
}
