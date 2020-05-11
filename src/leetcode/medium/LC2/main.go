package main

import (
	"fmt"
	. "leetcode"
)

func main() {
	fmt.Println(addTwoNumbers(ToListNode([]int{2, 4, 3}), ToListNode([]int{5, 6, 4})))
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	cur := res
	var carry int
	for l1 != nil || l2 != nil || carry != 0 {
		n := carry
		if l1 != nil {
			n += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n += l2.Val
			l2 = l2.Next
		}
		carry = n / 10
		n %= 10
		cur.Next = &ListNode{Val: n, Next: nil}
		cur = cur.Next
	}
	return res.Next
}
