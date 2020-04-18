package main

import "fmt"

func main() {
	fmt.Println(minDeletionSize([]string{"cba", "daf", "ghi"}) == 1)
	fmt.Println(minDeletionSize([]string{"a", "b"}) == 0)
	fmt.Println(minDeletionSize([]string{"zyx", "wvu", "tsr"}) == 3)
}
func minDeletionSize(A []string) int {
	ans := 0
	for c := 0; c < len(A[0]); c++ {
		for r := 0; r < len(A)-1; r++ {
			if A[r][c] > A[r+1][c] {
				ans++
				break
			}
		}
	}
	return ans
}
