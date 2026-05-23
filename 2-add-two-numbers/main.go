package main

import (
	"fmt"
	ll "main/linked-list"
)

func addTwoNumbers(l1 *ll.ListNode, l2 *ll.ListNode) *ll.ListNode {

}

func main() {
	l1 := ll.NewLinkedList([]int{2, 4, 3})
	l2 := ll.NewLinkedList([]int{5, 6, 4})

	l := addTwoNumbers(l1, l2)
	fmt.Println(ll.LinkedListToArr(l))
}
