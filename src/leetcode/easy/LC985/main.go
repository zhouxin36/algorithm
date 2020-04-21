package main

import (
	"fmt"
)

func main() {
	//8,6,2,4
	fmt.Println(sumEvenAfterQueries([]int{1, 2, 3, 4}, [][]int{{1, 0}, {-3, 1}, {-4, 0}, {2, 3}}))
}
func sumEvenAfterQueries(A []int, queries [][]int) []int {
	sumEven := 0
	res := make([]int, len(A))
	for _, a := range A {
		if a&1 == 0 {
			sumEven += a
		}
	}
	for index, q := range queries {
		if A[q[1]]&1 == 0 {
			sumEven -= A[q[1]]
		}
		A[q[1]] += q[0]
		if A[q[1]]&1 == 0 {
			sumEven += A[q[1]]
		}
		res[index] = sumEven
	}
	return res
}
