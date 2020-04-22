package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(largestSumAfterKNegations([]int{4, 2, 3}, 1) == 5)
	fmt.Println(largestSumAfterKNegations([]int{3, -1, 0, 2}, 3) == 6)
	fmt.Println(largestSumAfterKNegations([]int{2, -3, -1, 5, -4}, 2) == 13)
}
func largestSumAfterKNegations(A []int, K int) int {
	sort.Ints(A)
	i := 0
	for ; i < len(A) && K > 0; i++ {
		if A[i] < 0 {
			A[i] = -A[i]
			K--
		} else {
			break
		}
	}
	if K > 0 {
		if K&1 != 0 {
			if i == 0 {
				A[i] = -A[i]
			} else {
				if A[i] < A[i-1] {
					A[i] = -A[i]
				} else {
					A[i-1] = -A[i-1]
				}
			}
		}
	}
	var total int
	for _, a := range A {
		total += a
	}
	return total
}
