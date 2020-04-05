package main

import (
	"fmt"
	"leetcode"
)

//noinspection GoDuplicate
type TreeNode = leetcode.TreeNode

func main() {
	root := TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 7}}
	fmt.Println(root)
}
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	} else if root.Val == val {
		return root
	}
	if val > root.Val {
		return searchBST(root.Right, val)
	} else {
		return searchBST(root.Left, val)
	}
}
