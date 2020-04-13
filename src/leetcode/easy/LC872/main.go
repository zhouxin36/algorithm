package main

import (
	"fmt"
	"leetcode"
)

type TreeNode = leetcode.TreeNode

func main() {
	fmt.Println(!leafSimilar(&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}, &TreeNode{Val: 1, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 2}}))
}
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	r1 := doLeafSimilar(root1)
	r2 := doLeafSimilar(root2)
	if len(r1) != len(r2) {
		return false
	}
	for i := 0; i < len(r1); i++ {
		if r1[i] != r2[i] {
			return false
		}
	}
	return true
}
func doLeafSimilar(root1 *TreeNode) []int {
	if root1 == nil {
		return nil
	}
	if root1.Left == nil && root1.Right == nil {
		return []int{root1.Val}
	}
	return append(doLeafSimilar(root1.Left), doLeafSimilar(root1.Right)...)
}
