package main

import (
	"fmt"
	. "leetcode"
)

func main() {
	//== "[2,4,3,1]"
	fmt.Println(IntToString(sortArrayByParity([]int{3, 1, 2, 4})))
}
func sortArrayByParity(A []int) []int {
	head, tag := 0, len(A)-1
	for head < tag {
		for head <= tag && A[head]&1 == 0 {
			head++
		}
		for tag >= head && A[tag]&1 == 1 {
			tag--
		}
		if head >= tag {
			break
		}
		A[head] ^= A[tag]
		A[tag] ^= A[head]
		A[head] ^= A[tag]
	}
	return A
}
