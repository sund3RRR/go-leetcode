// 74. Search a 2D Matrix
//
// https://leetcode.com/problems/search-a-2d-matrix/description/
//
// You are given an m x n integer matrix matrix with the following two properties:
//   - Each row is sorted in non-decreasing order.
//   - The first integer of each row is greater than the last integer of the previous row.
//
// Given an integer target, return true if target is in matrix or false otherwise.
// You must write an algorithm with O(log(m * n)) runtime complexity.
package main

import "fmt"

func main() {
	m := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	fmt.Println(searchMatrix2(m, 30))
}

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	left, right := 0, m*n-1

	for left <= right {
		mid := left + (right-left)/2
		guess := matrix[mid/n][mid%n]
		if guess == target {
			return true
		} else if guess < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}

func searchMatrix2(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	l1, r1 := 0, m-1
	for l1 <= r1 {
		m1 := l1 + (r1-l1)/2
		if matrix[m1][0] == target {
			return true
		} else if matrix[m1][0] < target {
			l2, r2 := 0, n-1
			for l2 <= r2 {
				m2 := l2 + (r2-l2)/2
				if matrix[m1][m2] == target {
					return true
				} else if matrix[m1][m2] < target {
					l2 = m2 + 1
				} else {
					r2 = m2 - 1
				}
			}

			l1 = m1 + 1
		} else {
			r1 = m1 - 1
		}
	}
	return false
}
