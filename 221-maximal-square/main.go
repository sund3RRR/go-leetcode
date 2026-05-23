// 221. Maximal Square
//
// https://leetcode.com/problems/maximal-square/description/
//
// Given an m x n binary matrix filled with 0's and 1's, find the largest square containing only 1's and return its area.
package main

import "fmt"

func main() {
	m := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},

		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '2', '2'},
		{'1', '0', '1', '2', '3'},
		{'1', '0', '0', '1', '0'},
	}
	// m := [][]byte{
	// 	{'0', '0', '0', '1', '0', '1', '1', '1'},
	// 	{'0', '1', '1', '0', '0', '1', '0', '1'},
	// 	{'1', '0', '1', '1', '1', '1', '0', '1'},
	// 	{'0', '0', '0', '1', '0', '0', '0', '0'},
	// 	{'0', '0', '1', '0', '0', '0', '1', '0'},
	// 	{'1', '1', '1', '0', '0', '1', '1', '1'},
	// 	{'1', '0', '0', '1', '1', '0', '0', '1'},
	// 	{'0', '1', '0', '0', '1', '1', '0', '0'},
	// 	{'1', '0', '0', '1', '0', '0', '0', '0'},
	// }
	fmt.Println(maximalSquare(m))
}

func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	n := len(matrix[0])

	var result int

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == '0' {
				continue
			}

		sizeLoop:
			for size := 0; i+size < m && j+size < n; size++ {
				nj := j + size
				for ni := i; ni <= i+size; ni++ {
					if matrix[ni][nj] != '1' {
						break sizeLoop
					}
				}

				ni := i + size
				for nj := j; nj <= j+size; nj++ {
					if matrix[ni][nj] != '1' {
						break sizeLoop
					}
				}

				result = max(result, size+1)
			}
		}
	}

	return result * result
}
