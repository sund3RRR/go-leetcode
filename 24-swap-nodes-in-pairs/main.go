// 24. Swap Nodes in Pairs
//
// https://leetcode.com/problems/swap-nodes-in-pairs/
//
// Given a linked list, swap every two adjacent nodes and return its head.
// You must solve the problem without modifying the values in the list's nodes (i.e., only nodes themselves may be changed.)
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummyHead := &ListNode{Next: head}
	prev := dummyHead
	curr := head

	for curr != nil && curr.Next != nil {
		curr2 := curr.Next
		next := curr2.Next

		prev.Next = curr2
		curr2.Next = curr
		curr.Next = next

		prev = curr
		curr = next
	}

	return dummyHead.Next
}
