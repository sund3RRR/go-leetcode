// 23. Merge k Sorted Lists
//
// https://leetcode.com/problems/merge-k-sorted-lists/description/
//
// You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.
// Merge all the linked-lists into one sorted linked-list and return it.
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists ...*ListNode) *ListNode {
	var result *ListNode

	for _, list := range lists {
		result = mergeTwoLists(result, list)
	}

	return result
}

func mergeTwoLists(a, b *ListNode) *ListNode {
	if a == nil {
		return b
	}

	if b == nil {
		return a
	}

	var head = &ListNode{}
	var merged = head
	for a != nil && b != nil {
		if a.Val < b.Val {
			merged.Next = a
			a = a.Next
		} else {
			merged.Next = b
			b = b.Next
		}
		merged = merged.Next
	}

	if a == nil {
		merged.Next = b
	}

	if b == nil {
		merged.Next = a
	}

	return head.Next
}
