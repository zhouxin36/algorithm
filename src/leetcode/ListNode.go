package leetcode

import "strings"

type ListNode struct {
	Val  int
	Next *ListNode
}

func ToListNode(arr []int) *ListNode {
	root := ListNode{Val: arr[0]}
	next := &root
	for i := 1; i < len(arr); i++ {
		next.Next = &ListNode{Val: arr[i]}
		next = next.Next
	}
	return &root
}
func (node *ListNode) String() string {
	var res strings.Builder
	head := node
	for head != nil {
		res.WriteByte(byte('0' + head.Val))
		head = head.Next
	}
	return res.String()
}
