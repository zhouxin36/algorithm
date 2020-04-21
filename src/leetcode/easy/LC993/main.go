package main

import (
	"fmt"
	"leetcode"
	"math"
)

type TreeNode = leetcode.TreeNode

func main() {
	null := math.MinInt32
	fmt.Println(isCousins(leetcode.ToTreeNode([]int{1, 2, 3, null, 4, null, 5}), 5, 4))
	fmt.Println(!isCousins(leetcode.ToTreeNode([]int{1, 2, 3, 4}), 4, 3))
	fmt.Println(!isCousins(leetcode.ToTreeNode([]int{1, 2, 3, null, 4}), 2, 3))
}
func isCousins(root *TreeNode, x int, y int) bool {
	return depCousins(root, x, y, 0) == -2
}
func depCousins(root *TreeNode, x int, y int, level int) int {
	if root == nil {
		return -1
	}
	res := -1
	if root.Val == x || root.Val == y {
		fmt.Println(level)
		res = level
	}
	left := depCousins(root.Left, x, y, level+1)
	if left < -1 {
		return left
	}
	if res != -1 && left != -1 {
		if res == left && res != level+1 {
			return -2
		} else {
			return -3
		}
	} else {
		if left != -1 {
			res = left
		}
	}
	right := depCousins(root.Right, x, y, level+1)
	if right < -1 {
		return right
	}
	if res != -1 && right != -1 {
		if res == right && res != level+1 {
			return -2
		} else {
			return -3
		}
	} else {
		if right != -1 {
			return right
		} else {
			return res
		}
	}
}
