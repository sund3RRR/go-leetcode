// 234. Palindrome Linked List
//
// https://leetcode.com/problems/palindrome-linked-list/description/
//
// Given the head of a singly linked list, return true if it is a palindrome or false otherwise.
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slow, fast := head, head

	var prev *ListNode

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next

		next := slow.Next
		slow.Next = prev
		prev = slow
		slow = next
	}

	if fast != nil {
		slow = slow.Next
	}

	for prev != nil && slow != nil {
		if prev.Val != slow.Val {
			return false
		}
		prev = prev.Next
		slow = slow.Next
	}

	return true
}

func isPalindrome2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	stack := make([]int, 0)

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		stack = append(stack, slow.Val)
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast != nil {
		slow = slow.Next
	}

	for len(stack) != 0 {
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if item != slow.Val {
			return false
		}

		slow = slow.Next
	}

	return true
}

func main() {
	l := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
		},
	}
	fmt.Println(isPalindrome2(l))
}
