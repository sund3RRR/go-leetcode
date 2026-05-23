// 206. Reverse Linked List
//
// https://leetcode.com/problems/reverse-linked-list/description/
//
// Given the head of a singly linked list, reverse the list, and return the reversed list.
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseLinkedList(head *ListNode) *ListNode {
	var prev *ListNode
	var curr = head

	for curr != nil {
		next := curr.Next

		curr.Next = prev

		prev = curr
		curr = next
	}

	return prev
}
