// 36. Valid Sudoku
//
// https://leetcode.com/problems/valid-sudoku/description/
//
// Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:
//   - Each row must contain the digits 1-9 without repetition.
//   - Each column must contain the digits 1-9 without repetition.
//   - Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
//
// Note:
// A Sudoku board (partially filled) could be valid but is not necessarily solvable.
// Only the filled cells need to be validated according to the mentioned rules.
package main

import "fmt"

func main() {
	sudoku := [][]byte{
		{'.', '8', '7', '6', '5', '4', '3', '2', '1'},
		{'2', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'3', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'4', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'5', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'6', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'7', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'8', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'9', '.', '.', '.', '.', '.', '.', '.', '.'},
	}
	fmt.Println(isValidSudoku(sudoku))
}

func isValidSudoku(board [][]byte) bool {
	rows, cols := [9][10]int{}, [9][10]int{}
	cells := [3][3][10]int{}

	for i := range board {
		for j := range board[i] {
			if board[i][j] == '.' {
				continue
			}

			num := int(board[i][j] - 0x30)

			if rows[i][num] == 1 {
				return false
			}
			rows[i][num] += 1

			if cols[j][num] == 1 {
				return false
			}
			cols[j][num] += 1

			k, l := i/3, j/3
			if cells[k][l][num] == 1 {
				return false
			}
			cells[k][l][num] += 1
		}
	}
	return true
}
