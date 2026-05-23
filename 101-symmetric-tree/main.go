// 101. Symmetric Tree
//
// https://leetcode.com/problems/symmetric-tree/description/
//
// Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSymmetricRec(root.Left, root.Right)
}

func isSymmetricRec(a, b *TreeNode) bool {
	if a == nil || b == nil {
		return a == b
	}

	return a.Val == b.Val && isSymmetricRec(a.Left, b.Right) && isSymmetricRec(a.Right, b.Left)
}
