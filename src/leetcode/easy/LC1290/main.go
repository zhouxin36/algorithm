package main

import (
	"fmt"
	. "leetcode"
)

func main() {
	fmt.Println(getDecimalValue(ToListNode([]int{0})) == 0)
	fmt.Println(getDecimalValue(ToListNode([]int{1})) == 1)
	fmt.Println(getDecimalValue(ToListNode([]int{1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0})) == 18880)
	fmt.Println(getDecimalValue(ToListNode([]int{0, 0})) == 0)
}
func getDecimalValue(head *ListNode) int {
	var res int
	for head != nil {
		res <<= 1
		res |= head.Val
		head = head.Next
	}
	return res
}
