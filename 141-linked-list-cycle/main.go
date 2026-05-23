// 141. Linked List Cycle
//
// https://leetcode.com/problems/linked-list-cycle/description/
//
// Given head, the head of a linked list, determine if the linked list has a cycle in it.
// There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the Next pointer.
// Internally, pos is used to denote the index of the node that tail's Next pointer is connected to.
// Note that pos is not passed as a parameter.
// Return true if there is a cycle in the linked list. Otherwise, return false.
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	var bunny, tortoise = head, head
	for bunny != nil && bunny.Next != nil {
		bunny = bunny.Next.Next
		tortoise = tortoise.Next

		if bunny == tortoise && bunny != head {
			return true
		}
	}

	return true
}

func main() {
	sec := &ListNode{Val: 2}
	head := &ListNode{Val: 1, Next: sec}

	hg := hasCycle(head)

	fmt.Println(hg)
}
