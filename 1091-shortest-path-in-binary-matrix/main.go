// 1091. Shortest Path in Binary Matrix
//
// https://leetcode.com/problems/shortest-path-in-binary-matrix
//
// Given an n x n binary matrix grid, return the length of the shortest clear path in the matrix.
// If there is no clear path, return -1.
// A clear path in a binary matrix is a path from the top-left cell (i.e., (0, 0)) to the
// bottom-right cell (i.e., (n - 1, n - 1)) such that:
//   - All the visited cells of the path are 0.
//   - All the adjacent cells of the path are 8-directionally connected (i.e., they are different and they share an edge or a corner).
//   - The length of a clear path is the number of visited cells of this path.
package main

import "fmt"

func main() {
	// grid := [][]int{
	// 	{0, 0, 0},
	// 	{1, 1, 0},
	// 	{1, 1, 0},
	// }
	grid := [][]int{
		{0, 0, 0, 0},
		{1, 0, 1, 0},
		{1, 1, 0, 0},
		{1, 1, 1, 0},
	}

	fmt.Println(shortestPathBinaryMatrix(grid))
}

type Item struct {
	coords [2]int
	depth  int
}

func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	q := make([]Item, 0)

	if grid[0][0] == 1 {
		return -1
	}

	q = append(q, Item{coords: [2]int{0, 0}})

	for len(q) != 0 {
		item := q[0]
		q = q[1:]

		x := item.coords[0]
		y := item.coords[1]
		depth := item.depth + 1

		if item.coords == [2]int{n - 1, n - 1} {
			return depth
		}

		// 8 cases
		q = appendQueue(q, x-1, y-1, depth, grid)
		q = appendQueue(q, x, y-1, depth, grid)
		q = appendQueue(q, x-1, y, depth, grid)
		q = appendQueue(q, x+1, y-1, depth, grid)
		q = appendQueue(q, x+1, y, depth, grid)
		q = appendQueue(q, x+1, y+1, depth, grid)
		q = appendQueue(q, x, y+1, depth, grid)
		q = appendQueue(q, x-1, y+1, depth, grid)

		grid[x][y] = 1
	}

	return -1
}

func appendQueue(q []Item, x, y int, depth int, grid [][]int) []Item {
	if isValid(x, y, len(grid)) && grid[x][y] == 0 {
		grid[x][y] = 1
		return append(q, Item{coords: [2]int{x, y}, depth: depth})
	}
	return q
}

func isValid(x, y int, n int) bool {
	if x < 0 || y < 0 || x == n || y == n {
		return false
	}

	return true
}
