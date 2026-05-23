package main

import (
	"fmt"
	ll "main/linked-list"
)

func rotateRight(head *ll.ListNode, k int) *ll.ListNode {
	if head.Next == nil {
		return head
	}

	var current = head
	var target = current
	var skip int
	for {
		if current.Next == nil {
			skip++

			if skip < 2 {
			}
			result := target.Next
			target.Next = nil
			current.Next = head
			return result
		}

		if skip < k {
			skip++
			current = current.Next
			continue
		}
		current = current.Next
		target = target.Next
	}
}

func main() {
	linkedList := ll.NewLinkedList([]int{1, 2, 3, 4, 5})
	fmt.Println(ll.LinkedListToArr(rotateRight(linkedList, 20)))
}
