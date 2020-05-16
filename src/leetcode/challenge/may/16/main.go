package main

import (
	"fmt"
	. "leetcode"
)

func main() {
	fmt.Println(oddEvenList(ToListNode([]int{1, 2, 3, 4, 5})))
	fmt.Println(oddEvenList(ToListNode([]int{2, 1, 3, 5, 6, 4, 7})))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	odd := head
	even := head.Next
	evenHead := head.Next
	if even == nil || even.Next == nil {
		return head
	}
	node := even.Next
	for node != nil {
		odd.Next = node
		odd = odd.Next
		node = node.Next
		if node != nil {
			even.Next = node
			even = even.Next
			node = node.Next
		}
	}
	odd.Next = evenHead
	even.Next = nil
	return head
}
