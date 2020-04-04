package main

import (
	"fmt"
	"leetcode"
)

type TreeNode = leetcode.TreeNode

func main() {
	a2 := TreeNode{
		Val: 1, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 5}},
	}
	fmt.Println(longestUnivaluePath(&a2) == 2)
	a1 := TreeNode{
		Val: 5, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 5}},
	}
	fmt.Println(longestUnivaluePath(&a1) == 2)
}

var ans int

func longestUnivaluePath(root *TreeNode) int {
	ans = 0
	doLongestUnivaluePath(root)
	return ans
}

func doLongestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := doLongestUnivaluePath(root.Left)
	right := doLongestUnivaluePath(root.Right)
	var arrowLeft, arrowRight int
	if root.Left != nil && root.Val == root.Left.Val {
		arrowLeft = left + 1
	}
	if root.Right != nil && root.Val == root.Right.Val {
		arrowRight = right + 1
	}
	if arrowLeft+arrowRight > ans {
		ans = arrowRight + arrowLeft
	}
	if arrowRight > arrowLeft {
		return arrowRight
	} else {
		return arrowLeft
	}
}
