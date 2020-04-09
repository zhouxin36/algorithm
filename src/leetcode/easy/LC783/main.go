package main

import (
	"fmt"
	"leetcode"
	"math"
)

type TreeNode = leetcode.TreeNode

func main() {
	fmt.Println(minDiffInBST(&TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 6}}) == 1)
	fmt.Println(minDiffInBST(&TreeNode{Val: 90, Left: &TreeNode{Val: 69, Left: &TreeNode{Val: 49, Right: &TreeNode{Val: 52}}, Right: &TreeNode{Val: 89}}}) == 1)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var pref int
var flag bool
var min int

func minDiffInBST(root *TreeNode) int {
	pref = 0
	flag = false
	min = math.MaxInt32
	doMinDiffInBST(root)
	return min
}

func doMinDiffInBST(root *TreeNode) int {
	if root == nil {
		return min
	}
	doMinDiffInBST(root.Left)
	if flag {
		v := root.Val - pref
		if v < min {
			min = v
		}
	} else {
		flag = true
	}
	pref = root.Val
	doMinDiffInBST(root.Right)
	return min
}
