package main

import (
	"fmt"
	"leetcode"
	"math"
)

type TreeNode = leetcode.TreeNode

func main() {
	null := math.MinInt32
	fmt.Println(rangeSumBST(leetcode.ToTreeNode([]int{10, 5, 15, 3, 7, null, 18}), 7, 15) == 32)
	fmt.Println(rangeSumBST(leetcode.ToTreeNode([]int{10, 5, 15, 3, 7, 13, 18, 1, null, 6}), 6, 10) == 23)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}
	var sum int
	if root.Val >= L && root.Val <= R {
		sum += root.Val
	}
	if root.Val >= L {
		sum += rangeSumBST(root.Left, L, R)
	}
	if root.Val <= R {
		sum += rangeSumBST(root.Right, L, R)
	}
	return sum
}
