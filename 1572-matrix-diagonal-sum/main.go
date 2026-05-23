// 1572. Matrix Diagonal Sum
//
// https://leetcode.com/problems/matrix-diagonal-sum/description/
//
// Given a square matrix mat, return the sum of the matrix diagonals.
// Only include the sum of all the elements on the primary diagonal and all the elements
// on the secondary diagonal that are not part of the primary diagonal.
package main

func diagonalSum(mat [][]int) int {
	var s int
	for i, j1, j2 := 0, 0, len(mat)-1; i < len(mat); i, j1, j2 = i+1, j1+1, j2-1 {
		if j1 == j2 {
			s += mat[i][j1]
			continue
		}
		s += mat[i][j1] + mat[i][j2]
	}
	return s
}
