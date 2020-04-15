package main

import (
	"fmt"
	"leetcode"
	"math"
)

type TreeNode = leetcode.TreeNode

func main() {
	null := math.MinInt32
	node := leetcode.ToTreeNode([]int{5, 3, 6, 2, 4, null, 8, 1, null, null, null, 7, 9})
	node = increasingBST(node)
	fmt.Println(leetcode.ToString(node))
}

var pref *TreeNode
var r *TreeNode

func increasingBST(root *TreeNode) *TreeNode {
	pref = nil
	r = nil
	return doIncreasingBST(root)
}
func doIncreasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return r
	}
	doIncreasingBST(root.Left)
	if r == nil {
		r = root
	}
	if pref == nil {
		pref = root
	} else {
		pref.Right = root
		pref.Left = nil
		pref = pref.Right
	}
	doIncreasingBST(root.Right)
	return r
}
