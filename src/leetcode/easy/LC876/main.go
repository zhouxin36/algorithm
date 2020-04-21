package main

import (
	"fmt"
	. "leetcode"
)

func main() {
	fmt.Println(middleNode(ToListNode([]int{1, 2, 3, 4, 5})).Val == 3)
	fmt.Println(middleNode(ToListNode([]int{1, 2, 3, 4, 5, 6})).Val == 4)
}
func middleNode(head *ListNode) *ListNode {
	var slow, fast = head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
