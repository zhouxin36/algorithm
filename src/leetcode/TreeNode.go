package leetcode

import (
	"math"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type LinkNode struct {
	Val  *TreeNode
	Next *LinkNode
}

func ToTreeNode(nums []int) *TreeNode {
	var head, tag *LinkNode
	head = &LinkNode{Val: &TreeNode{Val: nums[0]}}
	tag = head
	root := head.Val
	for i := 1; i < len(nums); {
		if nums[i] != math.MinInt32 {
			head.Val.Left = &TreeNode{Val: nums[i]}
			tag.Next = &LinkNode{Val: head.Val.Left}
			tag = tag.Next
		}
		i++
		if i < len(nums) && nums[i] != math.MinInt32 {
			head.Val.Right = &TreeNode{Val: nums[i]}
			tag.Next = &LinkNode{Val: head.Val.Right}
			tag = tag.Next
		}
		i++
		head = head.Next
	}
	return root
}
func ToString(root *TreeNode) string {
	var head, tag *LinkNode
	head = &LinkNode{Val: root}
	tag = head
	var arr []string
	for head != nil {
		cur := head.Val
		if cur != nil {
			arr = append(arr, strconv.Itoa(cur.Val))
			tag.Next = &LinkNode{Val: cur.Left}
			tag = tag.Next
			tag.Next = &LinkNode{Val: cur.Right}
			tag = tag.Next
		} else {
			arr = append(arr, "nil")
		}
		head = head.Next
	}
	return strings.Join(arr, ",")
}
