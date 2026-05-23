// 200. Number of Islands
//
// https://leetcode.com/problems/number-of-islands/description/
//
// Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water), return the number of islands.
// An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically.
// You may assume all four edges of the grid are all surrounded by water.
package main

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m := len(grid)
	n := len(grid[0])

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || j < 0 || i > n || j > m || grid[i][j] == 0 {
			return
		}

		grid[i][j] = 0
		dfs(i, j-1)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i+1, j)
	}

	var count int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				count++
				dfs(i, j)
			}
		}
	}

	return count
}

func numIslands2(grid [][]byte) int {
	n, m := len(grid), len(grid[0])

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || j < 0 || i >= n || j >= m || grid[i][j] == 0 {
			return
		}

		grid[i][j] = 0

		dfs(i-1, j)
		dfs(i, j-1)
		dfs(i+1, j)
		dfs(i, j+1)
	}

	var result int
	for i := range n {
		for j := range m {
			if grid[i][j] == 1 {
				result++
				dfs(i, j)
			}
		}
	}

	return result
}
