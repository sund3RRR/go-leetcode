// 1277. Count Square Submatquares with All Ones
//
// https://leetcode.com/problems/count-square-submatrices-with-all-ones/description/
//
// Given a m * n binary matrix filled with 0's and 1's, returns the number of square submatquares with all ones.
package main

import "fmt"

func main() {
	m := [][]int{
		{0, 1, 1, 1},
		{1, 1, 1, 1},
		{0, 1, 1, 1},
	}
	fmt.Println(countSquares(m))
}

func countSquares(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := range m {
		for j := range n {
			if matrix[i][j] == 1 {
				a := val(dp, i-1, j)
				b := val(dp, i, j-1)
				c := val(dp, i-1, j-1)
				dp[i][j] = min(a, b, c) + 1
			}
		}
	}

	var result int
	for i := range m {
		for j := range n {
			result += dp[i][j]
		}
	}

	return result
}

func val(m [][]int, i, j int) int {
	if i < 0 || j < 0 {
		return 0
	}
	return m[i][j]
}
