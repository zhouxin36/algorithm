package main

import (
	"fmt"
	. "leetcode"
	"math"
)

func main() {
	null := math.MinInt32
	fmt.Println(isUnivalTree(ToTreeNode([]int{1, 1, 1, 1, 1, null, 1})))
	fmt.Println(!isUnivalTree(ToTreeNode([]int{2, 2, 2, 5, 2})))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isUnivalTree(root *TreeNode) bool {
	return doUnivalTree(root, root.Val)
}
func doUnivalTree(root *TreeNode, Val int) bool {
	if root == nil {
		return true
	}
	if root.Val != Val {
		return false
	}
	return doUnivalTree(root.Right, Val) && doUnivalTree(root.Left, Val)
}
