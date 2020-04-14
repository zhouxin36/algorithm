package leetcode

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
