// 110. Balanced Binary Tree
//
// https://leetcode.com/problems/balanced-binary-tree/description/
//
// Given a binary tree, determine if it is height-balanced.
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := dfs(node.Left)
		if left == -1 {
			return -1
		}

		right := dfs(node.Right)
		if right == -1 {
			return -1
		}

		if abs(left-right) > 1 {
			return -1
		}

		return max(left, right) + 1
	}

	return dfs(root) != -1
}

func abs(num int) int {
	if num < 0 {
		return -num
	} else {
		return num
	}
}
