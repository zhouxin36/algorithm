package main

import (
	"fmt"
	. "leetcode"
)

func main() {
	fmt.Println(sumRootToLeaf(ToTreeNode([]int{1, 0, 1, 0, 1, 0, 1})) == 22)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func sumRootToLeaf(root *TreeNode) int {
	sum := 0
	return doSumRootToLeaf(root, &sum, 0)
}
func doSumRootToLeaf(root *TreeNode, sum *int, val int) int {
	if root == nil {
		return *sum
	}
	v := val<<1 | root.Val
	if root.Right == nil && root.Left == nil {
		*sum += v
	} else {
		doSumRootToLeaf(root.Left, sum, v)
		doSumRootToLeaf(root.Right, sum, v)
	}
	return *sum
}
