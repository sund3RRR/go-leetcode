// 19. Remove Nth Node From End of List
//
// https://leetcode.com/problems/remove-nth-node-from-end-of-list/
//
// Given the head of a linked list, remove the nth node from the end of the list and return its head.
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}

	dummyHead := &ListNode{Next: head}
	slow, fast := dummyHead, head
	var counter int
	for fast != nil {
		if counter < n {
			counter++
		} else {
			slow = slow.Next
		}

		fast = fast.Next
	}

	slow.Next = slow.Next.Next

	return dummyHead.Next
}
